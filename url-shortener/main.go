package main

import (
	"fmt"
	"log"
	"os"

	handlers "go-url-shortener/handlers"

	"github.com/valyala/fasthttp"
)

// func main() {
// 	http.HandleFunc("/shorten", CORS(handlers.ShortenHandler))
// 	http.HandleFunc("/", CORS(handlers.HomeOrRedirectHandler))

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}

// 	fmt.Printf("Server listening on port %s\n", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Println("Received inbound request")
		// ctx.Request.Header.Set("Access-Control-Allow-Origin", "*")
		// ctx.Request.Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// ctx.Request.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// route.POST("/shorten", handlers.ShortenHandler)
		// route.GET("/", CORS(handlers.HomeOrRedirectHandler))
		// route.GET("/test", CORS(handlers.TestHandler))
		var (
			allowOrigin      = "*"
			allowCredentials = "true"
			allowHeaders     = "Authorization accept Content-Type"
			allowedMethods   = "GET, POST"
		)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", allowOrigin)
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", allowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", allowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", allowedMethods)

		switch string(ctx.Path()) {
		case "/shorten":
			handlers.ShortenHandler(ctx)
		case "/":
			handlers.HomeOrRedirectHandler(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
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
	// else {
	// 	fmt.Printf("Server listening on port %s\n", port)
	// }
}

func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	var (
		allowOrigin      = "*"
		allowCredentials = "true"
		allowHeaders     = "Authorization accept Content-Type"
		allowedMethods   = "GET, POST"
	)
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", allowOrigin)
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", allowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", allowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", allowedMethods)

		next(ctx)
	}
}
