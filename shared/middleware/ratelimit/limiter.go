package ratelimit

import (
	"encoding/json"
	"net/http"

	prahariRedis "prahari/shared/redis"
	prahariLimit "prahari/shared/redis/ratelimit"
)

// RateLimiter wraps Redis token-bucket rate limiters into HTTP filters.
type RateLimiter struct {
	limiter *prahariLimit.TokenBucketLimiter
}

// NewRateLimiter constructs a RateLimiter.
func NewRateLimiter(client *prahariRedis.Client, capacity int64, fillRate float64) *RateLimiter {
	return &RateLimiter{
		limiter: prahariLimit.NewTokenBucketLimiter(client, capacity, fillRate),
	}
}

// Middleware creates a handler checking request rate quotas.
func (rl *RateLimiter) Middleware(keyFunc func(r *http.Request) string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := keyFunc(r)
			if key == "" {
				next.ServeHTTP(w, r)
				return
			}

			allowed, err := rl.limiter.Allow(r.Context(), key)
			if err != nil || !allowed {
				w.Header().Set("Content-Type", "application/problem+json")
				w.WriteHeader(http.StatusTooManyRequests)

				problem := map[string]interface{}{
					"type":   "about:blank",
					"title":  "Too Many Requests",
					"status": http.StatusTooManyRequests,
					"detail": "Rate limit quota exceeded. Please try again later.",
				}

				_ = json.NewEncoder(w).Encode(problem)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
