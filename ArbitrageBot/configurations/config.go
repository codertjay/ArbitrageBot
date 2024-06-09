package configurations

import (
	arbitrageABI "ArbitrageBot/ArbitrageBot/abi"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"
)

// Sushi swap deployed address doc https://dev.sushi.com/docs/Products/Classic%20AMM/Deployment%20Addresses
//  https://goethereumbook.org/event-subscribe/ check subscribing to event

type ConfigInterface interface {
	SetupRPCURL() *Config
	SetupDecentralizedExchange() *Config
	SetupMainTokenAddress() *Config
	SetupArbitrageAddress() *Config
	SetupAuthentication() *Config
	SetupTokens() (*Config, error)
	SetupETHClient(rpcURL string) (*ethclient.Client, error)
	SetUpRPCURLList() *Config
	Setup() (*Config, error)
	WatchSwapForExchange(client *ethclient.Client, exchange DecentralizedExchange)
	HandleSwapEvent(vLog types.Log, exchangeName, routerV2 string, pairAddress, mainToken, externalToken common.Address)
	CalculatePriceDifference(params CalculatePriceDifferenceParams)
	SetupHelper() *Config
}

func NewConfig() ConfigInterface {
	return &Config{}
}

type Config struct {
	DecentralizedExchanges []DecentralizedExchange
	Tokens                 []Tokens
	HTTPRPCURLList         []string
	HTTPRPCURL             string
	WSSRPCURL              string
	MainTokenAddress       common.Address
	ETHClient              *ethclient.Client
	ArbitrageAddress       common.Address
	Arbitrage              *arbitrageABI.FlashLoanArbitrage
	Authentication         *bind.TransactOpts
	Helper                 HelperInterface
}

type DecentralizedExchange struct {
	Name           string
	RouterV2       string
	FactoryAddress string
}

type Token struct {
	Name    string
	Decimal int
	Address common.Address
}

type Tokens struct {
	MainToken     Token // could be WBNB or WETH base on the main pairs
	ExternalToken Token // this is the external pair
}

func (cfg *Config) Setup() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg.SetupRPCURL()
	cfg.SetUpRPCURLList()
	cfg.SetupDecentralizedExchange()
	cfg.SetupMainTokenAddress()

	ethClient, err := cfg.SetupETHClient(cfg.HTTPRPCURL)
	if err != nil {
		return nil, err
	}
	cfg.ETHClient = ethClient

	cfg.SetupAuthentication()
	cfg.SetupHelper()
	cfg.SetupArbitrageAddress()

	cfg, err = cfg.LoadDexTokensFromFile()
	if err != nil {
		return nil, err
	}

	err = cfg.WatchSwap()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
func (cfg *Config) SetupRPCURL() *Config {
	cfg.HTTPRPCURL = "https://polygon-mainnet.alchemyapi.io/v2/v314L6MKTskIeFD4nXB-OEPeFrfWJ0ey"
	cfg.WSSRPCURL = "wss://polygon-mainnet.g.alchemy.com/v2/v314L6MKTskIeFD4nXB-OEPeFrfWJ0ey"
	return cfg
}

func (cfg *Config) SetupETHClient(rpcURL string) (*ethclient.Client, error) {
	client, err := rpc.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(client), nil
}

