package main

import (
	"fmt"
	"time"
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
	fmt.Printf("\n%c[1;40;32mHub started:\n %#v%c[0m\n", 0x1B, &workerhub, 0x1B)

	for {
		time.Sleep(time.Minute)
	}
}
