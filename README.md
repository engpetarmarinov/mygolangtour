# mygolangtour
![Build Status](https://travis-ci.org/wildalmighty/mygolangtour.svg?branch=master)

Just a playground to keep track of my GoLang learning curve
* Exercises taken from https://tour.golang.org/
###Requirements
* Go 1.13

###CLI Cheatsheet
####Building and running
* Get it
  ```
  go get github.com/wildalmighty/mygolangtour
  ```
* Run it with -race flog to generate warnings for potential race conditions
  ```
  go run -race .
  ```
* Install it
  ```
  go install github.com/wildalmighty/mygolangtour
  ```
* Consume it and use it
  ```
  import "github.com/wildalmighty/mygolangtour/concurrency"
  ...
  concurrency.Crawl("http://slavi.bg", 10)
  ```
####Testing
* Test all
    ```
    go test ./...
    ```
* Test Cover
    ```
    go test -covermode count -coverprofile ./testdata/cover_concurrency.out ./concurrency
    go tool cover -html=./testdata/cover_concurrency.out
    ```
* Test Trace profile
    ```
    go test -trace ./testdata/trace_concurrency.out ./concurrency
    go tool trace ./testdata/trace_concurrency.out
    ```
* Test Benchmark

    Add a benchmark in a test
    ```
    func BenchmarkCrawl(b *testing.B) {
        for i:=0; i<b.N;i++ {
            crawl("https://golang.org/", 4, fetcher)
        }
    }
    ```
    and then run
    ```
    go test -benchtime 2s -bench . ./concurrency
    ```
####Others
* List all dependencies
    ```
    go list -f {{.Deps}} ./concurrency
    ```
* Lists the current module and all its dependencies
    ```
    go list -m all
    ```
* List all versions of the module
    ```
    go list -m -versions github.com/wildalmighty/mygolangtour
    ```
  and if you want to install an older version:
    ```
    go get github.com/wildalmighty/mygolangtour@v0.0.1
    ```
* Get doc for the API of a package
    ```
    go doc github.com/wildalmighty/mygolangtour/concurrency
    ```
    or use godoc to spin up a web server with the documentation:
    ```
    godoc -http :8000
    ```
    and open http://localhost:8000/pkg/github.com/wildalmighty/mygolangtour/concurrency/