func (cfg *Config) SetUpRPCURLList() *Config {
	cfg.HTTPRPCURLList = []string{
		"https://polygon-mainnet.alchemyapi.io/v2/_E_R6fq4lVfmVYxXBFvuMQ0vL7pCW9Jk",
		"https://polygon-mainnet.alchemyapi.io/v2/Jhcj8wHIUWJEe__0oMD8VW7nNZh8ZwKp",
		"https://polygon-mainnet.alchemyapi.io/v2/ydQb_t8kmgv_Q8lBoXYxD9wRJOP0SYFA",
		"https://polygon-mainnet.alchemyapi.io/v2/psn_5bn2xDgaZ6JrcW4Pthht1YLaKx6e",
		"https://polygon-mainnet.alchemyapi.io/v2/DBBhNLOybHWsr_Rj_iT-m5K9dPWVfsAw",
		"https://polygon-mainnet.alchemyapi.io/v2/MqotdQE4eTNuHJa165AVlJw8d7iu1lyV",
		"https://polygon-mainnet.alchemyapi.io/v2/CDRsfhmJ3sfhou6IeQY5kEadLr08Cyn1",
		"https://polygon-mainnet.alchemyapi.io/v2/udtdESou598cCF8AwxFDwwjNlngtF860",
		"https://polygon-mainnet.alchemyapi.io/v2/zXGOpsx4OFS0OhrxuBlpWHaCdshXtYAt",
		"https://polygon-mainnet.alchemyapi.io/v2/v314L6MKTskIeFD4nXB-OEPeFrfWJ0ey",
		"https://polygon-mainnet.alchemyapi.io/v2/vJfuk98N3ctdLMM1ZWuqHKBftqXxNnBR",
		"https://polygon-mainnet.alchemyapi.io/v2/0AYzfz3g8RnWaWCABo2GMTfS_5IkV1JD",
	}

	return cfg
}

func (cfg *Config) SetupDecentralizedExchange() *Config {
	cfg.DecentralizedExchanges = []DecentralizedExchange{
		{
			Name:           "QUICKSWAP",
			RouterV2:       "0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff",
			FactoryAddress: "0x5757371414417b8c6caad45baef941abc7d3ab32",
		},
		{
			Name:           "SUSHISWAP",
			RouterV2:       "0x1b02da8cb0d097eb8d57a175b88c7d8b47997506",
			FactoryAddress: "0xc35dadb65012ec5796536bd9864ed8773abc74c4",
		},
		{
			Name:           "DFYN",
			RouterV2:       "0xA102072A4C07F06EC3B4900FDC4C7B80b6c57429",
			FactoryAddress: "0xE7Fb3e833eFE5F9c441105EB65Ef8b261266423B", // dont have much pairs
		},
	}

	return cfg
}

func (cfg *Config) SetupMainTokenAddress() *Config {
	cfg.MainTokenAddress = common.HexToAddress(os.Getenv("MAIN_TOKEN_ADDRESS"))
	return cfg
}

func (cfg *Config) SetupArbitrageAddress() *Config {
	var err error

	cfg.ArbitrageAddress = common.HexToAddress(os.Getenv("ARBITRAGE_ADDRESS"))

	cfg.Arbitrage, err = arbitrageABI.NewFlashLoanArbitrage(cfg.ArbitrageAddress, cfg.ETHClient)
	if err != nil {
		log.Fatal("Error creating arbitrage contract:", err)
		return nil
	}

	return cfg
}

func (cfg *Config) SetupAuthentication() *Config {
	chainIDStr := os.Getenv("CHAIN_ID")
	if chainIDStr == "" {
		log.Fatalf("CHAIN_ID environment variable not set")
	}
	chainIDInt, err := strconv.Atoi(chainIDStr)
	if err != nil {
		log.Fatalf("Failed to convert CHAIN_ID to integer: %v", err)
	}
	chainID := big.NewInt(int64(chainIDInt))

	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		log.Fatalf("PRIVATE_KEY environment variable not set")
	}

	privateKeyHex = privateKeyHex[2:]

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	gasPrice, err := cfg.ETHClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	auth.GasLimit = uint64(900000)
	auth.GasPrice = gasPrice

	cfg.Authentication = auth
	return cfg
}

func (cfg *Config) SetupHelper() *Config {
	helper := NewHelper(cfg.ETHClient)
	cfg.Helper = helper
	return cfg
}

type DexTokens struct {
	MainToken     common.Address
	ExternalToken common.Address
	Exist         bool
}

