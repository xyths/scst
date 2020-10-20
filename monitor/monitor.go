package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/seroclient"
	"github.com/shopspring/decimal"
	"github.com/xyths/scst/coral/swap"
	"log"
	"math/big"
	"os"
)

type Config struct {
	Url     string // like "http://127.0.0.1:8545"
	Address string
	Pairs   []Pair
	Output  string
}
type Pair struct {
	A         string
	B         string
	AmountIn  float64 `json:"amountIn"`
	AmountOut float64 `json:"amountOut,omitempty"`
}

func Monitor(ctx context.Context, config Config) error {
	conn, err := seroclient.Dial(config.Url)
	if err != nil {
		return err
	}
	addr := common.Base58ToAddress(config.Address)
	ex, err := swap.NewSwapExchange(addr, conn)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(config.Output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, pair := range config.Pairs {
		if err := onePair(ex, &pair); err != nil {
			log.Printf("%s-%s error: %s", pair.A, pair.B, err)
			continue
		}
		b2, err2 := json.Marshal(pair)
		if err2 != nil {
			log.Print(err2)
			continue
		}
		_, err2 = fmt.Fprintf(f, "%s\n", string(b2))
		if err2 != nil {
			log.Print(err2)
			continue
		}
	}
	return nil
}

func onePair(ex *swap.SwapExchange, pair *Pair) error {
	log.Printf("A: %s, B: %s", pair.A, pair.B)
	var ba, bb [32]byte
	copy(ba[:], []byte(pair.A))
	copy(bb[:], []byte(pair.B))
	key, err := ex.HashKey(nil, ba, bb)
	if err != nil {
		return err
	}
	has, err := ex.HasPair(nil, key)
	if err != nil {
		return err
	}
	if !has {
		log.Printf("no %s-%s pair", pair.A, pair.B)
		return nil
	}
	p, err := ex.PairInfo(nil, key)
	if err != nil {
		return err
	}
	if p.TokenA != ba {
		log.Printf("%s-%s should be %s-%s", pair.A, pair.B, pair.B, pair.A)
		pair.A, pair.B = pair.B, pair.A
	}
	amountIn := big.NewInt(int64(pair.AmountIn * 1e18))
	amountOut, err := ex.EstimateSwap(nil, key, p.TokenA, amountIn)
	if err != nil {
		return err
	}
	o, _ := decimal.NewFromBigInt(amountOut, 0).Div(decimal.NewFromFloat(1e18)).Float64()
	pair.AmountOut = o
	return nil
}
