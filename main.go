package main

import (
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

	ping()

	waitForAnotes()

	setScript()

	callScript()
}
