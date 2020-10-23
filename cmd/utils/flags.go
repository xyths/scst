package utils

import "github.com/urfave/cli/v2"

var (
	MonitorGlobalConfigFlag = &cli.StringFlag{
		Name:    "global",
		Aliases: []string{"g"},
		Value:   "config.json",
		Usage:   "load global configuration from `file`",
	}
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Value:   "config.json",
		Usage:   "load configuration from `file`",
	}
)
