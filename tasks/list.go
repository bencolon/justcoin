package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Market struct {
	Id     string
	Last   string
	High   string
	Low    string
	Bid    string
	Ask    string
	Volume string
	Scale  int
}

func List() {
	data, err := ioutil.ReadFile(os.Getenv("HOME") + "/.justcoin")
	if err != nil || len(data) == 0 {
		log.Println("Please setup your Justcoin api key first")
		log.Println("Run `justcoin setup`")
		os.Exit(0)
	} else {
		resp, err := http.Get("https://justcoin.com/api/v1/markets?key=" + string(data))
		defer resp.Body.Close()
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		if resp.StatusCode != 200 {
			log.Fatal("HTTP Error " + resp.Status)
			os.Exit(0)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}

		var markets []Market
		err = json.Unmarshal(body, &markets)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		for _, v := range markets {
			price, _ := strconv.ParseFloat(v.Last, 16)
			fmt.Printf("%s %.2f\n", v.Id, price)
		}
	}
}
