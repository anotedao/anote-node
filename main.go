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

	initSeedFile()

	fmt.Printf("Node Address: %s\n", NodeAddress)
	fmt.Printf("Owner Address: %s\n", OwnerAddress)

	setup := flag.Bool("setup", false, "Setup your Anote Node.")
	flag.Parse()

	if *setup {
		ping()

		waitForAnotes()

		setScript()

		waitForScript()

		callScript()
	}
}
