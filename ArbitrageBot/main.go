package main

import (
	"ArbitrageBot/ArbitrageBot/config"
	"log"
)

//go:generate abigen --abi ./out/Arbitrage.sol/Arbitrage.json --pkg arbitrageABI --type Arbitrage --out ./ArbitrageBot/abi/arbitrage.go
//go:generate abigen --abi ./out/FlashLoanArbitrage.sol/FlashLoanArbitrage.json --pkg arbitrageABI --type FlashLoanArbitrage --out ./ArbitrageBot/abi/flashLoanArbitrage.go
//go:generate abigen --abi ./out/IUniswapV2Router02.sol/IUniswapV2Router02.json --pkg arbitrageABI --type IUniswapV2Router02 --out ./ArbitrageBot/abi/iuniswapV2Router02.go
//go:generate abigen --abi ./node_modules/@uniswap/v2-core/build/UniswapV2Factory.json --pkg arbitrageABI --type UniswapV2Factory --out ./ArbitrageBot/abi/uniswapV2Factory.go
//go:generate abigen --abi ./node_modules/@uniswap/v2-core/build/UniswapV2Pair.json --pkg arbitrageABI --type UniswapV2Pair --out ./ArbitrageBot/abi/uniswapV2Pair.go

func main() {
	cfg := &config.Config{}
	_, err := cfg.Setup()
	if err != nil {
		log.Fatalf("Failed to set up config: %v", err)
	}
	select {}
}
