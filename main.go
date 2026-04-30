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

	http.HandleFunc("/gravatar/", gravatar.Handler)

	addr := ":8080"
	log.Printf("starting Gravatar proxy server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
