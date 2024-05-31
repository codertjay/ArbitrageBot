package configurations

import (
	"log"
	"math/big"

	arbitrageABI "ArbitrageBot/ArbitrageBot/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type HelperInterface interface {
	GetPair(factoryAddress, token0, token1 common.Address) (common.Address, error)
	GetReserves(pairAddress common.Address) (*Reserves, error)
	CalculatePrice(reserves *Reserves) *big.Float
	CalculatePriceDifference(uPrice, sPrice *big.Float) *big.Float
	CheckArbitrageOpportunity(uPrice, sPrice *big.Float, threshold float64) (bool, string)
}

type Helper struct {
	client *ethclient.Client
}

func NewHelper(client *ethclient.Client) HelperInterface {
	return &Helper{client: client}
}

// GetPair retrieves the pair address for a given token0 and token1
func (h *Helper) GetPair(factoryAddress, token0, token1 common.Address) (common.Address, error) {
	// Instantiate a contract object for the Uniswap V2 Factory
	factory, err := arbitrageABI.NewUniswapV2Factory(factoryAddress, h.client)
	if err != nil {
		return common.Address{}, err
	}

	// Retrieve the pair address
	pairAddress, err := factory.GetPair(nil, token0, token1)
	if err != nil {
		return common.Address{}, err
	}

	return pairAddress, nil
}

// Reserves struct to hold reserves of the pair
type Reserves struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}

// GetReserves retrieves reserves of the pair
func (h *Helper) GetReserves(pairAddress common.Address) (*Reserves, error) {

	pair, err := arbitrageABI.NewUniswapV2Pair(pairAddress, h.client)
	if err != nil {
		return nil, err
	}

	// Retrieve reserves of the pair
	reserves, err := pair.GetReserves(nil)
	if err != nil {
		return nil, err
	}

	return &Reserves{
		Reserve0: new(big.Int).Set(reserves.Reserve0),
		Reserve1: new(big.Int).Set(reserves.Reserve1),
	}, nil
}

// CalculatePrice calculates the price of the token pair
func (h *Helper) CalculatePrice(reserves *Reserves) *big.Float {
	// Calculate price (reserve1 / reserve0)
	price := new(big.Float).Quo(
		new(big.Float).SetInt(reserves.Reserve1),
		new(big.Float).SetInt(reserves.Reserve0),
	)
	return price
}

// CalculatePriceDifference calculates the percentage difference between two prices
func (h *Helper) CalculatePriceDifference(uPrice, sPrice *big.Float) *big.Float {
	// Calculate price difference percentage ((uPrice - sPrice) / sPrice) * 100
	difference := new(big.Float).Sub(uPrice, sPrice)
	difference.Quo(difference, sPrice)
	difference.Mul(difference, big.NewFloat(100))
	return difference
}

// CheckArbitrageOpportunity checks if there's a profitable arbitrage opportunity
func (h *Helper) CheckArbitrageOpportunity(priceA, priceB *big.Float, threshold float64) (bool, string) {
	differenceAtoB := h.CalculatePriceDifference(priceA, priceB)
	differenceBtoA := h.CalculatePriceDifference(priceB, priceA)
	absDifferenceAtoB := new(big.Float).Abs(differenceAtoB)
	absDifferenceBtoA := new(big.Float).Abs(differenceBtoA)
	thresholdBig := big.NewFloat(threshold)

	if absDifferenceAtoB.Cmp(thresholdBig) >= 0 {
		log.Print("The price difference is ", absDifferenceAtoB)
		return true, "AtoB"
	}
	if absDifferenceBtoA.Cmp(thresholdBig) >= 0 {
		log.Print("The price difference is ", absDifferenceBtoA)

		return true, "BtoA"
	}
	return false, ""
}