func (cfg *Config) SetupTokens() (*Config, error) {
	var dexTokens []DexTokens
	const maxConcurrentGoroutines = 10000

	semaphore := make(chan struct{}, maxConcurrentGoroutines)
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for counter, exchange := range cfg.DecentralizedExchanges {

		// Create a new instance of the Uniswap V2 Factory contract
		factoryAddress := common.HexToAddress(exchange.FactoryAddress)
		factory, err := arbitrageABI.NewUniswapV2Factory(factoryAddress, cfg.ETHClient)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		// Get the total number of pairs
		totalPairs, err := factory.AllPairsLength(nil)
		if err != nil {
			log.Println(dexTokens)
			log.Println("the pair factory address is: ", factoryAddress)
			log.Fatal("an error occurred getting the total pairs", err)
		}

		// Create a rate limiter that allows 5 calls per second
		rateLimit := time.Tick(time.Second / 5)

		// Iterate over all pairs and check if they contain WBNB
		for i := big.NewInt(0); i.Cmp(totalPairs) < 0; i.Add(i, big.NewInt(1)) {
			semaphore <- struct{}{}
			wg.Add(1)

			go func(i *big.Int) {
				defer wg.Done()
				defer func() {
					<-semaphore
				}()

				<-rateLimit

				pairAddress, err := factory.AllPairs(nil, i)
				if err != nil {
					log.Println("Error fetching pair address:", err)
					return
				}

				log.Printf("pairAddress from  %s and the pair address is %s \n", exchange.Name, pairAddress.String())

				// Instantiate UniswapV2Pair contract
				pairContract, err := arbitrageABI.NewUniswapV2Pair(pairAddress, cfg.ETHClient)
				if err != nil {
					log.Println("Error creating pair contract:", err)
					return
				}

				// Get tokens in the pair
				token0, err := pairContract.Token0(nil)
				if err != nil {
					log.Println("Error fetching token0:", err)
					return
				}
				token1, err := pairContract.Token1(nil)
				if err != nil {
					log.Println("Error fetching token1:", err)
					return
				}

				_, err = factory.GetPair(nil, cfg.MainTokenAddress, token0)
				if err != nil {
					log.Println("Error getting pair:", err)
					return
				}

				mu.Lock()
				defer mu.Unlock()

				log.Println("token0: ", token0.String())

				if counter == 0 {
					dexToken := DexTokens{
						MainToken:     cfg.MainTokenAddress,
						ExternalToken: token0,
						Exist:         false,
					}
					dexTokens = append(dexTokens, dexToken)

				} else {
					// For subsequent exchanges, check if tokens exist in dexTokens and update their existence status
					for k, tokens := range dexTokens {
						if (tokens.MainToken == token0 && tokens.ExternalToken == token1) ||
							(tokens.MainToken == token1 && tokens.ExternalToken == token0) {
							dexTokens[k].Exist = true
							break
						}
					}
				}

			}(new(big.Int).Set(i))

			// Check if we've reached the maximum number of concurrent goroutines
			if i.Cmp(big.NewInt(maxConcurrentGoroutines)) >= 0 {
				break
			}

		}

		// Wait for all goroutines to complete before moving to the next exchange
		wg.Wait()

		// After looping a dex completely, update the dexTokens' existence status to false
		if counter != len(cfg.DecentralizedExchanges)-1 {
			for i := range dexTokens {
				dexTokens[i].Exist = false
			}
		}

	}

	for _, tokens := range dexTokens {

		if !tokens.Exist {
			continue
		}
		if tokens.MainToken == cfg.MainTokenAddress {
			cfg.Tokens = append(cfg.Tokens, Tokens{
				MainToken: Token{
					Name:    "",
					Decimal: 0,
					Address: tokens.MainToken,
				},
				ExternalToken: Token{
					Name:    "",
					Decimal: 0,
					Address: tokens.ExternalToken,
				},
			})
		} else if tokens.ExternalToken == cfg.MainTokenAddress {
			cfg.Tokens = append(cfg.Tokens, Tokens{
				MainToken: Token{
					Name:    "",
					Decimal: 0,
					Address: tokens.ExternalToken,
				},
				ExternalToken: Token{
					Name:    "",
					Decimal: 0,
					Address: tokens.MainToken,
				},
			})
		}
	}

	data, err := json.Marshal(cfg.Tokens)
	if err != nil {
		log.Println("Error marshaling DexTokens to JSON:", err)
		return nil, err
	}

	// Write JSON data to file
	err = ioutil.WriteFile("tokens.json", data, 0644)
	if err != nil {
		log.Println("Error writing DexTokens JSON to file:", err)
		return nil, err
	}

	log.Println("DexTokens saved to  tokens.json")
	return cfg, nil
}

