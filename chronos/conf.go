package chronos


import (
	"fmt"
	"os"

	"encoding/json"
	"github.com/bagaking/bagakit/console"
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
	console.PrintCLn(console.SDefault, console.BGYellow, console.FBlue, "start load conf from ", path)
	file, _ := os.Open(path)

	conf := Config{}

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)

	if err != nil {
		console.PrintCLn(console.SDefault, console.BGCLEAR, console.FRed, "Error:\n", err)
	} else {
		fmt.Println("Conf:\n", conf)
	}
	return &conf
}
