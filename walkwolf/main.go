package main

import (
	"os"
	"walkwolf/action"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "walkwolf"
	app.Commands = []*cli.Command{
		{
			Name:        "client",
			Aliases:     []string{"c"},
			Usage:       "make a client",
			Description: "make a client to send request",
			Action:      action.Client,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "protocol",
					Aliases: []string{"p"},
				},
				&cli.StringFlag{
					Name:    "url",
					Aliases: []string{"u"},
				},
				&cli.StringFlag{
					Name:    "casename",
					Aliases: []string{"cn"},
				},
			},
		},
		{
			Name:        "walkwolf",
			Aliases:     []string{"wl"},
			Usage:       "let a walkwolf get all you need",
			Description: "a wolf is greater than a spider",
			Action:      action.Walk,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "cmd",
					Aliases: []string{"c"},
				},
				&cli.StringFlag{
					Name:    "rooturl",
					Aliases: []string{"ru"},
				},
				&cli.StringFlag{
					Name:    "rootdir",
					Aliases: []string{"rd"},
				},
				&cli.StringFlag{
					Name:    "depth",
					Aliases: []string{"d"},
				},
			},
		},
	}
	app.Run(os.Args)
}
