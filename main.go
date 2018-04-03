package main

import (
	"fmt"
)

const (
	confPATH = ".kh.chronos.json"
)

func main() {

	conf := loadConfig(confPATH)
	fmt.Printf("Conf loaded %s\n", conf)
	fmt.Println("Hub started")
	workerhub := &Hub{}
	for _, workerCfg := range conf.Workers {
		workerhub.Insert(workerCfg.Workername, workerCfg.Timespan, workerCfg.Srcpath)
	}
	workerhub.Start()
}
