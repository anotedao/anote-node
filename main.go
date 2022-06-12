package main

import (
	"fmt"
	"log"
)

var conf *Config

var m *Monitor

var NodeAddress string

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	initWaves()

	if checkFlags() {
		m = initMonitor()

		fmt.Println("Shutting down gracefully.")
	}
}
