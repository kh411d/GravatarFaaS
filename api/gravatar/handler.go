package gravatar

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	apiKey := strings.TrimSpace(os.Getenv("GRAVATAR_API_KEY"))
	if apiKey == "" {
		http.Error(w, "GRAVATAR_API_KEY environment variable is required", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	profileIdentifier := strings.TrimSpace(os.Getenv("GRAVATAR_PROFILE_ID"))
	if profileIdentifier == "" {
		http.Error(w, "GRAVATAR_PROFILE_ID environment variable is required", http.StatusBadRequest)
		return
	}

	proxyGravatarProfile(w, r, apiKey, profileIdentifier)
}

func proxyGravatarProfile(w http.ResponseWriter, r *http.Request, apiKey, profileIdentifier string) {
	apiURL := fmt.Sprintf("https://api.gravatar.com/v3/profiles/%s", profileIdentifier)
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "failed to call Gravatar API", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)

	_, copyErr := io.Copy(w, resp.Body)
	if copyErr != nil {
		log.Printf("error copying response body: %v", copyErr)
	}
}
