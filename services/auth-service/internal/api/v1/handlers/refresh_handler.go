package handlers

import (
	"net/http"

	"prahari/services/auth-service/internal/api"
	"prahari/services/auth-service/internal/domain"
	"prahari/services/auth-service/internal/dto"
)

type RefreshHandler struct {
	useCase domain.AuthUseCase
}

func NewRefreshHandler(useCase domain.AuthUseCase) *RefreshHandler {
	return &RefreshHandler{useCase: useCase}
}

func (h *RefreshHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req dto.RefreshRequest
	if err := api.DecodeJSON(r, &req); err != nil {
		api.WriteError(w, http.StatusBadRequest, "INVALID_ARGUMENT", err.Error(), nil)
		return
	}

	tokens, err := h.useCase.Refresh(r.Context(), req.RefreshToken)
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
