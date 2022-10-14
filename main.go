package main

import (
	"fmt"
	"log"
	"os"
)

var NodeAddress string

var OwnerAddress string

var PublicKey string

var PrivateKey string

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	initSeedFile()

	fmt.Printf("Node Address: %s\n", NodeAddress)

	// setup := flag.Bool("setup", false, "Setup your Anote Node.")
	// update := flag.Bool("update", false, "Update your Anote Node.")
	// flag.Parse()

	if len(os.Args) == 2 {
		OwnerAddress = os.Args[1]

		fmt.Printf("Owner Address: %s\n", OwnerAddress)
		fmt.Println("Installing Anote Node... Please wait!")

		ping()

		waitForAnotes()

		setScript()

		waitForScript()

		callScript()

		fmt.Println("Anote Node installation is now done.")
	}
}
