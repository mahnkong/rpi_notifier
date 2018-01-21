package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port      int
	ButtonPin int
	Leds      map[string][]int
}

func NewConfigFromFile(configFile string) *Config {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	conf := Config{Port: 7777}
	json.Unmarshal(data, &conf)
	return &conf
}
