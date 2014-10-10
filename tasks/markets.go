package tasks

import (
	"fmt"
	"strconv"

	"github.com/bencolon/justcoin/helpers"
	"github.com/bencolon/justcoin/market"

	"github.com/codegangsta/cli"
)

func Markets(context *cli.Context) {
	markets := helpers.ReadMarkets()
	displayMarkets(markets, context)
}

func displayMarkets(markets []market.DataStruct, context *cli.Context) {
	for _, v := range markets {
		if helpers.ArgumentsCount(context) == 0 || context.IsSet(v.Id) {
			price, _ := strconv.ParseFloat(v.Last, 32)

			if context.IsSet("trend") {
				color := helpers.TrendColor(v)

				switch color {
				case "green":
					fmt.Printf("%s = \033[32m%.2f\033[0m\n", v.Id, price)
				case "red":
					fmt.Printf("%s = \033[31m%.2f\033[0m\n", v.Id, price)
				}
			} else {
				fmt.Printf("%s = %.2f\n", v.Id, price)
			}
		}
	}
}
