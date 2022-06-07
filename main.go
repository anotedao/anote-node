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

	m = initMonitor()

	fmt.Println("Done.")
}
