package concurrency

import (
	"fmt"
	"testing"
)

func TestCrawl(t *testing.T) {
	fetchedUrls := crawl("https://golang.org/", 4, fetcher)
	expected := []string{
		"https://golang.org", "https://golang.org/pkg", "https://golang.org/pkg/fmt", "https://golang.org/pkg/os", "https://golang.org/cmd",
	}
	for _, fetchedUrl := range fetchedUrls {
		if !contains(expected, fetchedUrl) {
			t.Errorf("Fetched url %s is not expected", fetchedUrl)
		}
	}
	notExpected := []string{
		"https://golang.org/cmd/",
	}
	for _, fetchedUrl := range fetchedUrls {
		if contains(notExpected, fetchedUrl) {
			t.Errorf("Fetched url %s should not be found", fetchedUrl)
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*result

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &result{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &result{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &result{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &result{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
