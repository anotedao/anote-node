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

	NodeAddress = getAddress()

	fmt.Printf("Node Address: %s\n", NodeAddress)

	if len(os.Args) == 2 {
		OwnerAddress = os.Args[1]

		fmt.Printf("Owner Address: %s\n", OwnerAddress)
		fmt.Println("Installing Anote Node... Please wait!")

		setScript()

		callScript()

		fmt.Println("Anote Node installation is now done.")
	}
}
