package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/wildalmighty/mygolangtour/concurrency"
	"github.com/wildalmighty/mygolangtour/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	WebCrawlerPrint()

	concurrency.FibonacciPrint(100)
	concurrency.PlantABomb(3)
	concurrency.CountTo(1000)
	concurrency.CompareEquivalentBinaryTreesTest()
}

func WebCrawlerPrint() {
	crawlerResult := make(chan []string)
	go func() {
		crawlerResult <- concurrency.Crawl("http://slavi.bg", 10)
	}()
	defer func() {
		links := <-crawlerResult
		fmt.Printf("Fetched links: %v, Links: %v\n", len(links), links)
	}()
}
