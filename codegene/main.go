package main

import (
	"codegene/action"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "format",
			Aliases: []string{"ft"},
			Value:   "py",
			Usage:   "file format, support: py, cpp",
		},
		&cli.StringFlag{
			Name:    "filename",
			Value:   "default",
			Usage:   "output file name",
			Aliases: []string{"fn"},
		},
	}

	app.Action = action.CodeGene
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
