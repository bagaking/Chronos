package chronos

import (
	"fmt"
	"os/exec"
	"time"
	"runtime"
	"os"
	"github.com/bagaking/bagakit/sbuilder"
	"github.com/bagaking/bagakit/console"
)

var isWindows bool = runtime.GOOS == "windows"

func getShell() string {
	var shell string
	if isWindows {
		shell = os.Getenv("COMSPEC")
		if shell == "" {
			shell = "C:\\WINDOWS\\System32\\cmd.exe"
		}
	} else {
		shell = os.Getenv("SHELL")
		if shell == "" {
			shell = "sh"
		}
	}
	return shell
}

// Worker : bash script executor
type Worker struct {
	timespan    time.Duration
	triggertime time.Time
	name        string
	scriptPth   string
}

func (worker *Worker) Print() {
	sb := &sbuilder.B{}
	sb.Appends("=w= ", worker.name, "(", worker.scriptPth ,")").NewLine()
	sb.Appends("timespan", worker.timespan).NewLine()
	sb.Appends("triggertime", worker.triggertime).NewLine()
	console.PrintCF("%s\n", console.SNONE, console.BGCLEAR, console.FGreen, sb.String())
}

func (worker *Worker) tryTrigger(triggerTime time.Time) (out []byte, ok bool, err error) {
	if ok = worker.triggertime.Unix() <= 0 || worker.triggertime.Add(worker.timespan).Unix() <= triggerTime.Unix(); ok {
		fmt.Printf("enter trigger %s\n", worker.name)
		worker.triggertime = triggerTime
		//exec.Command("/bin/chmod", "a+x", worker.scriptPth).Output()
		executor := exec.Command(getShell(), worker.scriptPth)
		out, err = executor.Output() // read the new src
		fmt.Println("mission complete.")
	}
	// else {
	// 	fmt.Printf("%#v %#v", worker.triggertime.Unix(), worker.triggertime.Add(worker.timespan).Unix())
	// }
	return
}
