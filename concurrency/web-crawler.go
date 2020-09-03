package concurrency

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type fetchedUrls struct {
	urls map[string]bool
	mut  sync.Mutex
}

func (f *fetchedUrls) Add(url string) bool {
	f.mut.Lock()
	defer f.mut.Unlock()
	url = strings.Trim(url, "/")
	if _, ok := f.urls[url]; ok {
		return false
	}
	f.urls[url] = false
	return true
}

func (f *fetchedUrls) Update(url string, fetched bool) bool {
	f.mut.Lock()
	defer f.mut.Unlock()
	url = strings.Trim(url, "/")
	if _, ok := f.urls[url]; !ok {
		return false
	}
	f.urls[url] = fetched
	return true
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl the given URL and gather all links
func Crawl(url string, depth int) []string {
	return crawl(url, depth, httpFetcher{})
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func crawl(url string, depth int, fetcher Fetcher) []string {
	quit := make(chan bool)
	fetchedUrls := fetchedUrls{urls: make(map[string]bool)}
	go crawlRecursive(url, depth, fetcher, &fetchedUrls, quit)

	// We will not quit until we have something
	// in the "quit" channel
	<-quit
	keysUrls := make([]string, 0, len(fetchedUrls.urls))
	for u, fetched := range fetchedUrls.urls {
		if fetched {
			keysUrls = append(keysUrls, u)
		}
	}

	return keysUrls
}

func crawlRecursive(url1 string, depth int, fetcher Fetcher, fetchedUrls *fetchedUrls, quit chan bool) {
	defer func() {
		fetchedUrls.Update(url1, true)
		quit <- true
	}()
	if depth <= 0 {
		return
	}
	if fetchedUrls.Add(url1) != true {
		return
	}

	_, urls, err := fetcher.Fetch(url1)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("found: %s %q\n", url1, body)

	quitChild := make(chan bool)
	for _, u := range urls {
		if !strings.Contains(u, url1) {
			continue
		}
		go crawlRecursive(u, depth-1, fetcher, fetchedUrls, quitChild)
	}
	// To exit goroutines. This channel will always be filled
	for _, u := range urls {
		if !strings.Contains(u, url1) {
			continue
		}
		<-quitChild
	}
}

// http fetcher that fetches an url
type httpFetcher struct {
	url    string
	result result
}

type result struct {
	body string
	urls []string
}

func (f httpFetcher) Fetch(url1 string) (string, []string, error) {
	resp, err := http.Get(url1)
	if err != nil {
		return "", nil, fmt.Errorf("not found: %s", url1)
	}

	if resp.StatusCode != http.StatusOK {
		return "", nil, fmt.Errorf("HTTP status not OK: %s", url1)
	}

	//get all links
	doc, err := html.Parse(resp.Body)
	var links []string
	if err != nil {
		fmt.Printf("parsing %s as HTML: %v\n", url1, err)
		links = nil
	} else {
		visitNode := func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key != "href" {
						continue
					}
					if strings.HasPrefix(a.Val, "/") {
						//relative url, prefix with the domain
						parsedUrl, err := url.Parse(url1)
						if err != nil {
							continue
						}
						a.Val = parsedUrl.Scheme + "://" + parsedUrl.Host + a.Val
					}

					links = append(links, a.Val)
				}
			}
		}
		forEachNode(doc, visitNode, nil)
	}

	return doc.Data, links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
