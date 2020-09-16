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
	ValidateMyReader()

	go WebCrawlerPrint()

	concurrency.FibonacciPrint(10000)
	concurrency.PlantABomb(3)
	concurrency.CountTo(1000)
	concurrency.CompareEquivalentBinaryTreesTest()
}

func WebCrawlerPrint() {
	links := concurrency.Crawl("http://slavi.bg", 10)
	fmt.Printf("Fetched links: %v, Links: %v\n", len(links), links)
}
