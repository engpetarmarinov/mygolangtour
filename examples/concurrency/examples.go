package main

import (
	"context"
	"fmt"
	"github.com/engpetarmarinov/mygolangtour/concurrency"
	"github.com/engpetarmarinov/mygolangtour/utils"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg = sync.WaitGroup{}
	ctx := context.Background()

	listOfFunctions := []func(){
		func() {
			links := concurrency.Crawl("https://slavi.bg", 2)
			fmt.Printf("Fetched links: %v, Links: %v\n", len(links), links)
		},
		func() {
			concurrency.FibonacciPrint(100)
		},
		func() {
			//pipeline
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			random := func() interface{} { return rand.Int() }
			for num := range concurrency.Take(ctx, concurrency.RepeatFn(ctx, random), 10) {
				fmt.Printf("Pipeline: take rand %d\n", num)
			}
		},
		func() {
			//fan-out and fan-in
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			concurrency.TestFanOutPrimeFinder(ctx)
		},
		func() {
			defer utils.TimeTrack(time.Now(), "or-channel")
			sig := func(after time.Duration) <-chan struct{} {
				c := make(chan struct{})
				tick := time.Tick(after)
				go func() {
					defer close(c)
					<-tick
				}()

				return c
			}

			<-concurrency.Or(
				sig(2*time.Hour),
				sig(1*time.Second),
				sig(3*time.Second),
				sig(1*time.Minute),
				sig(3*time.Minute),
			)
		},
		func() {
			for result := range concurrency.CheckStatus(ctx.Done(), "https://slavi.bg", "http://ffoooo.baar") {
				if result.Error != nil {
					fmt.Printf("CheckStatus error %v\n", result.Error)
					continue
				}

				fmt.Printf("CheckStatus response %v\n", result.Response)
			}
		},
		func() {
			concurrency.PlantABomb(3)
		},
		func() {
			concurrency.CountTo(10000000)
		},
		func() {
			concurrency.CompareEquivalentBinaryTreesTest()
		},
		func() {
			t1, _ := concurrency.NewTriangle(3, 3)
			s1, _ := concurrency.NewSquare(5.0)

			fmt.Printf("Shapes: FanOut sum of the areas is %f\n", concurrency.CalculateArea(t1, s1))
			fmt.Printf("Shapes: WG sum of the areas is %f\n", concurrency.CalculateAreaWithWG(t1, s1))
			fmt.Printf("Shapes: JASE sum of the areas is %f\n", concurrency.CalculateAreaWithoutChan(t1, s1))
		},
	}

	for _, fun := range listOfFunctions {
		execInWaitGroup(fun, &wg)
	}

	wg.Wait()
}

func execInWaitGroup(funcToExec func(), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		funcToExec()
		wg.Done()
	}()
}
