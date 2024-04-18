package main

import "context"

type DEXInterface interface {
	GetPairs() ([]Pair, error)
}

type Pair struct {
	Token0 string
	Token1 string
}

type DEX struct {
}

func (s *DEX) GetPairs(ctx context.Context, v2Factory string) ([]Pair, error) {
	return nil, nil
}

func main() {

}
