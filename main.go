package main

import (
	"os"

	"github.com/bencolon/justcoin/helpers"
	"github.com/bencolon/justcoin/tasks"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Justcoin"
	app.Version = "0.1.0"
	app.Usage = "List Justcoin.com markets values"

	app.Commands = []cli.Command{
		{
			Name:      "setup",
			ShortName: "s",
			Usage:     "Setup your justcoin.com api key",
			Action: func(c *cli.Context) {
				tasks.Setup()
			},
		},
		{
			Name:      "markets",
			ShortName: "l",
			Usage:     "List all the justcoin.com markets",
			Action: func(c *cli.Context) {
				tasks.Markets(c)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "BTCEUR",
					Value: "1",
					Usage: "BTC vs EUR trade",
				},
				cli.StringFlag{
					Name:  "BTCLTC",
					Value: "1",
					Usage: "BTC vs LTC trade",
				},
				cli.StringFlag{
					Name:  "BTCNOK",
					Value: "1",
					Usage: "BTC vs NOK trade",
				},
				cli.StringFlag{
					Name:  "BTCSTR",
					Value: "1",
					Usage: "BTC vs STR trade",
				},
				cli.StringFlag{
					Name:  "BTCUSD",
					Value: "1",
					Usage: "BTC vs USD trade",
				},
				cli.StringFlag{
					Name:  "BTCXRP",
					Value: "1",
					Usage: "BTC vs XRP trade",
				},
			},
		},
		{
			Name:      "wallet",
			ShortName: "w",
			Usage:     "Display your wallet amounts depending the options",
			Action: func(c *cli.Context) {
				tasks.Wallet(c)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "BTC",
					Value: "amount",
					Usage: "BTC amount [at least one crypto amount is mandatory]",
				},
				cli.StringFlag{
					Name:  "LTC",
					Value: "amount",
					Usage: "LTC amount",
				},
				cli.StringFlag{
					Name:  "STR",
					Value: "amount",
					Usage: "STR amount",
				},
				cli.StringFlag{
					Name:  "XRP",
					Value: "amount",
					Usage: "XRP amount",
				},
				cli.StringFlag{
					Name:  "curr",
					Value: "currency",
					Usage: "Wallet curency : EUR (default) or USD or NOK) [mandatory]",
				},
			},
		},
	}

	app.Action = func(c *cli.Context) {
		listCommand, _ := helpers.CliCommands(app.Commands).Find("markets")
		listCommand.Run(c)
	}

	app.Run(os.Args)
}
