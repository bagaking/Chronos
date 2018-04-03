package main

import (
	"fmt"
	"os"

	"encoding/json"
)

// ConfigEntry : config of worker
type ConfigEntry struct {
	Workername string `json:"workername"`
	Srcpath    string `json:"srcpath"`
	Timespan   string `json:"timespan"`
}

// Config : global config
type Config struct {
	Version string        `json:"version"`
	Workers []ConfigEntry `json:"workers"`
}

func loadConfig(path string) *Config {
	fmt.Println("start load conf from " + path)
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	conf := Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:\n", err)
	} else {
		fmt.Println("Conf:\n", conf)
	}
	return &conf
}
