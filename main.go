package main

import (
	"gravatarfaas/api/gravatar"
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

	http.HandleFunc("/profile/", gravatar.Handler)
	addr := ":3000"
	http.ListenAndServe(addr, nil)
}
