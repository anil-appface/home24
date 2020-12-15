package main

import (
	"net/http"
	"os"
	"text/template"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	//load the assets
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	handler := newIndexHandler()
	mux.Handle("/", handler)

	//start and serve the server
	http.ListenAndServe(":"+port, mux)
}
