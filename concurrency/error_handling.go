package concurrency

import "net/http"

type Result struct {
	Error    error
	Response *http.Response
}

func CheckStatus(done <-chan struct{}, urls ...string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)

		for _, url := range urls {
			resp, err := http.Get(url)
			result := Result{
				Error:    err,
				Response: resp,
			}
			select {
			case <-done:
				return
			case results <- result:
			}
		}
	}()

	return results
}
