package main

import (
	"github.com/wildalmighty/mygolangtour/rest/product"
	"log"
	"net/http"
)

const apiAddr = ":5000"
const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)

	log.Fatal(http.ListenAndServe(apiAddr, nil))
}
