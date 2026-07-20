package jwt

import (
	"errors"
	"net/http"
	"strings"
)

var (
	ErrEmptyAuthHeader   = errors.New("authorization header is empty")
	ErrInvalidAuthFormat = errors.New("authorization header format must be Bearer token")
	ErrCookieNotFound    = errors.New("target cookie not found in request")
)

// ExtractTokenFromHeader parses the Bearer token from the Authorization header.
func ExtractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", ErrInvalidAuthFormat
	}

	return parts[1], nil
}

// ExtractTokenFromCookie retrieves the token from the specified cookie name.
func ExtractTokenFromCookie(r *http.Request, cookieName string) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return "", ErrCookieNotFound
		}
		return "", err
	}

	return cookie.Value, nil
}
