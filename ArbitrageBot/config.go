package main

// Sushi swap deployed address doc https://dev.sushi.com/docs/Products/Classic%20AMM/Deployment%20Addresses

type DecentralizedExchange struct {
	Name           string
	RouterV2       string
	FactoryAddress string
}

type Config struct {
	DecentralizedExchanges []DecentralizedExchange
	Tokens                 []string
	RPCUrl                 string
}

type ExchangeInterface interface {
	GetRouterV2() string
	GetFactoryAddress() string
}

func (dex DecentralizedExchange) GetRouterV2() string {
	return dex.RouterV2
}

func (dex DecentralizedExchange) GetFactoryAddress() string {
	return dex.FactoryAddress
}

func main() {
	// Example usage
	config := Config{
		DecentralizedExchanges: []DecentralizedExchange{
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
		},
		Tokens: []string{"Token1", "Token2", "Token3"},
		RPCUrl: "https://127.0.0.1:8545",
	}

	// Accessing config values
	println("RPC URL:", config.RPCUrl)
	println("Tokens:", config.Tokens)

	// Accessing decentralized exchanges
	for _, dex := range config.DecentralizedExchanges {
		println("RouterV2:", dex.GetRouterV2())
		println("Factory Address:", dex.GetFactoryAddress())
	}
}
