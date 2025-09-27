package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "port to listen to")
	flag.Parse()

	log.Printf("Starting the server at %s\n", *addr)
	http.ListenAndServe(*addr, routes())
}
