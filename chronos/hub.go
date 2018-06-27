package chronos

import (
	"fmt"
	"time"
)

// Hub : manager of worker
type Hub struct {
	Workers []Worker
}

// Start : start the hub and all workers
func (hub *Hub) Start() {
	println("enter hub")
	<-time.NewTimer(time.Second * 1).C
	println("prepare hub")
	tickChan := time.NewTicker(time.Millisecond * 1000).C

	hubexe := func() {
		//println("enter gocorou")
		fmt.Printf(" ======== cmd hub start %#v ======== \n", &hub)
		for {
			ctime := <-tickChan
			count := 0
			for i := range hub.Workers {
				worker := &hub.Workers[i]
				result, ok, err := worker.tryTrigger(ctime)
				if ok {
					count++
					fmt.Printf("== worker [%s] @ %s == >\n%s\n", worker.name, worker.triggertime.Format("2006-01-02 15:04:05"), result)
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
func (hub *Hub) Insert(name string, timespan string, srcPath string) *Worker {
	dtimespan, _ := time.ParseDuration(timespan)
	newworker := Worker{
		name:      name,
		timespan:  dtimespan,
		scriptPth: srcPath,
	}
	fmt.Printf("\ncreate worker [%#v] < %#v > : %#v  \n", newworker.name, newworker.timespan, newworker.scriptPth)
	hub.Workers = append(hub.Workers, newworker)
	return &newworker
}

