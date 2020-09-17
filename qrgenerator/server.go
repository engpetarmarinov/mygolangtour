package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ *template.Template

type Page struct {
	Title string
	Url   string
}

func init() {
	templateFilename := "./templates/index.gohtml"
	tpl, err := template.ParseFiles(templateFilename)
	if err != nil {
		log.Fatalln(err)
	}
	templ = tpl
}

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	err := templ.Execute(w, Page{
		Title: "QR Code Generator",
		Url:   req.FormValue("s"),
	})
	if err != nil {
		log.Fatal("QR handler", err)
	}
}
