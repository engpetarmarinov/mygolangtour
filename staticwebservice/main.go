package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const domain = "localhost"
const port = "3000"

func main() {
	http.Handle("/files", http.FileServer(http.Dir("./staticwebservice/root/")))

	http.HandleFunc("/json", func(writer http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]

		jsonObject := map[string]string{}

		var name string
		if len(names) == 1 {
			name = names[0]
			jsonObject = map[string]string{"name": name}
		}
		enc := json.NewEncoder(writer)
		err := enc.Encode(jsonObject)
		if err != nil {
			log.Fatal(err)
		}
	})
	http.HandleFunc("/info", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("Info page"))
	})

	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("User page"))
	})

	//gorilla mux
	r := mux.NewRouter()

	productsSubrouter := r.PathPrefix("/products").Methods("GET").Subrouter()

	productsSubrouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("V1 products page"))
	}).
		MatcherFunc(func(request *http.Request, match *mux.RouteMatch) bool {
			apiVersion := request.Header.Get("api-version")
			return apiVersion == "1.0" || apiVersion == "2.0"
		}).
		Schemes("http")

	productsSubrouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("V3 products page"))
	}).
		Headers("api-version", "3.0").
		Schemes("http")

	http.Handle("/", r)

	fmt.Println("Starting Web Service on: ", "http://"+domain+":"+port)
	err := http.ListenAndServe(domain+":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
