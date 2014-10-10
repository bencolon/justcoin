package tasks

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bencolon/justcoin/helpers"
	"github.com/bencolon/justcoin/market"

	"github.com/codegangsta/cli"
)

var Cryptos = [](string){"BTC", "LTC", "STR", "XRP"}

func Wallet(context *cli.Context) {
	markets := helpers.ReadMarkets()

	if cryptoPresent(context) && context.IsSet("curr") {
		currencies := loadCurrencies(markets)
		displayWallet(currencies, context)
	} else {
		log.Println("Incorrect Usage. Check the help with `justcoin help`.")
		os.Exit(0)
	}
}

func cryptoPresent(context *cli.Context) (present bool) {
	for _, c := range Cryptos {
		if context.IsSet(c) {
			present = true
			break
		}
	}
	return
}

func loadCurrencies(markets []market.DataStruct) (currencies map[string]float64) {
	currencies = make(map[string]float64)
	for _, v := range markets {
		currencies[v.Id], _ = strconv.ParseFloat(v.Last, 64)
	}
	return
}

func displayWallet(currencies map[string]float64, context *cli.Context) {
	total := 0.
	for _, c := range Cryptos {
		if context.IsSet(c) {
			amount, _ := strconv.ParseFloat(context.String(c), 32)

			backToBtc := 1.0
			if c != "BTC" {
				backToBtc = currencies["BTC"+c]
			}

			value := amount / backToBtc * currencies["BTC"+context.String("curr")]
			fmt.Printf("%.2f %s = %.2f %s\n", amount, c, value, context.String("curr"))

			if context.IsSet("tot") {
				total += value
			}
		}
	}

	if context.IsSet("tot") {
		fmt.Printf("TOTAL = %.2f %s\n", total, context.String("curr"))
	}
}
