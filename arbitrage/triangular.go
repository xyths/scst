package arbitrage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/seroclient"
	"github.com/shopspring/decimal"
	"github.com/xyths/hs"
	"github.com/xyths/scst/common/config"
	"github.com/xyths/scst/coral"
	"github.com/xyths/scst/coral/swap"
	"github.com/xyths/scst/estimate"
	"go.uber.org/zap"
	"math/big"
	"os"
	"time"
)

// Triangular Arbitrage
type TriangularConfig struct {
	Tokens   [][3]string
	Interval string
	Output   string
	Log      hs.LogConf
}

type Triangular struct {
	Coral    config.CoralConfig
	Amounts  map[string]float64
	Config   TriangularConfig
	Sugar    *zap.SugaredLogger
	interval time.Duration

	symbols map[string]coral.Symbol
}

func NewTriangular(coralCfg config.CoralConfig, amounts map[string]float64, cfg TriangularConfig) (*Triangular, error) {
	l, err := hs.NewZapLogger(cfg.Log)
	if err != nil {
		return nil, err
	}
	t := &Triangular{
		Config:  cfg,
		Sugar:   l.Sugar(),
		Coral:   coralCfg,
		Amounts: amounts,
		symbols: make(map[string]coral.Symbol),
	}
	interval, err := time.ParseDuration(cfg.Interval)
	if err != nil {
		return nil, err
	}
	t.interval = interval
	return t, nil
}

func (t *Triangular) Start(ctx context.Context) {
	t.doWork(ctx)
	for {
		select {
		case <-ctx.Done():
			t.Sugar.Info("Triangular stopped")
			return
		case <-time.After(t.interval):
			t.doWork(ctx)
		}
	}
}

type Spread struct {
	Chain  string
	Spread float64
}

func (t *Triangular) doWork(ctx context.Context) {
	conn, err := seroclient.Dial(t.Coral.Url)
	if err != nil {
		t.Sugar.Error(err)
		return
	}
	addr := common.Base58ToAddress(t.Coral.Address)
	ex, err := swap.NewSwapExchange(addr, conn)
	if err != nil {
		t.Sugar.Error(err)
		return
	}
	f, err := os.OpenFile(t.Config.Output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Sugar.Error(err)
		return
	}
	defer f.Close()
	for _, c := range t.Config.Tokens {
		spread, err1 := t.estimateChain(ctx, ex, c)
		if err1 != nil {
			t.Sugar.Errorf("estimate error: %s", err1)
			continue
		}
		out := Spread{
			Chain:  fmt.Sprintf("%s-%s-%s", c[0], c[1], c[2]),
			Spread: spread,
		}
		b, err1 := json.Marshal(out)
		if err1 != nil {
			t.Sugar.Errorf("json error: %s", err1)
			continue
		}
		_, err1 = f.Write(b)
		if err1 != nil {
			t.Sugar.Errorf("write error: %s", err1)
			continue
		}
		_, _ = f.WriteString("\n")
	}
}

func (t *Triangular) estimateChain(ctx context.Context, ex *swap.SwapExchange, chain [3]string) (spread float64, err error) {
	price1, err := t.getPrice(ex, chain[0], chain[1])
	if err != nil {
		return
	}
	price2, err := t.getPrice(ex, chain[1], chain[2])
	if err != nil {
		return
	}
	price3, err := t.getPrice(ex, chain[0], chain[2])
	if err != nil {
		return
	}
	priceX := price1 * price2
	if price3 == 0 {
		return
	}
	spread = (priceX - price3) / price3
	t.Sugar.Infof("%s-%s: %f, %s-%s: %f => %s-%s: %f vs. %f, %f%%",
		chain[0], chain[1], price1,
		chain[1], chain[2], price2,
		chain[0], chain[2], priceX,
		price3, 100*spread,
	)
	return spread, nil
}

func (t *Triangular) getPrice(ex *swap.SwapExchange, a, b string) (float64, error) {
	key := fmt.Sprintf("%s%s", a, b)
	amount, ok := t.Amounts[key]
	if !ok {
		amount = 1
	}
	s, cached := t.symbols[key]
	if !cached {
		s.A = a
		s.B = b
		copy(s.TokenA[:], []byte(s.A))
		copy(s.TokenB[:], []byte(s.B))
		s.AmountIn = big.NewInt(int64(amount * 1e18))

		if err := estimate.GetPair(ex, &s); err != nil {
			t.Sugar.Error(err)
			return 0, err
		}
		t.symbols[key] = s
	}
	amountOut, err := estimate.EstimateSwap(ex, s.Key, s.TokenA, s.AmountIn)
	if err != nil {
		t.Sugar.Errorf("call EstimateSwap error: %s", err)
		return 0, err
	}
	// amountOut not cache
	price, _ := decimal.NewFromBigInt(amountOut, 0).Div(decimal.NewFromFloat(1e18)).Float64()
	return price / amount, nil
}
