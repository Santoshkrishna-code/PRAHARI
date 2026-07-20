package handlers

import (
	"net"
	"net/http"
	"strings"

	"prahari/services/auth-service/internal/api"
	"prahari/services/auth-service/internal/domain"
	"prahari/services/auth-service/internal/dto"
)

type LoginHandler struct {
	useCase domain.AuthUseCase
}

func NewLoginHandler(useCase domain.AuthUseCase) *LoginHandler {
	return &LoginHandler{useCase: useCase}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req dto.LoginRequest
	if err := api.DecodeJSON(r, &req); err != nil {
		api.WriteError(w, http.StatusBadRequest, "INVALID_ARGUMENT", err.Error(), nil)
		return
	}

	// Extract IP and UserAgent context for safety audit events publishing
	ip := getClientIP(r)
	userAgent := r.Header.Get("User-Agent")

	tokens, err := h.useCase.Login(r.Context(), req.Email, req.Password, ip, userAgent)
	if err != nil {
		api.MapError(w, err)
		return
	}

	resp := dto.LoginResponse{
		AccessToken:  tokens.AccessToken,
		IDToken:      tokens.IDToken,
		RefreshToken: tokens.RefreshToken,
		ExpiresIn:    tokens.ExpiresIn,
	}

	api.WriteJSON(w, http.StatusOK, resp)
}

func getClientIP(r *http.Request) string {
	// Check standard proxy header first
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
