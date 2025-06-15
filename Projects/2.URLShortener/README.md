

# To create a URL shortener in Go, we'll need a few components:

1.  **A way to generate short codes:** These codes will be unique identifiers for our long URLs.
2.  **A storage mechanism:** To store the mapping between short codes and long URLs. For simplicity, we'll use an in-memory map, but in a real-world scenario, you'd use a database (like Redis, PostgreSQL, or MongoDB).
3.  **A web server:** To handle requests for shortening URLs and redirecting short URLs to their original destinations.

Here's a basic URL shortener application in Go:

---

## 🏗️ Go URL Shortener Application

This application provides a simple URL shortening service with two main endpoints:

* **`/shorten` (POST):** Accepts a long URL and returns a short URL.
* **`/{shortCode}` (GET):** Redirects to the original long URL associated with the given short code.

### 📁 Project Structure

```
url-shortener/
├── main.go
└── go.mod
```

### 📋 Prerequisites

* Go installed (version 1.16 or higher recommended)

### 🚀 Step-by-Step Implementation

#### 1. Initialize Go Module

First, create a new directory for your project and initialize a Go module:

```bash
mkdir url-shortener
cd url-shortener
go mod init url-shortener
```

#### 2. Create `main.go`

Now, create a `main.go` file inside the `url-shortener` directory and paste the following code:

```go
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

```

### 📝 Code Explanation

* **`URLRequest` and `URLResponse` structs:** Define the JSON structure for incoming requests and outgoing responses.
* **`urlStore`:** A global `map[string]string` that acts as our in-memory database, storing `shortCode -> longURL` mappings.
* **`mu sync.Mutex`:** A mutex is used to protect `urlStore` from concurrent access issues. Since `urlStore` is a shared resource, multiple goroutines accessing it simultaneously could lead to race conditions. The mutex ensures that only one goroutine can modify `urlStore` at a time.
* **`generateShortCode()`:**
    * Uses `crypto/rand` for cryptographically secure random bytes. This is important for generating unpredictable short codes.
    * `base64.URLEncoding` is used to encode these bytes into a web-safe string, ensuring the short codes don't contain characters that need special URL encoding.
    * It generates an 8-character short code.
* **`shortenURLHandler(w http.ResponseWriter, r *http.Request)`:**
    * Handles `POST` requests to the `/shorten` endpoint.
    * Decodes the JSON request body to get the `LongURL`.
    * Generates a `shortCode`.
    * **Crucially**, it checks if the `longURL` already exists in the `urlStore`. If it does, it returns the existing short URL to prevent creating multiple short codes for the same long URL.
    * Stores the `shortCode -> LongURL` mapping in `urlStore`.
    * Constructs and returns the full short URL in a JSON response.
* **`redirectHandler(w http.ResponseWriter, r *http.Request)`:**
    * Handles `GET` requests to any path that is not `/shorten`. It assumes these paths are `shortCode`s.
    * Extracts the `shortCode` from the URL path.
    * Looks up the corresponding `longURL` in `urlStore`.
    * If found, it performs an `http.Redirect` with `http.StatusFound` (302).
* **`main()`:**
    * Registers our handlers for the `/shorten` and `/` paths.
    * Starts the HTTP server on port `8080`.

### ▶️ How to Run

1.  **Navigate to the project directory:**
    ```bash
    cd url-shortener
    ```
2.  **Run the application:**
    ```bash
    go run main.go
    ```
    You should see the output: `URL Shortener server listening on :8080`

### 🧪 How to Test

You can use `curl` or any API client (like Postman, Insomnia) to test the application.

#### 1. Shorten a URL

Open a new terminal window and run:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"long_url": "https://www.example.com/this-is-a-very-long-url-that-we-want-to-shorten"}' http://localhost:8080/shorten
```

**Expected Output:**

```json
{"short_url":"http://localhost:8080/ABCDabcd"}
```
*(The actual short code will be different each time unless the long URL was already shortened.)*

#### 2. Redirect to the Original URL

Take the `short_url` from the previous step (e.g., `http://localhost:8080/ABCDabcd`) and paste it into your web browser, or use `curl` with the `-L` flag (to follow redirects):

```bash
curl -L http://localhost:8080/ABCDabcd
```

This `curl` command should redirect you to `https://www.example.com/this-is-a-very-long-url-that-we-want-to-shorten`. You can verify this by checking the browser's address bar or the `curl` output (it will show the content of the redirected page).

---

### ⚠️ Limitations & Future Improvements

This is a basic in-memory shortener. For a production-ready application, you would need to consider:

* **Persistent Storage:** Replace the `urlStore` map with a database (e.g., PostgreSQL, MongoDB, Redis) to store mappings permanently. This prevents data loss when the server restarts.
* **Error Handling:** More robust error handling and user-friendly error messages.
* **Concurrency:** While a mutex is used, a database would naturally handle more complex concurrency challenges.
* **Short Code Collision:** While `crypto/rand` makes collisions highly unlikely for a small number of URLs, a larger-scale system would need a strategy for handling potential collisions (e.g., retrying with a new code).
* **Custom Short Codes:** Allow users to specify their desired short code.
* **Analytics:** Track clicks on short URLs.
* **API Key/Authentication:** Protect the `/shorten` endpoint.
* **Custom Domain:** Use a custom domain for short URLs (e.g., `myshort.link/ABCDabcd`).

This application provides a solid foundation for understanding the core concepts of a URL shortener in Go!
