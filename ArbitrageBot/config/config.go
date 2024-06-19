package config

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
	"time"
)

// ConfigInterface definition remains unchanged

type Config struct {
	ETHClient              *ethclient.Client
	Arbitrage              *arbitrageABI.FlashLoanArbitrage
	Authentication         *bind.TransactOpts
	DecentralizedExchanges []DecentralizedExchange
	TrendingTokens         []string
	PairAddresses          map[common.Address]Tokens
	HTTPRPCURLList         []string
	HTTPRPCURL             string
	WSSRPCURL              string
	MainTokenAddress       common.Address
	ArbitrageAddress       common.Address
}

type DecentralizedExchange struct {
	Name           string
	RouterV2       string
	FactoryAddress string
	MKey           string
}

type Token struct {
	Name    string
	Decimal int
	Address common.Address
}

type Tokens struct {
	MainToken              Token
	ExternalToken          Token
	DecentralizedExchanges []DecentralizedExchange
}

func (cfg *Config) Setup() (*Config, error) {
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
	cfg.SetupArbitrageAddress()

	cfg, err = cfg.LoadTrendingTokens()
	if err != nil {
		return nil, err
	}

	err = cfg.SetupGetPairAddresses()
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
	cfg.HTTPRPCURL = "https://polygon-mainnet.alchemyapi.io/v2/ydQb_t8kmgv_Q8lBoXYxD9wRJOP0SYFA"
	cfg.WSSRPCURL = "wss://polygon-mainnet.g.alchemy.com/v2/ydQb_t8kmgv_Q8lBoXYxD9wRJOP0SYFA"
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
		"https://eth-mainnet.alchemyapi.io/v2/_E_R6fq4lVfmVYxXBFvuMQ0vL7pCW9Jk",
		"https://eth-mainnet.alchemyapi.io/v2/Jhcj8wHIUWJEe__0oMD8VW7nNZh8ZwKp",
		"https://eth-mainnet.alchemyapi.io/v2/ydQb_t8kmgv_Q8lBoXYxD9wRJOP0SYFA",
		"https://eth-mainnet.alchemyapi.io/v2/psn_5bn2xDgaZ6JrcW4Pthht1YLaKx6e",
		"https://eth-mainnet.alchemyapi.io/v2/DBBhNLOybHWsr_Rj_iT-m5K9dPWVfsAw",
		"https://eth-mainnet.alchemyapi.io/v2/MqotdQE4eTNuHJa165AVlJw8d7iu1lyV",
		"https://eth-mainnet.alchemyapi.io/v2/CDRsfhmJ3sfhou6IeQY5kEadLr08Cyn1",
		"https://eth-mainnet.alchemyapi.io/v2/udtdESou598cCF8AwxFDwwjNlngtF860",
		"https://eth-mainnet.alchemyapi.io/v2/zXGOpsx4OFS0OhrxuBlpWHaCdshXtYAt",
		"https://eth-mainnet.alchemyapi.io/v2/v314L6MKTskIeFD4nXB-OEPeFrfWJ0ey",
		"https://eth-mainnet.alchemyapi.io/v2/vJfuk98N3ctdLMM1ZWuqHKBftqXxNnBR",
		"https://eth-mainnet.alchemyapi.io/v2/0AYzfz3g8RnWaWCABo2GMTfS_5IkV1JD",
	}
	return cfg
}

