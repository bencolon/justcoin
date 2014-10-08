package tasks

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Setup() {
	key, err := ioutil.ReadFile(os.Getenv("HOME") + "/.justcoin")

	if err != nil || len(key) == 0 {
		key := askForKey()
		saveKey(key)
	} else {
		keyPresent(key)
	}
}

func askForKey() (input string) {
	input = ""
	log.Println("Please enter your Justcoin api key :")
	for len(input) == 0 {
		fmt.Scanln(&input)
	}
	return
}

func saveKey(key string) {
	file, err := os.Create(os.Getenv("HOME") + "/.justcoin")
	defer file.Close()

	if err != nil {
		log.Printf("Unable to create config file (%s). Try again...", err)
		os.Exit(0)
	} else {
		file.WriteString(key)
		log.Println("Justcoin api key successfully saved")
	}
}

func keyPresent(key []byte) {
	log.Println("Your Justcoin api key is ready : " + string(key))
	log.Println("Run `justcoin list`")
	log.Println("To use another api key, just remove " + os.Getenv("HOME") + "/.justcoin file")
}
