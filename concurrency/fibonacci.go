package concurrency

import (
	"fmt"
	"github.com/engpetarmarinov/mygolangtour/utils"
	"time"
)

func Fibonacci(n int, c chan<- uint64) {
	x, y := uint64(0), uint64(1)
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func FibonacciPrint(n int) {
	defer utils.TimeTrack(time.Now(), fmt.Sprintf("fibonacci of %v", n))
	c := make(chan uint64, n)
	go Fibonacci(n, c)

	fmt.Printf("Fibonacci up to %vth element:", n)
	for i := range c {
		fmt.Printf(" %v", i)
	}
	fmt.Println("")
}
