package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	aiPlatformURL := os.Getenv("AI_PLATFORM_URL")
	if aiPlatformURL == "" {
		aiPlatformURL = "http://localhost:8000"
	}

	targetURL, err := url.Parse(aiPlatformURL)
	if err != nil {
		log.Fatalf("Invalid AI Platform target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":      "UP",
			"service":     "gateway-service",
			"timestamp":   time.Now().Format(time.RFC3339),
			"ai_platform": aiPlatformURL,
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[GATEWAY PROXY] %s %s -> %s", r.Method, r.URL.Path, aiPlatformURL)
		r.Host = targetURL.Host
		proxy.ServeHTTP(w, r)
	})

	log.Printf("PRAHARI Go API Gateway Listening on :%s (Proxying to %s)", port, aiPlatformURL)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("Gateway server error: %v", err)
	}
}
