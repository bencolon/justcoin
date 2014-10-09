package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bencolon/justcoin/market"
)

func ReadMarkets() (markets []market.DataStruct) {
	key, err := ioutil.ReadFile(os.Getenv("HOME") + "/.justcoin")

	if err != nil || len(key) == 0 {
		needSetup()
	} else {
		resp := getJsonData(string(key))
		defer resp.Body.Close()

		body := readBody(resp)
		markets = extractJson(body)
	}

	return
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

func extractJson(body []byte) (markets []market.DataStruct) {
	err := json.Unmarshal(body, &markets)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	return
}
