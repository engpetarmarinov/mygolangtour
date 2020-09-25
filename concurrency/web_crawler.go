package concurrency

import (
	"fmt"
	"github.com/wildalmighty/mygolangtour/utils"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"sync"
	"time"
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
	defer utils.TimeTrack(time.Now(), "crawler")

	runtime.GOMAXPROCS(runtime.NumCPU())

	fetchedUrls := fetchedUrls{urls: make(map[string]bool)}
	var waitGrp = sync.WaitGroup{}
	waitGrp.Add(1)
	go crawlRecursive(url, depth, fetcher, &fetchedUrls, &waitGrp)

	waitGrp.Wait()

	keysUrls := make([]string, 0, len(fetchedUrls.urls))
	for u, fetched := range fetchedUrls.urls {
		if fetched {
			keysUrls = append(keysUrls, u)
		}
	}

	return keysUrls
}

func crawlRecursive(urlToFetch string, depth int, fetcher Fetcher, fetchedUrls *fetchedUrls, waitGrp *sync.WaitGroup) {
	defer func() {
		fetchedUrls.Update(urlToFetch, true)
		waitGrp.Done()
	}()
	if depth <= 0 {
		return
	}
	if fetchedUrls.Add(urlToFetch) != true {
		return
	}

	_, urls, err := fetcher.Fetch(urlToFetch)
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGrpChildren = sync.WaitGroup{}

	for _, u := range urls {
		if !strings.Contains(u, urlToFetch) {
			continue
		}
		waitGrpChildren.Add(1)
		go crawlRecursive(u, depth-1, fetcher, fetchedUrls, &waitGrpChildren)
	}
	waitGrpChildren.Wait()
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
	client := http.Client{Timeout: time.Second * 3}
	resp, err := client.Get(url1)
	if err != nil {
		return "", nil, fmt.Errorf("error fetching: %s, error: %s", url1, err)
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
