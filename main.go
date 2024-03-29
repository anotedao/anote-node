package main

import (
	"flag"
	"fmt"
	"log"
)

var NodeAddress string

var OwnerAddress string

var PublicKey string

var PrivateKey string

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	init := flag.Bool("init", false, "Initialize your Anote Node with secret file.")
	install := flag.String("install", "", "Install your Anote Node.")
	flag.Parse()

	if *init {
		initSecretsFile()

		fmt.Printf("Node Address: %s\n", NodeAddress)
	} else if len(*install) > 0 {
		initSecretsFile()

		OwnerAddress = getOwnerAddress(*install)

		fmt.Printf("Node Address: %s\n", NodeAddress)
		fmt.Printf("Owner Address: %s\n", OwnerAddress)
		fmt.Println("Installing Anote Node... Please wait!")

		err := setScript()

		waitForScript()

		err1 := callScript()

		if err == nil && err1 == nil {
			fmt.Println("Anote Node installation is now done.")
		} else {
			fmt.Println("Errror occured.")
		}
	} else {
		flag.Usage()
	}
}
