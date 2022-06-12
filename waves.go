package main

import (
	"log"
	"time"

	"github.com/anonutopia/gowaves"
)

func initWaves() {
	gowaves.WNC.Host = "http://localhost"
	gowaves.WNC.Port = 6869
	gowaves.WNC.ApiKey = conf.ApiKey
	// gowaves.WNC.Debug = true

	a, err := gowaves.WNC.Addresses()
	if err != nil {
		log.Println(err.Error())
		for err != nil {
			time.Sleep(time.Second * 10)
			a, err = gowaves.WNC.Addresses()
			if err != nil {
				log.Println(err.Error())
			}
		}
	}

	ar := *a
	NodeAddress = ar[0]
	log.Println(NodeAddress)
}
