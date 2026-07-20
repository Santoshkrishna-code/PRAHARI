package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	prahariMid "prahari/shared/middleware"
	"prahari/shared/middleware/auth"
	prahariJWT "prahari/shared/security/jwt"
)

func TestAuthorizeRole_NoClaims(t *testing.T) {
	handler := auth.AuthorizeRole("Admin")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Errorf("expected 403 Forbidden on missing claims, got %d", rec.Code)
	}
}

func TestAuthorizeRole_RoleMismatchAndMatch(t *testing.T) {
	handler := auth.AuthorizeRole("Admin")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	claims := &prahariJWT.Claims{Role: "Worker"}

	// Test 1: Mismatched role -> Expect 403 Forbidden
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx1 := prahariMid.WithClaims(req1.Context(), claims)
	rec1 := httptest.NewRecorder()
	handler.ServeHTTP(rec1, req1.WithContext(ctx1))

	if rec1.Code != http.StatusForbidden {
		t.Errorf("expected 403 Forbidden on role mismatch, got %d", rec1.Code)
	}

	// Test 2: Matching role -> Expect 200 OK
	claims.Role = "Admin"
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx2 := prahariMid.WithClaims(req2.Context(), claims)
	rec2 := httptest.NewRecorder()
	handler.ServeHTTP(rec2, req2.WithContext(ctx2))

	if rec2.Code != http.StatusOK {
		t.Errorf("expected 200 OK on matching role, got %d", rec2.Code)
	}
}
