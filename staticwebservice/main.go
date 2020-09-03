package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const domain = "localhost"
const port = "3000"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./staticwebservice/root/")))

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

	fmt.Println("Starting Web Service on: ", "http://"+domain+":"+port)
	err := http.ListenAndServe(domain+":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
