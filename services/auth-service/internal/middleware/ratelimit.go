package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"prahari/services/auth-service/internal/api"
)

type ipBucket struct {
	tokens     float64
	lastRefill time.Time
}

type IPRateLimiter struct {
	mu           sync.Mutex
	clients      map[string]*ipBucket
	rate         float64 // tokens per second
	burst        float64 // max bucket size
	clientMaxAge time.Duration
}

func NewIPRateLimiter(rate, burst float64) *IPRateLimiter {
	limiter := &IPRateLimiter{
		clients:      make(map[string]*ipBucket),
		rate:         rate,
		burst:        burst,
		clientMaxAge: 10 * time.Minute,
	}
	
	// Start clean-up goroutine for stale clients
	go limiter.startCleanupRoutine()
	return limiter
}

func (l *IPRateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}
		
		if !l.allow(ip) {
			api.WriteError(w, http.StatusTooManyRequests, "TOO_MANY_REQUESTS", "Rate limit exceeded. Please try again later.", nil)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func (l *IPRateLimiter) allow(ip string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	now := time.Now()
	bucket, exists := l.clients[ip]
	if !exists {
		l.clients[ip] = &ipBucket{
			tokens:     l.burst - 1, // Consume one token on first request
			lastRefill: now,
		}
		return true
	}
	
	// Refill tokens based on time passed
	elapsed := now.Sub(bucket.lastRefill).Seconds()
	bucket.lastRefill = now
	bucket.tokens += elapsed * l.rate
	if bucket.tokens > l.burst {
		bucket.tokens = l.burst
	}
	
	// Check if token can be consumed
	if bucket.tokens >= 1 {
		bucket.tokens -= 1
		return true
	}
	
	return false
}

func (l *IPRateLimiter) startCleanupRoutine() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		l.mu.Lock()
		now := time.Now()
		for ip, bucket := range l.clients {
			// If client has not requested in the max age window, clean up memory
			if now.Sub(bucket.lastRefill) > l.clientMaxAge {
				delete(l.clients, ip)
			}
		}
		l.mu.Unlock()
	}
}
