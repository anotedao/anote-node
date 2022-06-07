package main

import (
	"log"

	"github.com/anonutopia/gowaves"
)

func balance() int {
	a, err := gowaves.WNC.Addresses()
	ar := *a

	abr, err := gowaves.WNC.AddressesBalance(ar[0])
	if err != nil {
		log.Println(err.Error())
	}
	return abr.Balance
}

func initWaves() {
	gowaves.WNC.Host = "http://localhost"
	gowaves.WNC.Port = 6869
	gowaves.WNC.ApiKey = conf.ApiKey
}
