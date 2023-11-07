package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "go-url-shortener/handlers"
)

func main() {
	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.HomeOrRedirectHandler)

    port := os.Getenv("PORT")
    if port == "" { port = "8080" }

    fmt.Printf("Server listening on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}