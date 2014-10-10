package helpers

import (
	"errors"
	"strconv"

	"github.com/bencolon/justcoin/market"

	"github.com/codegangsta/cli"
)

type CliCommands []cli.Command

func (commands CliCommands) Find(name string) (cli.Command, error) {
	for _, v := range commands {
		if v.Name == name {
			return v, nil
		}
	}
	return cli.Command{}, errors.New("Command not found")
}

func ArgumentsCount(context *cli.Context) (count int) {
	markets := []string{"BTCEUR", "BTCLTC", "BTCNOK", "BTCSTR", "BTCUSD", "BTCXRP"}
	for _, v := range markets {
		if context.IsSet(v) {
			count++
		}
	}
	return
}

func TrendColor(market market.DataStruct) (color string) {
	last, _ := strconv.ParseFloat(market.Last, 32)
	high, _ := strconv.ParseFloat(market.High, 32)
	low, _ := strconv.ParseFloat(market.Low, 32)

	trend := low + ((high - low) / 2)
	if trend > last {
		return "red"
	} else {
		return "green"
	}
}
