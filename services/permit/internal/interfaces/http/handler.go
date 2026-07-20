package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	permitApp "prahari/services/permit/internal/application/permit"
	approvalApp "prahari/services/permit/internal/application/approval"
	riskApp "prahari/services/permit/internal/application/riskassessment"
	searchApp "prahari/services/permit/internal/application/search"
	reportingApp "prahari/services/permit/internal/application/reporting"
	exportApp "prahari/services/permit/internal/application/export"
	approvalDomain "prahari/services/permit/internal/domain/approval"
	searchDomain "prahari/services/permit/internal/domain/search"
)

// Handler delegates incoming requests to target application services.
type Handler struct {
	permit     *permitApp.Service
	approval   *approvalApp.Service
	risk       *riskApp.Service
	search     *searchApp.Service
	reporting  *reportingApp.Service
	export     *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	permit *permitApp.Service,
	approval *approvalApp.Service,
	risk *riskApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		permit:     permit,
		approval:   approval,
		risk:       risk,
		search:     search,
		reporting:  reporting,
		export:     export,
	}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// CreatePermit handles POST /permits.
func (h *Handler) CreatePermit(w http.ResponseWriter, r *http.Request) {
	var cmd permitApp.CreateIncidentCommand // Using CreatePermitCommand equivalent
	var realCmd permitApp.CreatePermitCommand
	if err := json.NewDecoder(r.Body).Decode(&realCmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	p, err := h.permit.CreatePermit(r.Context(), realCmd)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to create permit", err))
		return
	}
	writeJSON(w, http.StatusCreated, p)
}

// GetPermit handles GET /permits/{id}.
func (h *Handler) GetPermit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	p, err := h.permit.GetPermit(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("permit not found", err))
		return
	}
	writeJSON(w, http.StatusOK, p)
}

// ListPermits handles GET /permits.
func (h *Handler) ListPermits(w http.ResponseWriter, r *http.Request) {
	permits, total, err := h.permit.ListPermits(r.Context(), 1, 20)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"items":       permits,
		"total_count": total,
	})
}

// UpdatePermit handles PUT /permits/{id}.
func (h *Handler) UpdatePermit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var cmd permitApp.UpdatePermitCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	p, err := h.permit.UpdatePermit(r.Context(), id, cmd, "actor-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, p)
}

// SubmitPermit handles POST /permits/{id}/submit.
func (h *Handler) SubmitPermit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := permitApp.TransitionStatusCommand{
		PermitID:   id,
		TargetCode: "SUBMITTED",
		ActorID:    "applicant-id",
	}
	if err := h.permit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "submitted"})
}

// ApprovePermit handles POST /permits/{id}/approve.
func (h *Handler) ApprovePermit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var body struct {
		ApproverID string `json:"approver_id"`
		Role       string `json:"role"`
		Signature  string `json:"signature"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	app, err := h.approval.SubmitApproval(r.Context(), id, body.ApproverID, approvalDomain.Role(body.Role), approvalDomain.DecisionApproved, "approved", body.Signature)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, app)
}

// RejectPermit handles POST /permits/{id}/reject.
func (h *Handler) RejectPermit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := permitApp.TransitionStatusCommand{
		PermitID:   id,
		TargetCode: "REJECTED",
		ActorID:    "reviewer-id",
	}
	if err := h.permit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "rejected"})
}

// IssuePermit handles POST /permits/{id}/issue.
func (h *Handler) IssuePermit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := permitApp.TransitionStatusCommand{
		PermitID:   id,
		TargetCode: "ISSUED",
		ActorID:    "issuer-id",
	}
	if err := h.permit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "issued"})
}

// ActivatePermit handles POST /permits/{id}/activate.
func (h *Handler) ActivatePermit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := permitApp.TransitionStatusCommand{
		PermitID:   id,
		TargetCode: "ACTIVE",
		ActorID:    "receiver-id",
	}
	if err := h.permit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "activated"})
}

// SuspendPermit handles POST /permits/{id}/suspend.
func (h *Handler) SuspendPermit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := permitApp.TransitionStatusCommand{
		PermitID:   id,
		TargetCode: "SUSPENDED",
		ActorID:    "safety-officer-id",
	}
	if err := h.permit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "suspended"})
}

// CompletePermit handles POST /permits/{id}/complete.
func (h *Handler) CompletePermit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := permitApp.TransitionStatusCommand{
		PermitID:   id,
		TargetCode: "COMPLETED",
		ActorID:    "receiver-id",
	}
	if err := h.permit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "completed"})
}

// ClosePermit handles POST /permits/{id}/close.
func (h *Handler) ClosePermit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := permitApp.TransitionStatusCommand{
		PermitID:   id,
		TargetCode: "CLOSED",
		ActorID:    "issuer-id",
	}
	if err := h.permit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "closed"})
}

// ExtendPermit handles POST /permits/{id}/extend.
func (h *Handler) ExtendPermit(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "extended"})
}

// AddComment handles POST /permits/{id}/comments.
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "comment added"})
}

// UploadAttachment handles POST /permits/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// SearchPermits handles POST /permits/search.
func (h *Handler) SearchPermits(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	res, err := h.search.Search(r.Context(), &criteria)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, res)
}

// GetMetricsReport handles GET /reports.
func (h *Handler) GetMetricsReport(w http.ResponseWriter, r *http.Request) {
	report, err := h.reporting.GenerateMetricsReport(r.Context())
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, report)
}

// ExportCSV handles GET /export/csv.
func (h *Handler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	criteria := &searchDomain.Criteria{}
	data, err := h.export.ExportCSV(r.Context(), criteria)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=permits.csv")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// ExportPDF handles GET /export/pdf/{id}.
func (h *Handler) ExportPDF(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	data, err := h.export.ExportPDF(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=permit_sheet.pdf")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// Health handles GET /health.
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
}

// Ready handles GET /ready.
func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

// Live handles GET /live.
func (h *Handler) Live(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "alive"})
}

// Version handles GET /version.
func (h *Handler) Version(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"service": "permit-service",
		"version": "1.0.0",
	})
}
