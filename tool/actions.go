package tool

import (
	"github.com/glstr/futty_golang/logger"
	"github.com/urfave/cli/v2"
)

func ActionHttpServer(ctx *cli.Context) error {
	if err := initModule(); err != nil {
		panic(err)
	}
	logger.Notice("start http service")
	if err := StartHttpServer(); err != nil {
		panic(err)
	}

	return nil
}
