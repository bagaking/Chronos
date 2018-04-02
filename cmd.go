package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

type Worker struct {
	timespan    time.Duration
	triggertime time.Time
	name        string
	scriptPth   string
}

func (worker *Worker) tryTrigger(triggerTime time.Time) ([]byte, error) {
	if worker.triggertime.Unix() == 0 || worker.triggertime.Add(worker.timespan).Unix() <= triggerTime.Unix() {
		worker.triggertime = triggerTime
		srcData, err := ioutil.ReadFile(worker.scriptPth)
		srcStr := string(srcData[:])
		out, err := exec.Command(srcStr).Output() // read the new src
		return out, err
	}
	return nil, errors.New("u shell not")
}

type Hub struct {
	workers []Worker
}

func (hub *Hub) Start() {
	timeChan := time.NewTimer(time.Second * 2).C
	tickChan := time.NewTicker(time.Millisecond * 500).C

	go func() {
		fmt.Printf("cmd hub start %s \n", <-timeChan)
		go func() {
			for {
				ctime := <-tickChan
				for _, worker := range hub.workers {
					result, err := worker.tryTrigger(ctime)
					if err != nil {
						fmt.Printf("worker %s @ %s \nconsole :\n%s\n", worker.name, ctime.Unix(), result)
					}
				}
			}
		}()
	}()
}

func (hub *Hub) Insert(timespan string, srcPath string) {
	dtimespan, _ := time.ParseDuration(timespan)
	newworker := Worker{
		timespan:  dtimespan,
		scriptPth: srcPath,
	}
	fmt.Printf("create worker %s ", newworker)
	hub.workers = append(hub.workers, newworker)
}
