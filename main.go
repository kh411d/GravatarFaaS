package main

import (
	"gravatarfaas/api"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	apiKey := strings.TrimSpace(os.Getenv("GRAVATAR_API_KEY"))
	if apiKey == "" {
		log.Fatal("GRAVATAR_API_KEY environment variable is required")
	}

	http.HandleFunc("/profile/", api.Profile)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}
