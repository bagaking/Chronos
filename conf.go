package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin/json"
)

type ConfigEntry struct {
	Workername string
	Srcpath    string
	Timespan   string
}

type Config struct {
	Version string
	Workers []ConfigEntry
}

func loadConfig(path string) *Config {
	fmt.Println("start load conf from " + path)
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	fmt.Printf("decoder:\n", decoder)
	conf := Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:\n", err)
	} else {
		fmt.Println("Conf:\n", decoder)
	}
	return &conf
}
