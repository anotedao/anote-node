package main

import (
	"fmt"
	"log"
)

var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	fmt.Println("Done.")
}
