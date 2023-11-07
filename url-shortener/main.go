package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	handlers "go-url-shortener/handlers"
)

// Global map to store the URL and its corresponding short URL
// For now, this is just a simple Map and is not persisted anywhere.
// todo later: Persist the map in a proper database with a caching layer.
var urlMap sync.Map

func main() {
	http.HandleFunc("/shorten", CORS(handlers.ShortenHandler))
	http.HandleFunc("/", CORS(handlers.HomeOrRedirectHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}
