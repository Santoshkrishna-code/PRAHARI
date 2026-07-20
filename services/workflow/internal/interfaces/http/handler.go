package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"
	prahariApproval "prahari/services/workflow/internal/application/approval"
)

// Handler maps REST endpoints controllers.
type Handler struct {
	approvalSvc *prahariApproval.Service
}

// NewHandler constructs a Handler.
func NewHandler(approvalSvc *prahariApproval.Service) *Handler {
	return &Handler{approvalSvc: approvalSvc}
}

// CompleteTask processes human/approval task completions.
func (h *Handler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	var body struct {
		TaskID     string `json:"task_id"`
		ApproverID string `json:"approver_id"`
		Decision   string `json:"decision"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request format", err))
		return
	}

	err := h.approvalSvc.SubmitDecision(r.Context(), body.TaskID, body.ApproverID, body.Decision)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to record task decision", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "decision recorded successfully",
	})
}
