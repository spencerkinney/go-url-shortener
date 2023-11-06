package utils

import (
	"math/rand"
	"time"
)

// Since we're using random numbers, we need to seed the random number generator.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Utility function to generate a short URL.
func GenerateShortURL() string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}