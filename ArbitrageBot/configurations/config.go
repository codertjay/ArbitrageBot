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
	Arbitrage              *arbitrageABI.Arbitrage
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
	cfg.SetupRPCURL()
	cfg.SetUpRPCURLList()
	cfg.SetupDecentralizedExchange()
	cfg.SetupMainTokenAddress()
	cfg.SetupArbitrageAddress()
	cfg.SetupHelper()
	cfg.SetupAuthentication()

	ethClient, err := cfg.SetupETHClient(cfg.HTTPRPCURL)
	if err != nil {
		return nil, err
	}
	cfg.ETHClient = ethClient

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
	cfg.HTTPRPCURL = "https://bsc-mainnet.core.chainstack.com/2b143ce4e436f2bc1261f7b0851d272d" // rpc-url needed
	cfg.WSSRPCURL = "wss://bsc-mainnet.core.chainstack.com/ws/2b143ce4e436f2bc1261f7b0851d272d" // rpc-url needed
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
		"https://bsc-dataseed1.binance.org/",
		"https://bsc-dataseed2.binance.org/",
		"https://bsc-dataseed3.binance.org/",
		"https://bsc-dataseed4.binance.org/",
		"https://bsc-dataseed1.defibit.io/",
		"https://bsc-dataseed2.defibit.io/",
		"https://bsc-dataseed3.defibit.io/",
		"https://bsc-dataseed4.defibit.io/",
		"https://bsc-dataseed1.ninicoin.io/",
		"https://bsc-dataseed2.ninicoin.io/",
		"https://bsc-dataseed3.ninicoin.io/",
		"https://bsc-dataseed4.ninicoin.io/",
	}

	return cfg
}

