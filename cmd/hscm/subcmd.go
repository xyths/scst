package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xyths/hs"
	"github.com/xyths/scst/arbitrage"
	"github.com/xyths/scst/cmd/utils"
	"github.com/xyths/scst/common/config"
	"github.com/xyths/scst/monitor"
)

var (
	monitorCommand = &cli.Command{
		Action: monitorAction,
		Name:   "monitor",
		Usage:  "Monitor the Coral exchange corswap price",
		Flags: []cli.Flag{
			utils.ConfigFlag,
		},
	}
	estimateCommand = &cli.Command{
		Action: estimateAction,
		Name:   "estimate",
		Usage:  "Estimate the price spread between Coral exchange symbols",
		Flags: []cli.Flag{
			utils.ConfigFlag,
		},
	}
)

func monitorAction(ctx *cli.Context) error {
	globalConfigFilename := ctx.String(utils.MonitorGlobalConfigFlag.Name)
	global := config.MonitorGlobalConfig{}
	if err := hs.ParseJsonConfig(globalConfigFilename, &global); err != nil {
		return err
	}
	configFile := ctx.String(utils.ConfigFlag.Name)
	cfg := monitor.Config{}
	if err := hs.ParseJsonConfig(configFile, &cfg); err != nil {
		return err
	}
	if err := monitor.Monitor(ctx.Context,global.Coral, getAmounts(global.Pairs), cfg); err != nil {
		return err
	}
	return nil
}

func estimateAction(ctx *cli.Context) error {
	globalConfigFilename := ctx.String(utils.MonitorGlobalConfigFlag.Name)
	global := config.MonitorGlobalConfig{}
	if err := hs.ParseJsonConfig(globalConfigFilename, &global); err != nil {
		return err
	}
	triangularFilename := ctx.String(utils.ConfigFlag.Name)
	triangular := arbitrage.TriangularConfig{}
	if err := hs.ParseJsonConfig(triangularFilename, &triangular); err != nil {
		return  err
	}
	t, err := arbitrage.NewTriangular(global.Coral, getAmounts(global.Pairs), triangular)
	if err != nil {
		return err
	}
	t.Start(ctx.Context)
	return nil
}

func getAmounts(pairs []config.PairConfig) map[string]float64 {
	m := make(map[string]float64)
	for _, p := range pairs {
		k := fmt.Sprintf("%s%s", p.A, p.B)
		m[k] = p.AmountIn
	}
	return m
}
