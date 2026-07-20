package ratelimit_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis/v2"
	prahariMid "prahari/shared/middleware/ratelimit"
	prahariRedis "prahari/shared/redis"
)

func TestRateLimiterMiddleware(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()

	client, _ := prahariRedis.NewClient(prahariRedis.Config{Address: mr.Addr()})
	defer client.Close()

	// Capacity 1 (1 requests limit), fill rate 1 token/sec
	rl := prahariMid.NewRateLimiter(client, 1, 1.0)

	middleware := rl.Middleware(func(r *http.Request) string {
		return "user-key-abc"
	})

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Request 1 -> Expect 200 OK
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec1 := httptest.NewRecorder()
	handler.ServeHTTP(rec1, req1)
	if rec1.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rec1.Code)
	}

	// Request 2 -> Expect 429 Too Many Requests
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec2 := httptest.NewRecorder()
	handler.ServeHTTP(rec2, req2)
	if rec2.Code != http.StatusTooManyRequests {
		t.Errorf("expected 429 Too Many Requests, got %d", rec2.Code)
	}
}
