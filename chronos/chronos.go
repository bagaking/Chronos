package chronos

import (
	"fmt"
	"github.com/bagaking/bagakit/console"
)

const ConfPATH = ".kh.chronos.json"

func Start() {
	conf := loadConfig(ConfPATH)
	fmt.Printf("\nConf loaded:\n %#v \n", conf)

	workerhub := &Hub{}
	fmt.Printf("\nHub loaded:\n %#v \n", workerhub)

	for _, workerCfg := range conf.Workers {
		workerhub.Insert(workerCfg.Workername, workerCfg.Timespan, workerCfg.Srcpath)
	}

	fmt.Println("")
	fmt.Println("Running workers :")
	fmt.Println("")
	for _, worker := range workerhub.Workers {
		worker.Print()
	}

	workerhub.Start()
	console.PrintCF("\nHub started:\n %v\n", console.SNONE, console.BGCLEAR, console.FYellow,  &workerhub)
}