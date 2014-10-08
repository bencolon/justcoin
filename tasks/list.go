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
	key, err := ioutil.ReadFile(os.Getenv("HOME") + "/.justcoin")

	if err != nil || len(key) == 0 {
		needSetup()
	} else {
		resp := getJsonData(string(key))
		defer resp.Body.Close()

		body := readBody(resp)
		var markets []Market = extractJson(body)

		displayMarkets(markets)
	}
}

func needSetup() {
	log.Println("Please setup your Justcoin api key first")
	log.Println("Run `justcoin setup`")
	os.Exit(0)
}

func getJsonData(key string) (resp *http.Response) {
	resp, err := http.Get("https://justcoin.com/api/v1/markets?key=" + key)

	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	} else if resp.StatusCode != 200 {
		log.Fatal("HTTP Error " + resp.Status)
		os.Exit(0)
	}

	return
}

func readBody(resp *http.Response) (body []byte) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	return
}

func extractJson(body []byte) (markets []Market) {
	err := json.Unmarshal(body, &markets)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	return
}

func displayMarkets(markets []Market) {
	for _, v := range markets {
		price, _ := strconv.ParseFloat(v.Last, 16)
		fmt.Printf("%s %.2f\n", v.Id, price)
	}
}
