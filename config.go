package main

import (
	"encoding/json"
	"log"
	"os"
)

// Config struct holds all our configuration
type Config struct {
	ApiKey string `json:"api_key"`
}

// Load method loads configuration file to Config struct
func (c *Config) load(configFile string) {
	file, err := os.Open(configFile)

	if err != nil {
		log.Println(err)
		// logTelegram(err.Error())
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&c)

	if err != nil {
		log.Println(err)
		// logTelegram(err.Error())
	}
}

func initConfig() *Config {
	c := &Config{}
	c.load("config.json")
	return c
}
