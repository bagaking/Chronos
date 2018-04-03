package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin/json"
)

type ConfigEntry struct {
	Workername string `json:"workername"`
	Srcpath    string `json:"srcpath"`
	Timespan   string `json:"timespan"`
}

type Config struct {
	Version string        `json:"version"`
	Workers []ConfigEntry `json:"workers"`
}

func loadConfig(path string) *Config {
	fmt.Println("start load conf from " + path)
	srcData, _ := ioutil.ReadFile(path)
	fmt.Println(string(srcData[:]))

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
