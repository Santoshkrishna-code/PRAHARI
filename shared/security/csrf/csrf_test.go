package csrf_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"prahari/shared/security/csrf"
)

func TestGenerateToken(t *testing.T) {
	tok, err := csrf.GenerateToken()
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}
	if len(tok) != 64 { // 32 bytes hex encoded = 64 characters
		t.Errorf("expected 64 characters, got %d", len(tok))
	}
}

func TestCSRFMiddleware_GetGeneratesCookie(t *testing.T) {
	handler := csrf.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	resp := rec.Result()
	cookies := resp.Cookies()

	found := false
	for _, c := range cookies {
		if c.Name == csrf.CookieName {
			found = true
			if len(c.Value) != 64 {
				t.Errorf("expected CSRF cookie to hold token, got value length %d", len(c.Value))
			}
		}
	}
	if !found {
		t.Error("expected GET request to seed missing CSRF token cookie")
	}
}

func TestCSRFMiddleware_PostValidation(t *testing.T) {
	handler := csrf.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Case 1: Post request with missing token -> expected 403 Forbidden
	req1 := httptest.NewRequest(http.MethodPost, "/", nil)
	rec1 := httptest.NewRecorder()
	handler.ServeHTTP(rec1, req1)
	if rec1.Code != http.StatusForbidden {
		t.Errorf("expected 403 Forbidden on missing token, got %d", rec1.Code)
	}

	// Case 2: Post request with mismatched token -> expected 403 Forbidden
	req2 := httptest.NewRequest(http.MethodPost, "/", nil)
	req2.AddCookie(&http.Cookie{Name: csrf.CookieName, Value: "token-a"})
	req2.Header.Set(csrf.HeaderName, "token-b")
	rec2 := httptest.NewRecorder()
	handler.ServeHTTP(rec2, req2)
	if rec2.Code != http.StatusForbidden {
		t.Errorf("expected 403 Forbidden on token mismatch, got %d", rec2.Code)
	}

	// Case 3: Post request with matching token -> expected 200 OK
	req3 := httptest.NewRequest(http.MethodPost, "/", nil)
	req3.AddCookie(&http.Cookie{Name: csrf.CookieName, Value: "token-matching"})
	req3.Header.Set(csrf.HeaderName, "token-matching")
	rec3 := httptest.NewRecorder()
	handler.ServeHTTP(rec3, req3)
	if rec3.Code != http.StatusOK {
		t.Errorf("expected 200 OK on matching tokens, got %d", rec3.Code)
	}
}
