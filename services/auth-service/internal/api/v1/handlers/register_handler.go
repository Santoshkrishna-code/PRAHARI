package handlers

import (
	"net/http"

	"prahari/services/auth-service/internal/api"
	"prahari/services/auth-service/internal/domain"
	"prahari/services/auth-service/internal/dto"
)

type RegisterHandler struct {
	useCase domain.AuthUseCase
}

func NewRegisterHandler(useCase domain.AuthUseCase) *RegisterHandler {
	return &RegisterHandler{useCase: useCase}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req dto.RegisterRequest
	if err := api.DecodeJSON(r, &req); err != nil {
		api.WriteError(w, http.StatusBadRequest, "INVALID_ARGUMENT", err.Error(), nil)
		return
	}

	user, err := h.useCase.Register(r.Context(), req.Email, req.Password, req.Role, req.FirstName, req.LastName)
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

	api.WriteJSON(w, http.StatusCreated, resp)
}
