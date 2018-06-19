package main

import (
	"fmt"
	"time"
	"github.com/bagaking/bagakit/console"
)

const (
	confPATH = ".kh.chronos.json"
)

func main() {
	conf := loadConfig(confPATH)
	fmt.Printf("\nConf loaded:\n %#v \n", conf)

	workerhub := &Hub{}
	fmt.Printf("\nHub loaded:\n %#v \n", workerhub)

	for _, workerCfg := range conf.Workers {
		workerhub.Insert(workerCfg.Workername, workerCfg.Timespan, workerCfg.Srcpath)
	}

	fmt.Println("")
	fmt.Println("Running workers :")
	fmt.Println("")
	for _, worker := range workerhub.workers {
		worker.Print()
	}

	workerhub.Start()
	console.Printf("\nHub started:\n %v\n", console.SDefault, console.BGBlack, console.FYellow,  &workerhub)

	for {
		time.Sleep(time.Minute)
	}
}
