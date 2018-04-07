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
	fmt.Printf("\nConf loaded:\n %#v\n", conf)

	workerhub := &Hub{}
	fmt.Printf("\nHub loaded:\n %#v\n", workerhub)

	for _, workerCfg := range conf.Workers {
		workerhub.Insert(workerCfg.Workername, workerCfg.Timespan, workerCfg.Srcpath)
	}
	fmt.Printf("\nWorkers loaded:\n %#v\n", workerhub.workers)

	workerhub.Start()
	fmt.Printf("\nHub started:\n %#v\n", &workerhub)

	for {
		time.Sleep(time.Minute)
	}
}
