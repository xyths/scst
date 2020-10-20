package main

import (
	"github.com/urfave/cli/v2"
	"github.com/xyths/hs"
	"github.com/xyths/scst/cmd/utils"
	"github.com/xyths/scst/monitor"
)

var (
	monitorCommand = &cli.Command{
		Action: monitorAction,
		Name:   "monitor",
		Usage:  "Monitor the Coral exchange corswap price",
	}
)

func monitorAction(ctx *cli.Context) error {
	configFile := ctx.String(utils.ConfigFlag.Name)
	cfg := monitor.Config{}
	if err := hs.ParseJsonConfig(configFile, &cfg); err != nil {
		return err
	}
	if err := monitor.Monitor(ctx.Context, cfg); err != nil {
		return err
	}
	return nil
}
