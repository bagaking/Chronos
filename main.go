package main
import (
	"time"
	"github.com/bagaking/chronos/chronos"
)

func main() {
	chronos.Start();

	for {
		time.Sleep(time.Minute)
	}
}
