package main

import (
	_ "net/http/pprof"

	"os"

	"github.com/glstr/futty_golang/tool"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "snow",
		Usage:  "run services",
		Action: tool.ActionHttpServer,
		Commands: []*cli.Command{
			{
				Name:   "httpserver",
				Usage:  "start http server",
				Action: tool.ActionHttpServer,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
