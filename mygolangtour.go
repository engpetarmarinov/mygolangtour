package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/wildalmighty/mygolangtour/concurrency"
	"github.com/wildalmighty/mygolangtour/morestrings"
	"log"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}

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
			links := concurrency.Crawl("http://slavi.bg", 10)
			fmt.Printf("Fetched links: %v, Links: %v\n", len(links), links)
			log.Println("WebCrawlerPrint DONE")
		},
		func() {
			concurrency.FibonacciPrint(100)
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
