package main

import (
	"fmt"
	"os/exec"
	"time"
)

type Worker struct {
	timespan    time.Duration
	triggertime time.Time
	name        string
	scriptPth   string
}

func (worker *Worker) tryTrigger(triggerTime time.Time) (out []byte, ok bool, err error) {
	if ok = worker.triggertime.Unix() <= 0 || worker.triggertime.Add(worker.timespan).Unix() <= triggerTime.Unix(); ok {
		worker.triggertime = triggerTime
		//exec.Command("/bin/chmod", "a+x", worker.scriptPth).Output()
		out, err = exec.Command("/bin/sh", worker.scriptPth).Output() // read the new src
	}
	//fmt.Printf("\n== %#v %#v %#v %#v\n", worker.triggertime.Unix(), worker.triggertime.Add(worker.timespan).Unix(), triggerTime.Unix(), worker.triggertime.Add(worker.timespan).Unix() <= triggerTime.Unix())
	return
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
				for i, _ := range hub.workers {
					worker := &hub.workers[i]
					result, ok, err := worker.tryTrigger(ctime)
					if ok {
						fmt.Printf("== worker [%s] @ %s == >\n%s", worker.name, worker.triggertime.Format("2006-01-02 15:04:05"), result)
					} else if err != nil {
						fmt.Println(err)
					}
				}
			}
		}()
	}()
}

func (hub *Hub) Insert(name string, timespan string, srcPath string) {
	dtimespan, _ := time.ParseDuration(timespan)
	newworker := Worker{
		name:      name,
		timespan:  dtimespan,
		scriptPth: srcPath,
	}
	fmt.Printf("\ncreate worker ==> %s \n", newworker)
	hub.workers = append(hub.workers, newworker)
}
