package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs requests.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Received request: %s %s", r.Method, r.URL.Path)
        next(w, r)
    }
}

// ErrorHandlingMiddleware handles panics and errors.
func ErrorHandlingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        next(w, r)
    }
}

// AuthMiddleware checks if the user is authenticated.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Dummy authentication check
        if r.Header.Get("Authorization") != "Bearer valid_token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next(w, r)
    }
}

// RateLimitMiddleware limits the rate of requests.
func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
    var lastRequestTime time.Time
    const rateLimit = 1 * time.Second

    return func(w http.ResponseWriter, r *http.Request) {
        now := time.Now()
        if now.Sub(lastRequestTime) < rateLimit {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }
        lastRequestTime = now
        next(w, r)
    }
}
