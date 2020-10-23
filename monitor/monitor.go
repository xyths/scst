package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/seroclient"
	"github.com/shopspring/decimal"
	"github.com/xyths/scst/common/config"
	"github.com/xyths/scst/coral/swap"
	"log"
	"math/big"
	"os"
)

type Config struct {
	Pairs  [][2]string
	Coral  config.CoralConfig
	Output string
}

func Monitor(ctx context.Context, coral config.CoralConfig, amounts map[string]float64, config Config) error {
	conn, err := seroclient.Dial(config.Coral.Url)
	if err != nil {
		return err
	}
	addr := common.Base58ToAddress(config.Coral.Address)
	ex, err := swap.NewSwapExchange(addr, conn)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(config.Output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, tokens := range config.Pairs {
		key := fmt.Sprintf("%s%s", tokens[0], tokens[1])
		amountIn, ok := amounts[key]
		if !ok {
			amountIn = 1
		}
		pair := Pair{
			A:        tokens[0],
			B:        tokens[1],
			AmountIn: amountIn,
		}
		if err := onePair(ex, &pair); err != nil {
			log.Printf("%s-%s error: %s", tokens[0], tokens[1], err)
			continue
		}
		b2, err2 := json.Marshal(tokens)
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

type Pair struct {
	A         string
	B         string
	AmountIn  float64
	AmountOut float64
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
