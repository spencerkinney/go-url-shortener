package main

import (
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
		// fmt.Println("Received inbound request")
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
	// else {
	// 	fmt.Printf("Server listening on port %s\n", port)
	// }
}

// func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
// 	// fmt.Println("Setting CORS for context")
// 	var (
// 		allowOrigin      = "*"
// 		allowCredentials = "true"
// 		allowHeaders     = "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"
// 		allowedMethods   = "GET, POST"
// 	)
// 	return func(ctx *fasthttp.RequestCtx) {
// 		// fmt.Println("Setting response headers for context")
// 		ctx.Response.Header.Set("Access-Control-Allow-Origin", allowOrigin)
// 		ctx.Response.Header.Set("Access-Control-Allow-Credentials", allowCredentials)
// 		ctx.Response.Header.Set("Access-Control-Allow-Headers", allowHeaders)
// 		ctx.Response.Header.Set("Access-Control-Allow-Methods", allowedMethods)

// 		next(ctx)
// 	}
// }
