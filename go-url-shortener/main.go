package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

import (
	handlers "go-url-shortener/handlers"
)

// Global map to store the URL and its corresponding short URL
var urlMap sync.Map

func main() {
	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.RedirectHandler)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}