package api

import (
	"errors"
	"net/http"

	"prahari/services/auth-service/internal/domain"
)

// MapError translates domain errors into standardized HTTP JSON responses.
func MapError(w http.ResponseWriter, err error) {
	if errors.Is(err, domain.ErrInvalidCredentials) {
		WriteError(w, http.StatusUnauthorized, "UNAUTHENTICATED", err.Error(), nil)
		return
	}
	
	if errors.Is(err, domain.ErrUserAlreadyExists) {
		WriteError(w, http.StatusConflict, "USER_ALREADY_EXISTS", err.Error(), nil)
		return
	}
	
	if errors.Is(err, domain.ErrUserNotFound) {
		WriteError(w, http.StatusNotFound, "USER_NOT_FOUND", err.Error(), nil)
		return
	}
	
	if errors.Is(err, domain.ErrTokenExpired) {
		WriteError(w, http.StatusUnauthorized, "TOKEN_EXPIRED", err.Error(), nil)
		return
	}
	
	if errors.Is(err, domain.ErrInvalidToken) {
		WriteError(w, http.StatusUnauthorized, "INVALID_TOKEN", err.Error(), nil)
		return
	}
	
	if errors.Is(err, domain.ErrValidationError) {
		WriteError(w, http.StatusUnprocessableEntity, "VALIDATION_FAILED", err.Error(), nil)
		return
	}
	
	// Fallback to internal server error to prevent raw database details leakage
	WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "An unexpected server error occurred", nil)
}
