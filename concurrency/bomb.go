package concurrency

import (
	"fmt"
	"time"
)

func PlantABomb(secs int) {
	tick := time.Tick(time.Second)
	boom := time.After(time.Duration(secs) * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		}
	}
}
