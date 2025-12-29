package main

import (
	"os"
)

// GetAPIURL returns the API URL from environment or default
func GetAPIURL() string {
	apiURL := os.Getenv("LISTY_API_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080"
	}
	return apiURL
}
