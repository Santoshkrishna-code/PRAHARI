package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"
	prahariFlow "prahari/services/notification/internal/application/dispatcher"
)

// Handler maps REST endpoints controllers.
type Handler struct {
	flow *prahariFlow.Flow
}

// NewHandler constructs a Handler.
func NewHandler(flow *prahariFlow.Flow) *Handler {
	return &Handler{flow: flow}
}

// Send dispatches outbound messages.
func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Recipient string `json:"recipient"`
		Channel   string `json:"channel"`
		Message   string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request format", err))
		return
	}

	err := h.flow.Dispatch(r.Context(), body.Channel, body.Recipient, body.Message)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to dispatch messaging", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "notification sent successfully",
	})
}
