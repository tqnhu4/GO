package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Define a struct for our URL request and response
type URLRequest struct {
	LongURL string `json:"long_url"`
}

type URLResponse struct {
	ShortURL string `json:"short_url"`
}

// In-memory storage for URL mappings
var urlStore = make(map[string]string)
var mu sync.Mutex // Mutex to protect urlStore for concurrent access

const shortCodeLength = 8 // Length of the generated short code

// generateShortCode creates a unique random short code
func generateShortCode() string {
	b := make([]byte, shortCodeLength)
	if _, err := rand.Read(b); err != nil {
		log.Printf("Error generating random bytes: %v", err)
		// Fallback to a less robust, but still unique (for this simple app) method
		return fmt.Sprintf("%x", b)[:shortCodeLength]
	}
	return base64.URLEncoding.EncodeToString(b)[:shortCodeLength]
}

// shortenURLHandler handles POST requests to /shorten
func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req URLRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.LongURL == "" {
		http.Error(w, "Long URL cannot be empty", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Check if the URL is already shortened to avoid duplicate short codes for the same long URL
	for short, long := range urlStore {
		if long == req.LongURL {
			// Found existing short URL for this long URL
			resp := URLResponse{ShortURL: fmt.Sprintf("http://localhost:8080/%s", short)}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
	}

	shortCode := generateShortCode()
	urlStore[shortCode] = req.LongURL

	resp := URLResponse{ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortCode)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// redirectHandler handles GET requests for short codes
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:] // Remove leading slash

	mu.Lock()
	longURL, ok := urlStore[shortCode]
	mu.Unlock()

	if !ok {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound) // 302 Found
}

func main() {
	http.HandleFunc("/shorten", shortenURLHandler)
	http.HandleFunc("/", redirectHandler) // Handles all other paths, assuming they are short codes

	port := ":8080"
	fmt.Printf("URL Shortener server listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}