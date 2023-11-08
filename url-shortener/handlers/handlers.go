package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	models "go-url-shortener/models"
	utils "go-url-shortener/utils"

	"github.com/valyala/fasthttp"
)

var urlMap sync.Map

// Handler to create a short URL
func ShortenHandler(ctx *fasthttp.RequestCtx) {
	// fmt.Println("Received shorten handler request")
	// fmt.Println(string(ctx.Response.Header.Peek("Access-Control-Allow-Origin")))
	if string(ctx.Method()) != http.MethodPost {
		ctx.Error("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.ShortenRequest

	if _, err := utils.ReadShortenRequestBody(ctx.PostBody(), req); err != nil {
		ctx.Error("Bad request", http.StatusBadRequest)
	}

	var shortCode string
	if req.CustomUrl == "" {
		shortCode = utils.GenerateShortURL()
	} else {
		if _, ok := urlMap.Load(req.CustomUrl); ok {
			url := fmt.Sprintf("http://%s/%s", string(ctx.Host()), req.CustomUrl)
			ctx.Error("Shortened URL "+url+" already exists!", http.StatusConflict)
			return
		}

		if len(req.CustomUrl) > 24 {
			ctx.Error("Custom URL cannot exceed 24 characters.", http.StatusBadRequest)
			return
		}

		shortCode = req.CustomUrl
	}

	urlMap.Store(shortCode, req.URL)

	resp := models.ShortenResponse{
		ShortURL: fmt.Sprintf("http://%s/%s", string(ctx.Host()), shortCode),
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(resp) // Assume error handling here as per previous discussion
	if bodyBytes, err := json.Marshal(resp); err == nil {
		ctx.SetStatusCode(http.StatusOK)
		ctx.SetBody(bodyBytes)

	} else {
		ctx.Error("Unknown error occurred", http.StatusInternalServerError)
	}
}

// Handler to redirect to the original URL or render homepage
func HomeOrRedirectHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Path()) == "/" {
		utils.ServeHTMLHomepage(ctx) // Serve the homepage if the path is just "/"
		return
	}

	// If it's not the homepage, then it's a short URL redirect request
	shortCode := string(ctx.Path())[1:]

	if url, ok := urlMap.Load(shortCode); ok {
		ctx.Redirect(url.(string), http.StatusFound)
		return
	}

	ctx.NotFound()
}
