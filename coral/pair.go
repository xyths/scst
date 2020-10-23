package coral

import "math/big"

//
type Symbol struct {
	A        string
	B        string
	TokenA   [32]byte
	TokenB   [32]byte
	AmountIn *big.Int

	Key      [32]byte
	Valid    bool
	Inverted bool

	AmountOut *big.Int
}
