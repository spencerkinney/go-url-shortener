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
	if !ctx.IsPost() {
		ctx.Error("Method not allowed", fasthttp.StatusMethodNotAllowed)
		return
	}

	var req models.ShortenRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		ctx.Error("Bad request", fasthttp.StatusBadRequest)
		return
	}

	var shortCode string
	if req.CustomUrl == "" {
		shortCode = utils.GenerateShortURL()
	} else {
		if _, ok := urlMap.Load(req.CustomUrl); ok {
			url := fmt.Sprintf("http://%s/%s", ctx.Host(), req.CustomUrl)
			ctx.Error("Shortened URL "+url+" already exists!", fasthttp.StatusConflict)
			return
		}

		if len(req.CustomUrl) > 24 {
			ctx.Error("Custom URL cannot exceed 24 characters.", fasthttp.StatusBadRequest)
			return
		}

		shortCode = req.CustomUrl
	}

	// Assuming urlMap is a thread-safe map
	urlMap.Store(shortCode, req.URL)

	// for debugging purposes, it will significally slow down RPS
	//fmt.Println("Short URL created:", shortCode)

	resp := models.ShortenResponse{
		ShortURL: fmt.Sprintf("http://%s/%s", ctx.Host(), shortCode),
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	if bodyBytes, err := json.Marshal(resp); err == nil {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(bodyBytes)
	} else {
		ctx.Error("Unknown error occurred", fasthttp.StatusInternalServerError)
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

	if _, ok := urlMap.Load(shortCode); ok {
		//ctx.Redirect(url.(string), http.StatusFound)
		ctx.Redirect("https://www.example.com", http.StatusFound)
		return
	}

	ctx.NotFound()
}
