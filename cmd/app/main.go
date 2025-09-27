package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting the server at %s\n", port)
	http.ListenAndServe("0.0.0.0:"+port, routes())
}
