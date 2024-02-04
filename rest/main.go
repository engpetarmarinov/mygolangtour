package main

import (
	"github.com/engpetarmarinov/mygolangtour/rest/database"
	"github.com/engpetarmarinov/mygolangtour/rest/product"
	"log"
	"net/http"
)

const apiAddr = ":5000"
const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)

	log.Fatal(http.ListenAndServe(apiAddr, nil))
}
