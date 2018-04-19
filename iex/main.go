package main

import (
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:  "iex",
		Usage: "iex api cli",
		Commands: []*cli.Command{
			{
				Name:    "price",
				Aliases: []string{"p"},
				Usage:   "query price",
				Action: func(c *cli.Context) error {
					getPrice()
					return nil
				},
			},
			{
				Name:    "news",
				Aliases: []string{"n"},
				Usage:   "query news",
				Action: func(c *cli.Context) error {
					getNews()
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
