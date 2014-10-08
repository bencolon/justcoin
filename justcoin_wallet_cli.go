package main

import (
	"os"

	"github.com/bencolon/justcoin_wallet_cli/helpers"
	"github.com/bencolon/justcoin_wallet_cli/tasks"

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
			Usage:     "Setup your api key",
			Action: func(c *cli.Context) {
				tasks.Setup()
			},
		},
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "List all the markets",
			Action: func(c *cli.Context) {
				tasks.List()
			},
		},
	}

	app.Action = func(c *cli.Context) {
		listCommand, _ := helpers.CliCommands(app.Commands).Find("list")
		listCommand.Run(c)
	}

	app.Run(os.Args)
}
