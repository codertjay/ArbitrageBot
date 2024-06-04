package newConfiguration

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
)

// ConfigInterface definition remains unchanged

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
		},
		//{
		//	Name:           "SUSHISWAP",
		//	RouterV2:       "0x1b02da8cb0d097eb8d57a175b88c7d8b47997506",
		//	FactoryAddress: "0xc35dadb65012ec5796536bd9864ed8773abc74c4",
		//},
		//{
		//	Name:           "DFYN",
		//	RouterV2:       "0xA102072A4C07F06EC3B4900FDC4C7B80b6c57429",
		//	FactoryAddress: "0xE7Fb3e833eFE5F9c441105EB65Ef8b261266423B",
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
	return cfg
}

func (cfg *Config) SetupHelper() *Config {
	cfg.Helper = &Helper{}
	return cfg
}

func (cfg *Config) LoadDexTokensFromFile() (*Config, error) {
	file, err := os.Open("tokens.json")
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	err = json.Unmarshal(bytes, &cfg.Tokens)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config file: %v", err)
	}
	return cfg, nil
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
	for _, dex := range cfg.DecentralizedExchanges {

		addresses = append(addresses, common.HexToAddress(dex.RouterV2))
	}

	query := ethereum.FilterQuery{
		Addresses: addresses,
	}

	log.Println(query.Addresses)

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to filter logs: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case vLog := <-logs:
			cfg.HandleSwapLog(vLog)
		}
	}
}

func (cfg *Config) HandleSwapLog(vLog types.Log) {
	fmt.Println("Swap event detected:")
	fmt.Println("Address: ", vLog.Address.Hex())
	fmt.Println("Topics: ", vLog.Topics)
	fmt.Println("Data: ", vLog.Data)

	// Decode the log event here
	// The event signature for Uniswap V2 "Swap" event is:
	// Swap(address indexed sender, uint amount0In, uint amount1In, uint amount0Out, uint amount1Out, address indexed to)
	// Adjust the event signature based on the DEX you're monitoring

	swapEventSignature := []byte("Swap(address,uint256,uint256,uint256,uint256,address)")
	swapEventHash := crypto.Keccak256Hash(swapEventSignature)

	if vLog.Topics[0] == swapEventHash {
		//var sender, to common.Address
		//var amount0In, amount1In, amount0Out, amount1Out *big.Int

		//err := swapABI.UnpackIntoInterface(&swapEvent, "Swap", vLog.Data)
		//if err != nil {
		//	log.Fatalf("Failed to unpack log data: %v", err)
		//}
		//
		//sender = common.HexToAddress(vLog.Topics[1].Hex())
		//to = common.HexToAddress(vLog.Topics[2].Hex())
		//amount0In = swapEvent.Amount0In
		//amount1In = swapEvent.Amount1In
		//amount0Out = swapEvent.Amount0Out
		//amount1Out = swapEvent.Amount1Out
		//
		//fmt.Printf("Sender: %s, To: %s, Amount0In: %s, Amount1In: %s, Amount0Out: %s, Amount1Out: %s\n",
		//	sender.Hex(), to.Hex(), amount0In.String(), amount1In.String(), amount0Out.String(), amount1Out.String())
	}
}

func main() {
	cfg := &Config{}
	_, err := cfg.Setup()
	if err != nil {
		log.Fatalf("Failed to set up config: %v", err)
	}
}
