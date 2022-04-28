package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{},0,10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		//simulate dequeue
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.Signal()
	}

	for i:=0; i<10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			go removeFromQueue(1*time.Second)
			//wait for signal
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct {}{})
		c.L.Unlock()
	}
}
