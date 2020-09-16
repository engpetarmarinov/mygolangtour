package concurrency

import (
	"fmt"
)

func Fibonacci(n int, c chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func FibonacciPrint(n int) {
	c := make(chan int, n)
	go Fibonacci(cap(c), c)
	fmt.Printf("Fibonacci up to %vth element:", n)
	for i := range c {
		fmt.Printf(" %v", i)
	}
	fmt.Println("")
}
