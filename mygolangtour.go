package main

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/wildalmighty/mygolangtour/concurrency"
	"github.com/wildalmighty/mygolangtour/morestrings"
	"log"
	"sync"
	"time"
)

func main() {
	var wg = sync.WaitGroup{}
	ctx := context.Background()

	listOfFuncs := []func(){
		func() {
			fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
			fmt.Println(cmp.Diff("Hello World", "Hello Go"))
		},
		func() {
			ValidateMyReader()
		},
		func() {
			log.Println("WebCrawlerPrint STARTED")
			links := concurrency.Crawl("https://slavi.bg", 10)
			fmt.Printf("Fetched links: %v, Links: %v\n", len(links), links)
			log.Println("WebCrawlerPrint DONE")
		},
		func() {
			concurrency.FibonacciPrint(100)
		},
		func() {
			sig := func(after time.Duration) <-chan struct{} {
				c := make(chan struct{})
				tick := time.Tick(after)
				go func() {
					defer close(c)
					select {
					case <-tick:
					}
				}()

				return c
			}
			start := time.Now()
			<-concurrency.Or(
				sig(2*time.Hour),
				sig(1*time.Second),
				sig(3*time.Second),
			)
			fmt.Printf("or-channel done after %v\n", time.Since(start))
		},
		func() {
			for result := range concurrency.CheckStatus(ctx.Done(), "https://slavi.bg", "http://ffoooo.baar") {
				if result.Error != nil {
					fmt.Printf("CheckStatus error %v", result.Error)
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
	}

	for _, fun := range listOfFuncs {
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
