package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"
	"prahari/templates/microservice/internal/application"
)

// Handler maps HTTP routes controllers.
type Handler struct {
	svc *application.IncidentService
}

// NewHandler constructs a Handler.
func NewHandler(svc *application.IncidentService) *Handler {
	return &Handler{svc: svc}
}

// Create handles POST requests to instantiate incidents.
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	incident, err := h.svc.CreateIncident(r.Context(), body.Title)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to create incident", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(incident)
}