func (cfg *Config) SetupDecentralizedExchange() *Config {
	cfg.DecentralizedExchanges = []DecentralizedExchange{
		{
			Name:           "PANCAKESWAP",
			RouterV2:       "0x10ED43C718714eb63d5aA57B78B54704E256024E",
			FactoryAddress: "0xBCfCcbde45cE874adCB698cC183deBcF17952812",
		},
		{
			Name:           "APESWAP",
			RouterV2:       "0xcF0feBd3f17CEf5b47b0cD257aCf6025c5BFf3b7",
			FactoryAddress: "0x0841BD0B734E4F5853f0dD8d7Ea041c241fb0Da6",
		},
		//{
		//	Name:           "BAKERYSWAP",
		//	RouterV2:       "0xCDe540d7eAFE93aC5fE6233Bee57E1270D3E330F",
		//	FactoryAddress: "0x01bF7C66C6BD861915Cdaae475042d3c4BA9eF5d",
		//},
		//{
		//	Name:           "MDEX",
		//	RouterV2:       "0xc6aF770101dA859d680E0829380748CCcD8F7984",
		//	FactoryAddress: "0x3E5C63644E683549055b9Be8653de26E0B4CD36E",
		//},
		//{
		//	Name:           "UNISWAP",
		//	RouterV2:       "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
		//	FactoryAddress: "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f",
		//},
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

	cfg.Arbitrage, err = arbitrageABI.NewArbitrage(cfg.ArbitrageAddress, cfg.ETHClient)
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

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	auth.GasLimit = uint64(300000)
	auth.GasPrice = big.NewInt(20000000000)

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
	const maxConcurrentGoroutines = 5000

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

		// Iterate over all pairs and check if they contain WBNB
		for i := big.NewInt(0); i.Cmp(totalPairs) < 0; i.Add(i, big.NewInt(1)) {
			semaphore <- struct{}{}
			wg.Add(1)

			go func(i *big.Int) {
				defer wg.Done()
				defer func() { <-semaphore }()

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
	for _, exchange := range cfg.DecentralizedExchanges {
		wg.Add(1)
		go func(exchange DecentralizedExchange) {
			defer wg.Done()
			cfg.WatchSwapForExchange(client, exchange)
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
	rateLimit := time.Tick(time.Second / 5)
	var wg sync.WaitGroup

	for _, tokenPair := range cfg.Tokens {
		wg.Add(1)
		go func(tokenPair Tokens) {
			defer wg.Done()

			// Wait for the rate limiter
			<-rateLimit

			pairAddress, err := factory.GetPair(nil, tokenPair.MainToken.Address, tokenPair.ExternalToken.Address)
			if err != nil {
				log.Fatalf("Failed to get pair address: %v", err)
			}

			log.Println("The exchange Name ", exchange.Name, "\n The Pair Address pairAddress: ", pairAddress.String(), "\n The Main Token :", tokenPair.MainToken.Address.String(),
				"\n The External Token: ", tokenPair.ExternalToken.Address.String())

			query := ethereum.FilterQuery{
				Addresses: []common.Address{pairAddress},
			}

			logsCh := make(chan types.Log)
			sub, err := client.SubscribeFilterLogs(context.Background(), query, logsCh)
			if err != nil {
				log.Fatalf("Error subscribing to logs: %v", err)
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
	// todo: get the log informations

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
	var threshold float64 = 0.5
	var printAmount *big.Int = big.NewInt(20)

	for _, exchange := range cfg.DecentralizedExchanges {
		if exchange.Name == params.ExchangeName {
			// Perform calculations for the specified exchange
			reserves, err := cfg.Helper.GetReserves(params.PairAddress)
			if err != nil {
				log.Printf("Error getting reserves for exchange %s: %v\n", exchange.Name, err)
			}

			mainExchangePrice = cfg.Helper.CalculatePrice(reserves)
		}
	}

	// Let's calculate the price differences in other exchanges to see if there is an  opportunity
	for _, exchange := range cfg.DecentralizedExchanges {
		if exchange.Name != params.ExchangeName {

			factory := common.HexToAddress(exchange.FactoryAddress)
			pairAddress, err := cfg.Helper.GetPair(factory, params.MainToken, params.ExternalToken)
			if err != nil {
				log.Println("Error getting pair on ")
			}

			reserves, err := cfg.Helper.GetReserves(pairAddress)
			if err != nil {
				log.Printf("Error getting reserves for exchange %s: %v\n\n", exchange.Name, err)
			}

			externalExchangePrice := cfg.Helper.CalculatePrice(reserves)

			isArbitrage, direction := cfg.Helper.CheckArbitrageOpportunity(mainExchangePrice, externalExchangePrice, threshold)
			if isArbitrage {
				if direction == "AtoB" {
					fmt.Printf(" Arbitrage opportunity detected: Buy on %s at %v , Sell on %s at %v and the block number is %v and the token is %v \n",
						params.ExchangeName, mainExchangePrice, exchange.Name, externalExchangePrice, params.Log.BlockNumber, params.ExternalToken)

					transactionPrint, err := cfg.Arbitrage.PrintMoney(cfg.Authentication,
						common.HexToAddress(params.RouterV2), common.HexToAddress(exchange.RouterV2), params.MainToken, params.ExternalToken, printAmount)
					if err != nil {
						return
					}
					log.Println("The transaction Printed is: ", transactionPrint)
				} else {
					fmt.Printf(" Arbitrage opportunity detected: Buy on %s at %v, Sell on %s at %v  and the block number is %v annd the token is %v\n",
						exchange.Name, externalExchangePrice, params.ExchangeName, mainExchangePrice, params.Log.BlockNumber, params.ExternalToken)

					transactionPrint, err := cfg.Arbitrage.PrintMoney(cfg.Authentication,
						common.HexToAddress(exchange.RouterV2), common.HexToAddress(params.RouterV2), params.MainToken, params.ExternalToken, printAmount)
					if err != nil {
						return
					}
					log.Println("The transaction Printed is: ", transactionPrint)
				}
			} else {
				fmt.Println("No arbitrage opportunity.")
			}

		}
	}

}
