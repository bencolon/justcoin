package tasks

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Setup() {
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
}
