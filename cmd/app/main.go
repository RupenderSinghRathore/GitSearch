package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "port to listen to")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /{username}", serveUser)

	log.Printf("Starting the server at %s\n", *addr)
	http.ListenAndServe(*addr, mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
func serveUser(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("username")
	url := fmt.Sprintf("https://api.github.com/users/%s", user)
	res, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	userInfo := make(map[string]any)
	json.Unmarshal(data, &userInfo)
	// fmt.Printf("map: %v\n", userInfo)
	fmt.Printf("bio: %s\n", userInfo["bio"])
	w.Write(data)
}
