package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "port to listen to")
	flag.Parse()

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /{username}", serveUser)

	log.Printf("Starting the server at %s\n", *addr)
	http.ListenAndServe(*addr, mux)
}
