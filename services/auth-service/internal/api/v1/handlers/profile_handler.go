package handlers

import (
	"net/http"
	"strings"

	"prahari/services/auth-service/internal/api"
	"prahari/services/auth-service/internal/domain"
	"prahari/services/auth-service/internal/dto"
)

type ProfileHandler struct {
	useCase domain.AuthUseCase
}

func NewProfileHandler(useCase domain.AuthUseCase) *ProfileHandler {
	return &ProfileHandler{useCase: useCase}
}

func (h *ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		api.WriteError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Missing authorization header", nil)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		api.WriteError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Authorization header must be Bearer token", nil)
		return
	}

	accessToken := parts[1]
	user, err := h.useCase.GetUserProfile(r.Context(), accessToken)
	if err != nil {
		api.MapError(w, err)
		return
	}

	resp := dto.ProfileResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      string(user.Role),
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	api.WriteJSON(w, http.StatusOK, resp)
}
