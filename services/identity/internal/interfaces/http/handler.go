package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"
	"prahari/services/identity/internal/application"
)

// Handler maps REST endpoints controllers.
type Handler struct {
	authSvc *application.AuthService
	userSvc *application.UserService
}

// NewHandler constructs a Handler.
func NewHandler(authSvc *application.AuthService, userSvc *application.UserService) *Handler {
	return &Handler{
		authSvc: authSvc,
		userSvc: userSvc,
	}
}

// Login parses email/password, returning access tokens.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request format", err))
		return
	}

	token, err := h.authSvc.Authenticate(r.Context(), body.Email, body.Password)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewUnauthorizedError("invalid email or password", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"access_token": token,
		"token_type":   "Bearer",
	})
}

// Register parses user registration payloads.
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request format", err))
		return
	}

	user, err := h.userSvc.CreateUser(r.Context(), body.Email, body.Password)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to register user profile", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
