package concurrency

import (
	"context"
	"fmt"
	"github.com/wildalmighty/mygolangtour/utils"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func TestFanOutPrimeFinder(ctx context.Context) {
	defer utils.TimeTrack(time.Now(), "TestFanOutPrimeFinder")
	random := func() interface{} { return rand.Intn(50000000) }
	randIntStream := ToInt(ctx, RepeatFn(ctx, random))

	numFinders := runtime.NumCPU()
	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = naivePrimeFinder(ctx, randIntStream)
	}

	for prime := range Take(ctx, FanIn(ctx, finders...), 10) {
		fmt.Printf("TestFanOutPrimeFinder: found prime number %d\n", prime)
	}
}

func naivePrimeFinder(ctx context.Context, intStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for integer := range intStream {
			integer -= 1
			prime := true
			for divisor := integer - 1; divisor > 1; divisor-- {
				if integer%divisor == 0 {
					prime = false
					break
				}
			}

			if prime {
				select {
				case <-ctx.Done():
					return
				case primeStream <- integer:
				}
			}
		}
	}()

	return primeStream
}

func ToInt(ctx context.Context, stream <-chan interface{}) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)

		for {
			select {
			case <-ctx.Done():
				return
			case value := <-stream:
				intStream <- value.(int)
			}
		}
	}()

	return intStream
}

func FanIn(
	ctx context.Context,
	channels ...<-chan interface{},
) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case multiplexedStream <- i:
			}
		}
	}

	// Select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
