package models

// Struct to read the request payload
type ShortenRequest struct {
	URL string `json:"url"`
	CustomUrl string `json:"custom_url"`
}

// Struct for the response payload
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}