func (cfg *Config) SetupDecentralizedExchange() *Config {
	cfg.DecentralizedExchanges = []DecentralizedExchange{
		{
			Name:           "QUICKSWAP",
			RouterV2:       "0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff",
			FactoryAddress: "0x5757371414417b8c6caad45baef941abc7d3ab32",
			MKey:           "M1",
		},
		{
			Name:           "SUSHISWAP",
			RouterV2:       "0x1b02da8cb0d097eb8d57a175b88c7d8b47997506",
			FactoryAddress: "0xc35dadb65012ec5796536bd9864ed8773abc74c4",
			MKey:           "M2",
		},
		{
			Name:           "DFYN",
			RouterV2:       "0xA102072A4C07F06EC3B4900FDC4C7B80b6c57429",
			FactoryAddress: "0xE7Fb3e833eFE5F9c441105EB65Ef8b261266423B",
			MKey:           "M3",
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

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY")[2:])
	if err != nil {
		log.Fatalf("Failed to convert private key to ECDSA: %v", err)
	}
	cfg.Authentication, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// Set gas price to 50 Gwei (50 * 10^9 Wei)
	cfg.Authentication.GasPrice = big.NewInt(500 * 1000000000)

	return cfg
}

func (cfg *Config) LoadTrendingTokens() (*Config, error) {
	file, err := os.Open("trendingTokens.json")
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	err = json.Unmarshal(bytes, &cfg.TrendingTokens)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config file: %v", err)
	}
	return cfg, nil
}

func (cfg *Config) SetupGetPairAddresses() error {
	// Initialize PairAddresses map
	cfg.PairAddresses = make(map[common.Address]Tokens)

	trendingTokens := cfg.TrendingTokens

	for _, token := range trendingTokens {
		tokenAddress := common.HexToAddress(token)
		var dexesFound []DecentralizedExchange

		for _, exchange := range cfg.DecentralizedExchanges {

			factoryAddress := common.HexToAddress(exchange.FactoryAddress)
			factory, err := arbitrageABI.NewUniswapV2Factory(factoryAddress, cfg.ETHClient)
			if err != nil {
				log.Fatalf("Failed to instantiate a Uniswap V2 factory contract: %v", err)
			}

			pairAddress, err := factory.GetPair(nil, cfg.MainTokenAddress, tokenAddress)
			if err != nil {
				log.Printf("Failed to get pair address: %v\n", err)
			}

			if pairAddress != common.HexToAddress("0x0000000000000000000000000000000000000000") {
				dexesFound = append(dexesFound, exchange)
			}
		}

		if len(dexesFound) >= 2 {
			cfg.PairAddresses[tokenAddress] = Tokens{
				MainToken: Token{
					Address: cfg.MainTokenAddress,
				},
				ExternalToken: Token{
					Address: tokenAddress,
				},
				DecentralizedExchanges: dexesFound,
			}
		}

	}

	return nil
}

func (cfg *Config) WatchSwap() error {
	client, err := ethclient.Dial(cfg.WSSRPCURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	log.Println("Subscribing to DEX")

	// Filter query for multiple router addresses
	var addresses []common.Address
	for pairAddress, _ := range cfg.PairAddresses {
		addresses = append(addresses, pairAddress)
	}

	query := ethereum.FilterQuery{
		Addresses: addresses,
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to filter logs: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			// Attempt to resubscribe after a delay
			time.Sleep(10 * time.Second)
			sub, err = client.SubscribeFilterLogs(context.Background(), query, logs)
			if err != nil {
				log.Fatalf("Failed to resubscribe: %v", err)
			}
		case vLog := <-logs:
			cfg.HandleSwapLog(vLog)
		}
	}
}

func (cfg *Config) HandleSwapLog(vLog types.Log) {
	var threshold = big.NewInt(5000)
	// Define the print amount as 20 tokens with 18 decimals
	tokenDecimals := 6
	printAmount := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokenDecimals)), nil)
	printAmount.Mul(printAmount, big.NewInt(200)) // 20 * 10^18

	pairAddress := cfg.PairAddresses[vLog.Address]

	startSwapAddress := pairAddress.DecentralizedExchanges[1].RouterV2
	externalTokenAddress := pairAddress.ExternalToken.Address

	key1, err := cfg.keccak256(pairAddress.DecentralizedExchanges[1].MKey)
	if err != nil {
		log.Println("An error occurred checking arbitrage opportunity", err)
		return
	}

	for i, dex := range pairAddress.DecentralizedExchanges {
		if i == 1 {
			continue
		}
		key2, err := cfg.keccak256(dex.MKey)
		if err != nil {
			log.Println("An error occurred checking arbitrage opportunity", err)
			return
		}

		isProfitable, err := cfg.Arbitrage.CheckProfitability(nil, key1, key2, externalTokenAddress, printAmount, threshold)
		if err != nil {
			log.Println("An error occurred checking arbitrage opportunity", err)
			return
		}

		if isProfitable.IsProfitable {
			log.Println("The external token address ", externalTokenAddress, "The  TXHash", vLog.TxHash)
			log.Println("Is profitable ", isProfitable.IsProfitable, "The direction ", isProfitable.Direction, " The profit ", isProfitable.PercentageProfit, " The log number", vLog.BlockNumber)
			if isProfitable.Direction == "ATOB" {
				log.Println("the ATOB direction", common.HexToAddress(startSwapAddress), " ", common.HexToAddress(dex.RouterV2))
				input := arbitrageABI.FlashLoanArbitrageMakeInput{
					key1, key2, externalTokenAddress, printAmount, big.NewInt(5),
				}
				makeFlashLoanTX, err := cfg.Arbitrage.Milking(cfg.Authentication, []arbitrageABI.FlashLoanArbitrageMakeInput{input})
				if err != nil {
					log.Println("Error executing flash loan ", err)
					return
				}
				log.Println("The transaction hash ", makeFlashLoanTX.Hash())

			} else if isProfitable.Direction == "BTOA" {
				log.Println("the BTOA direction", common.HexToAddress(dex.RouterV2), " ", common.HexToAddress(startSwapAddress))

				input := arbitrageABI.FlashLoanArbitrageMakeInput{
					key2, key1, externalTokenAddress, printAmount, big.NewInt(5),
				}
				makeFlashLoanTX, err := cfg.Arbitrage.Milking(cfg.Authentication, []arbitrageABI.FlashLoanArbitrageMakeInput{input})
				if err != nil {
					log.Println("Error executing flash loan ", err)
					return
				}
				log.Println("The transaction hash ", makeFlashLoanTX.Hash())
			}
		}

	}

}

func (cfg *Config) keccak256(input string) ([32]byte, error) {
	// Compute the Keccak256 hash
	hash := crypto.Keccak256([]byte(input))

	// Convert hash to fixed-size array of 32 bytes
	var hashBytes [32]byte
	copy(hashBytes[:], hash)

	return hashBytes, nil
}
