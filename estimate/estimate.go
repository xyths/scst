package estimate

import (
	"github.com/xyths/scst/coral"
	"github.com/xyths/scst/coral/swap"
	"math/big"
)

func GetPair(ex *swap.SwapExchange, s *coral.Symbol) error {
	key, err := ex.HashKey(nil, s.TokenA, s.TokenB)
	if err != nil {
		return err
	}
	s.Key = key
	has, err := ex.HasPair(nil, key)
	if err != nil {
		return err
	}
	if !has {
		return nil
	} else {
		s.Valid = has
	}
	p, err := ex.PairInfo(nil, key)
	if err != nil {
		return err
	}
	if p.TokenA != s.TokenA {
		s.Inverted = true
	}
	return nil
}

func EstimateSwap(ex *swap.SwapExchange, key, tokenIn [32]byte, amountIn *big.Int) (*big.Int, error) {
	return ex.EstimateSwap(nil, key, tokenIn, amountIn)
}
