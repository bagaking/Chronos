package main

import (
	"fmt"
	"os/exec"
	"time"
)

// Worker : bash script executor
type Worker struct {
	timespan    time.Duration
	triggertime time.Time
	name        string
	scriptPth   string
}

func (worker *Worker) tryTrigger(triggerTime time.Time) (out []byte, ok bool, err error) {
	if ok = worker.triggertime.Unix() <= 0 || worker.triggertime.Add(worker.timespan).Unix() <= triggerTime.Unix(); ok {
		fmt.Printf("enter trigger %s", worker.name)
		worker.triggertime = triggerTime
		//exec.Command("/bin/chmod", "a+x", worker.scriptPth).Output()
		out, err = exec.Command("/bin/sh", worker.scriptPth).Output() // read the new src
	}
	// else {
	// 	fmt.Printf("%#v %#v", worker.triggertime.Unix(), worker.triggertime.Add(worker.timespan).Unix())
	// }
	return
}

// Hub : manager of worker
type Hub struct {
	workers []Worker
}

// Start : start the hub and all workers
func (hub *Hub) Start() {
	println("enter hub")
	<-time.NewTimer(time.Second * 1).C
	println("prepare hub")
	tickChan := time.NewTicker(time.Millisecond * 1000).C

	hubexe := func() {
		//println("enter gocorou")
		fmt.Printf(" ======== cmd hub start %#v ======== ", &hub)
		for {
			ctime := <-tickChan
			count := 0
			for i := range hub.workers {
				worker := &hub.workers[i]
				result, ok, err := worker.tryTrigger(ctime)
				if ok {
					count++
					fmt.Printf("== worker [%s] @ %s == >\n%s", worker.name, worker.triggertime.Format("2006-01-02 15:04:05"), result)
				} else if err != nil {
					fmt.Println(err)
				}
			}
			switch {
			case count > 0:
				fmt.Printf(" -%#v- \n", count)
			default:
				println(" --- ")
			}
		}
	}
	println("hub prepared")
	go hubexe()
}

// Insert : create a new worker and insert it to the hub
func (hub *Hub) Insert(name string, timespan string, srcPath string) {
	dtimespan, _ := time.ParseDuration(timespan)
	newworker := Worker{
		name:      name,
		timespan:  dtimespan,
		scriptPth: srcPath,
	}
	fmt.Printf("\ncreate worker [%#v] < %#v > : %#v  ", newworker.name, newworker.timespan, newworker.scriptPth)
	hub.workers = append(hub.workers, newworker)
}
