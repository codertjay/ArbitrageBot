package configurations

import (
	arbitrageABI "ArbitrageBot/ArbitrageBot/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

// Sushi swap deployed address doc https://dev.sushi.com/docs/Products/Classic%20AMM/Deployment%20Addresses

type ConfigInterface interface {
	SetupRPCURL() *Config
	SetupDecentralizedExchange() *Config
	SetupMainTokenAddress() *Config
	SetupTokens() (*Config, error)
	Load() (*Config, error)
}

func NewConfig() ConfigInterface {
	return &Config{}
}

type Config struct {
	DecentralizedExchanges []DecentralizedExchange
	Tokens                 []Tokens
	RPCUrl                 string
	MainTokenAddress       common.Address
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

func (cfg *Config) SetupRPCURL() *Config {
	cfg.RPCUrl = "https://bsc-dataseed2.defibit.io/" // rpc-url needed
	return cfg
}

func (cfg *Config) SetupDecentralizedExchange() *Config {
	cfg.DecentralizedExchanges = []DecentralizedExchange{
		{
			Name:           "UNISWAP",
			RouterV2:       "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
			FactoryAddress: "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f",
		},
		{
			Name:           "SUSHISWAP",
			RouterV2:       "0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F",
			FactoryAddress: "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac",
		},
	}
	return cfg
}

func (cfg *Config) SetupMainTokenAddress() *Config {
	cfg.MainTokenAddress = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	return cfg
}

type DexTokens struct {
	MainToken     common.Address
	ExternalToken common.Address
	Exist         bool
}

func (cfg *Config) SetupTokens() (*Config, error) {
	var dexTokens []DexTokens

	client, err := rpc.Dial(cfg.RPCUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer client.Close()
	ethClient := ethclient.NewClient(client)

	for counter, exchange := range cfg.DecentralizedExchanges {

		// Create a new instance of the Uniswap V2 Factory contract
		factoryAddress := common.HexToAddress(exchange.FactoryAddress)
		factory, err := arbitrageABI.NewUniswapV2Factory(factoryAddress, ethClient)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		// Get the total number of pairs
		totalPairs, err := factory.AllPairsLength(nil)
		if err != nil {
			log.Fatal(err)
		}

		// Iterate over all pairs and check if they contain WBNB
		for i := big.NewInt(0); i.Cmp(totalPairs) < 0; i.Add(i, big.NewInt(1)) {
			pairAddress, err := factory.AllPairs(nil, i)
			if err != nil {
				log.Println("Error fetching pair address:", err)
				continue
			}

			log.Println(pairAddress)

			// Instantiate UniswapV2Pair contract
			pairContract, err := arbitrageABI.NewUniswapV2Pair(pairAddress, ethClient)
			if err != nil {
				log.Println("Error creating pair contract:", err)
				continue
			}

			// Get tokens in the pair
			token0, err := pairContract.Token0(nil)
			if err != nil {
				log.Println("Error fetching token0:", err)
				continue
			}
			token1, err := pairContract.Token1(nil)
			if err != nil {
				log.Println("Error fetching token1:", err)
				continue
			}

			if token0 != cfg.MainTokenAddress || token1 != cfg.MainTokenAddress {
				continue
			}

			if counter == 0 {
				dexToken := DexTokens{
					MainToken:     token0,
					ExternalToken: token1,
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

		}

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

	return cfg, nil
}

func (cfg *Config) Load() (*Config, error) {

	cfg = cfg.SetupRPCURL()
	cfg = cfg.SetupDecentralizedExchange()
	cfg = cfg.SetupMainTokenAddress()
	cfg, err := cfg.SetupTokens()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
