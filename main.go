package main

import (
	"fmt"
	"github.com/aerfio/gobc/internal"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func main() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		log.Fatal(fmt.Errorf("while creating zap logger: %w", err))
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	app := &cli.App{
		Name:   "gobc",
		Action: internal.DeleteStaleBranchesWithExitCode(sugar),
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:    "depth",
				Aliases: []string{"d"},
				Value:   5,
				Usage:   "how many directories to go down",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Value:   false,
				Usage:   "turn on verbose logging",
			},
			&cli.StringFlag{
				Name:    "remote",
				Aliases: []string{"r"},
				Value:   "origin",
				Usage:   "name of git remote",
			},
			&cli.BoolFlag{
				Name:  "version",
				Value: false,
				Usage: "prints gobc binary version",
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		logger.Panic(err.Error())
	}
}
