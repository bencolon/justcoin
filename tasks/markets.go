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
			fmt.Printf("%s = %.2f\n", v.Id, price)
		}
	}
}
