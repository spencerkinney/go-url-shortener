package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"os"
)

import (
	handlers "go-url-shortener/handlers"
)

// Global map to store the URL and its corresponding short URL
// For now, this is just a simple Map and is not persisted anywhere.
// todo later: Persist the map in a proper database with a caching layer.
var urlMap sync.Map

func main() {
	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.HomeOrRedirectHandler)

    port := os.Getenv("PORT")
    if port == "" { port = "8080" }

    fmt.Printf("Server listening on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}