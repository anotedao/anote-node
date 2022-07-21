package main

import (
	"log"
)

var NodeAddress string

var OwnerAddress string

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	checkFlags()

	initAddresses()
}
