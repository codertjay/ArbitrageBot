package main

import (
	arbitrageABI "ArbitrageBot/ArbitrageBot/abi"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

func main() {
	// Connect to Ethereum client
	client, err := rpc.Dial("https://bsc-dataseed.binance.org")
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Close()

	ethClient := ethclient.NewClient(client)

	// Create a new instance of the Uniswap V2 Factory contract
	factoryAddress := common.HexToAddress("0x8909Dc15e40173Ff4699343b6eB8132c65e18eC6")
	factory, err := arbitrageABI.NewUniswapV2Factory(factoryAddress, ethClient)
	if err != nil {
		log.Fatal(err)
	}

	// Get the total number of pairs
	totalPairs, err := factory.AllPairsLength(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("totalPairs ", totalPairs)

	// Iterate over all pairs and check if they contain WBNB
	for i := big.NewInt(0); i.Cmp(totalPairs) < 0; i.Add(i, big.NewInt(1)) {
		pairAddress, err := factory.AllPairs(nil, i)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(pairAddress)

		// Instantiate UniswapV2Pair contract
		pairContract, err := arbitrageABI.NewUniswapV2Pair(pairAddress, ethClient)
		if err != nil {
			log.Fatal(err)
		}

		// Get tokens in the pair
		token0, err := pairContract.Token0(nil)
		if err != nil {
			log.Fatal(err)
		}
		token1, err := pairContract.Token1(nil)
		if err != nil {
			log.Fatal(err)
		}

		// Check if one of the tokens is WBNB
		if token0 == common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c") ||
			token1 == common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c") {
			fmt.Printf("Pair %s contains WBNB\n", pairAddress.Hex())
		}

	}
}