func (cfg *Config) LoadDexTokensFromFile() (*Config, error) {

	// Read JSON data from file
	data, err := ioutil.ReadFile("tokens.json")
	if err != nil {
		log.Println("Error reading DexTokens JSON from file:", err)
		return nil, err
	}

	err = json.Unmarshal(data, &cfg.Tokens)
	if err != nil {
		log.Println("Error unmarshalling DexTokens JSON:", err)
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) WatchSwap() error {
	client, err := ethclient.Dial(cfg.WSSRPCURL)
	if err != nil {
		log.Fatal("Error connecting to Ethereum client: ", err)
		return err
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 1)

	for _, exchange := range cfg.DecentralizedExchanges {
		wg.Add(1)

		go func(exchange DecentralizedExchange) {
			defer wg.Done()
			defer func() {
				<-semaphore
			}()
			_exchange := cfg.DecentralizedExchanges[0]
			cfg.WatchSwapForExchange(client, _exchange)
		}(exchange)
	}

	wg.Wait()
	return nil
}

func (cfg *Config) WatchSwapForExchange(client *ethclient.Client, exchange DecentralizedExchange) {
	// todo: add context in here and remove fatalf if i dont need it to run for life
	factoryAddress := common.HexToAddress(exchange.FactoryAddress)
	factory, err := arbitrageABI.NewUniswapV2Factory(factoryAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Uniswap V2 factory contract: %v", err)
	}

	// Create a rate limiter that allows 5 calls per second
	rateLimit := time.Tick(time.Second / 50)
	var wg sync.WaitGroup

	for _, tokenPair := range cfg.Tokens {
		wg.Add(1)
		go func(tokenPair Tokens) {
			defer wg.Done()

			// Wait for the rate limiter
			<-rateLimit

			pairAddress, err := factory.GetPair(nil, tokenPair.MainToken.Address, tokenPair.ExternalToken.Address)
			if err != nil {
				log.Println("Failed to get pair address: %v", err)
				return
			}

			if pairAddress == common.HexToAddress("0x0000000000000000000000000000000000000000") {
				log.Println("Invalid pair address, skipping...")
				return
			}

			log.Println("The exchange Name ", exchange.Name, "\n The Pair Address pairAddress: ", pairAddress.String(), "\n The Main Token :", tokenPair.MainToken.Address.String(),
				"\n The External Token: ", tokenPair.ExternalToken.Address.String())

			query := ethereum.FilterQuery{
				Addresses: []common.Address{pairAddress},
			}

			logsCh := make(chan types.Log)
			sub, err := client.SubscribeFilterLogs(context.Background(), query, logsCh)
			if err != nil {
				log.Println("Error subscribing to logs: %v", err)
				return
			}

			lastBlock := big.NewInt(0)

			for {
				select {
				case err := <-sub.Err():
					log.Println("Error in subscription:", err)
					time.Sleep(time.Second * 10) // Exponential backoff can be implemented here
					sub, err = client.SubscribeFilterLogs(context.Background(), query, logsCh)
					if err != nil {
						log.Fatalf("Error resubscribing to logs: %v", err)
					}
				case vLog := <-logsCh:
					if vLog.BlockNumber <= lastBlock.Uint64() {
						continue
					}
					lastBlock.SetUint64(vLog.BlockNumber)

					cfg.HandleSwapEvent(vLog, exchange.Name, exchange.RouterV2, pairAddress, tokenPair.MainToken.Address, tokenPair.ExternalToken.Address)
					log.Printf("Swap event detected on  Exchainge: %s MainToken: %s External Token: %s \n", exchange.Name, tokenPair.MainToken.Address.String(), tokenPair.ExternalToken.Address.String())
				}
			}
		}(tokenPair)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Block the goroutine indefinitely
	select {}
}

func (cfg *Config) HandleSwapEvent(vLog types.Log, exchangeName, routerV2 string, pairAddress, mainToken, externalToken common.Address) {
	// Extract relevant information from the event log
	// todo: get the log information's

	// Calculate price difference using the helper
	params := CalculatePriceDifferenceParams{
		ExchangeName:  exchangeName,
		RouterV2:      routerV2,
		MainToken:     mainToken,
		ExternalToken: externalToken,
		PairAddress:   pairAddress,
		Log:           vLog,
	}
	cfg.CalculatePriceDifference(params)
}

type CalculatePriceDifferenceParams struct {
	ExchangeName  string
	RouterV2      string
	MainToken     common.Address
	ExternalToken common.Address
	PairAddress   common.Address
	Log           types.Log
}

type PriceDifferenceExchange struct {
	ExchangeName string
}

func (cfg *Config) CalculatePriceDifference(params CalculatePriceDifferenceParams) {
	var mainExchangePrice *big.Float
	var mainRouterV2 string
	var threshold = big.NewInt(4)
	// Define the print amount as 20 tokens with 18 decimals
	tokenDecimals := 6
	printAmount := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokenDecimals)), nil)
	printAmount.Mul(printAmount, big.NewInt(500)) // 20 * 10^18

	for _, exchange := range cfg.DecentralizedExchanges {
		if exchange.Name == params.ExchangeName {
			// Perform calculations for the specified exchange
			reserves, err := cfg.Helper.GetReserves(params.PairAddress)
			if err != nil {
				log.Printf("Error getting reserves for exchange %s: %v\n", exchange.Name, err)
				return
			}

			mainExchangePrice = cfg.Helper.CalculatePrice(reserves)
			mainRouterV2 = exchange.RouterV2
		}
	}

	// Let's calculate the price differences in other exchanges to see if there is an  opportunity
	for _, exchange := range cfg.DecentralizedExchanges {
		if exchange.Name != params.ExchangeName {

			factory := common.HexToAddress(exchange.FactoryAddress)
			pairAddress, err := cfg.Helper.GetPair(factory, params.MainToken, params.ExternalToken)
			if err != nil {
				log.Println("Error getting pair on ")
				return
			}

			reserves, err := cfg.Helper.GetReserves(pairAddress)
			if err != nil {
				log.Printf("Error getting reserves for exchange %s: %v\n\n", exchange.Name, err)
				return
			}

			externalExchangePrice := cfg.Helper.CalculatePrice(reserves)

			isArbitrage, err := cfg.Arbitrage.CheckProfitability(nil, common.HexToAddress(mainRouterV2),
				common.HexToAddress(exchange.RouterV2), params.MainToken, params.ExternalToken, printAmount, threshold)
			if err != nil {
				log.Println("Error checking profitability: ", err)
				return
			}

			if isArbitrage.IsProfitable {
				log.Println("Arbitrage opportunity detected and the profit is ", isArbitrage.PercentageProfit)

				if isArbitrage.Direction == "ATOB" {
					fmt.Printf(" Arbitrage opportunity detected: Buy on %s at %v , Sell on %s at %v and the block number is %v and the token is %v \n",
						params.ExchangeName, mainExchangePrice, exchange.Name, externalExchangePrice, params.Log.BlockNumber, params.ExternalToken)

					transactionPrint, err := cfg.Arbitrage.MakeFlashLoan(cfg.Authentication,
						common.HexToAddress(params.RouterV2), common.HexToAddress(exchange.RouterV2), params.MainToken, params.ExternalToken, printAmount)
					if err != nil {
						log.Println("Error making flash loan: ", err)
						return
					}
					log.Println("The transaction hash is: ", transactionPrint.Hash().Hex())
				} else if isArbitrage.Direction == "BTOA" {
					fmt.Printf(" Arbitrage opportunity detected: Buy on %s at %v, Sell on %s at %v  and the block number is %v annd the token is %v\n",
						exchange.Name, externalExchangePrice, params.ExchangeName, mainExchangePrice, params.Log.BlockNumber, params.ExternalToken)

					transactionPrint, err := cfg.Arbitrage.MakeFlashLoan(cfg.Authentication,
						common.HexToAddress(exchange.RouterV2), common.HexToAddress(params.RouterV2), params.MainToken, params.ExternalToken, printAmount)
					if err != nil {
						log.Println("Error making flash loan: ", err)
						return
					}
					log.Println("The transaction hash is: ", transactionPrint.Hash().Hex())
				}
			} else {
				fmt.Println("No arbitrage opportunity.")
			}

		}
	}

}
