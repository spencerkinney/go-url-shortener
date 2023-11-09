package main

import (
	"log"
	"os"

	handlers "go-url-shortener/handlers"

	"github.com/valyala/fasthttp"
)

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		// Set CORS headers
		setCORSHeaders(ctx)

		// Handle preflight requests
		if string(ctx.Method()) == "OPTIONS" {
			// Preflight request only needs headers
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}

		path := string(ctx.Path())
		switch path {
		case "/shorten":
			handlers.ShortenHandler(ctx)
		default:
			handlers.HomeOrRedirectHandler(ctx)
		}

	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("listening on port", port)
	if err := fasthttp.ListenAndServe(":"+port, requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}

func setCORSHeaders(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Authorization, Accept, Content-Type")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
}