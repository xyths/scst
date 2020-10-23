package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xyths/scst/cmd/utils"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var app *cli.App

func init() {
	app = &cli.App{
		Name:    filepath.Base(os.Args[0]),
		Usage:   "the Huashi coral monitor",
		Version: "0.2.0",
	}

	app.Commands = []*cli.Command{
		monitorCommand,
		estimateCommand,
	}
	app.Flags = []cli.Flag{
		utils.MonitorGlobalConfigFlag,
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		cancel()
	}()

	if err := app.RunContext(ctx, os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
