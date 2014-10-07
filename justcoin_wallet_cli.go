package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
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

type cliCommands []cli.Command

func (commands cliCommands) Find(name string) (cli.Command, error) {
	for _, v := range commands {
		if v.Name == name {
			return v, nil
		}
	}
	return cli.Command{}, errors.New("Command not found")
}

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
				data, err := ioutil.ReadFile(os.Getenv("HOME") + "/.justcoin")
				if err != nil || len(data) == 0 {
					input := ""
					log.Println("Please enter your Justcoin api key :")
					for len(input) == 0 {
						fmt.Scanln(&input)
					}
					file, err := os.Create(os.Getenv("HOME") + "/.justcoin")
					defer file.Close()
					if err != nil {
						log.Printf("Unable to create config file (%s). Try again...", err)
						os.Exit(0)
					} else {
						file.WriteString(input)
						log.Println("Justcoin api key successfully saved")
					}
				} else {
					log.Println("Your Justcoin api key is ready : ")
					log.Println(string(data))
					log.Println("Run `justcoin list`")
					log.Println("To use another api key, just remove " + os.Getenv("HOME") + "/.justcoin file")
				}
			},
		},
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "List all the markets",
			Action: func(c *cli.Context) {
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
			},
		},
	}

	app.Action = func(c *cli.Context) {
		listCommand, _ := cliCommands(app.Commands).Find("list")
		listCommand.Run(c)
	}

	app.Run(os.Args)
}
