package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"go-url-shortener/models"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/valyala/fasthttp"
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
func ServeHTMLHomepage(ctx *fasthttp.RequestCtx) {
	// Attempt to read the content of index.html
	htmlContent, err := os.ReadFile("utils/index.html") // Ensure this is the correct relative path to the file
	if err != nil {
		// Log the error with more details and send an HTTP 500 Internal Server Error
		log.Printf("Error reading index.html file: %v", err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to inform the client that HTML is being sent
	ctx.Response.Header.Set("Content-Type", "text/html; charset=utf-8")

	// Write the HTML content to the response writer
	_, err = ctx.Write(htmlContent)
	if err != nil {
		// If an error occurs while writing the response, log it and send an HTTP 500 error
		log.Printf("Error writing response: %v", err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)
	}
}

// read the []byte request body using a buffer
func ReadShortenRequestBody(postBody []byte, reqObj models.ShortenRequest) (models.ShortenRequest, error) {
	buf := &bytes.Buffer{}
	buf.Write(postBody)
	_, buffWriteErr := buf.Write(postBody)
	buffReadErr := binary.Read(buf, binary.BigEndian, &reqObj)
	if buffWriteErr != nil || buffReadErr != nil {
		return models.ShortenRequest{}, errors.New("Error in reading request body")
	}
	return reqObj, nil
}
