package utils

import (
	"log"
	"math/rand"
	"net/http"
	"os"
)

// Utility function to generate a short URL.
func GenerateShortURL() string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Handler function just for a temporary home page so I can test deploy it to Heroku
func ServeHTMLHomepage(w http.ResponseWriter, r *http.Request) {
	// Attempt to read the content of index.html
	htmlContent, err := os.ReadFile("utils/index.html") // Ensure this is the correct relative path to the file
	if err != nil {
		// Log the error with more details and send an HTTP 500 Internal Server Error
		log.Printf("Error reading index.html file: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to inform the client that HTML is being sent
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write the HTML content to the response writer
	_, err = w.Write(htmlContent)
	if err != nil {
		// If an error occurs while writing the response, log it and send an HTTP 500 error
		log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}