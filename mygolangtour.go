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

	execInWaitGroup(func() {
		fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
		fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	}, &wg)

	execInWaitGroup(func() {
		ValidateMyReader()
	}, &wg)

	execInWaitGroup(func() {
		log.Println("WebCrawlerPrint STARTED")
		links := concurrency.Crawl("http://slavi.bg", 10)
		fmt.Printf("Fetched links: %v, Links: %v\n", len(links), links)
		log.Println("WebCrawlerPrint DONE")
	}, &wg)

	execInWaitGroup(func() {
		concurrency.FibonacciPrint(100)
	}, &wg)

	execInWaitGroup(func() {
		concurrency.PlantABomb(3)
	}, &wg)

	execInWaitGroup(func() {
		concurrency.CountTo(100000000)
	}, &wg)

	execInWaitGroup(func() {
		concurrency.CompareEquivalentBinaryTreesTest()
	}, &wg)

	wg.Wait()
}

func execInWaitGroup(funcToExec func(), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		funcToExec()
		wg.Done()
	}()
}
