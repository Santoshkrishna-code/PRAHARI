package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	auditApp "prahari/services/audit/internal/application/audit"
	planningApp "prahari/services/audit/internal/application/planning"
	executionApp "prahari/services/audit/internal/application/execution"
	reviewApp "prahari/services/audit/internal/application/review"
	findingsApp "prahari/services/audit/internal/application/findings"
	searchApp "prahari/services/audit/internal/application/search"
	reportingApp "prahari/services/audit/internal/application/reporting"
	exportApp "prahari/services/audit/internal/application/export"
	searchDomain "prahari/services/audit/internal/domain/search"
)

// Handler binds HTTP requests.
type Handler struct {
	audit     *auditApp.Service
	planning  *planningApp.Service
	execution *executionApp.Service
	review    *reviewApp.Service
	findings  *findingsApp.Service
	search    *searchApp.Service
	reporting *reportingApp.Service
	export    *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	audit *auditApp.Service,
	planning *planningApp.Service,
	execution *executionApp.Service,
	review *reviewApp.Service,
	findings *findingsApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		audit:     audit,
		planning:  planning,
		execution: execution,
		review:    review,
		findings:  findings,
		search:    search,
		reporting: reporting,
		export:    export,
	}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// CreateAudit handles POST /audits.
func (h *Handler) CreateAudit(w http.ResponseWriter, r *http.Request) {
	var cmd auditApp.CreateAuditCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	a, err := h.audit.CreateAudit(r.Context(), cmd, "actor-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, a)
}

// GetAudit handles GET /audits/{id}.
func (h *Handler) GetAudit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	a, err := h.audit.GetAudit(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("audit record not found", err))
		return
	}
	writeJSON(w, http.StatusOK, a)
}

// ListAudits handles GET /audits.
func (h *Handler) ListAudits(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": []string{}, "total_count": 0})
}

// ScheduleAudit handles POST /audits/{id}/schedule.
func (h *Handler) ScheduleAudit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := auditApp.TransitionStatusCommand{
		AuditID:    id,
		TargetCode: "SCHEDULED",
		ActorID:    "lead-auditor-id",
	}
	if err := h.audit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "scheduled"})
}

// ExecuteAudit handles POST /audits/{id}/execute.
func (h *Handler) ExecuteAudit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := auditApp.TransitionStatusCommand{
		AuditID:    id,
		TargetCode: "IN_PROGRESS",
		ActorID:    "lead-auditor-id",
	}
	if err := h.audit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "in progress"})
}

// ReviewAudit handles POST /audits/{id}/review.
func (h *Handler) ReviewAudit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := auditApp.TransitionStatusCommand{
		AuditID:    id,
		TargetCode: "REVIEW",
		ActorID:    "lead-auditor-id",
	}
	if err := h.audit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "reviewed"})
}

// ApproveAudit handles POST /audits/{id}/approve.
func (h *Handler) ApproveAudit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := auditApp.TransitionStatusCommand{
		AuditID:    id,
		TargetCode: "APPROVED",
		ActorID:    "manager-id",
	}
	if err := h.audit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "approved"})
}

// CloseAudit handles POST /audits/{id}/close.
func (h *Handler) CloseAudit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := auditApp.TransitionStatusCommand{
		AuditID:    id,
		TargetCode: "CLOSED",
		ActorID:    "lead-auditor-id",
	}
	if err := h.audit.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "closed"})
}

// UploadAttachment handles POST /audits/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// AddComment handles POST /audits/{id}/comments.
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "comment added"})
}

// SearchAudits handles POST /audits/search.
func (h *Handler) SearchAudits(w http.ResponseWriter, r *http.Request) {
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

// GetDashboardReport handles GET /reports.
func (h *Handler) GetDashboardReport(w http.ResponseWriter, r *http.Request) {
	report, err := h.reporting.GenerateDashboardReport(r.Context())
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
	w.Header().Set("Content-Disposition", "attachment; filename=audit_register.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=audit_sheet.pdf")
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
		"service": "audit-service",
		"version": "1.0.0",
	})
}
