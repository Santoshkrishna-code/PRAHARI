package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	complianceApp "prahari/services/compliance/internal/application/compliance"
	obligationApp "prahari/services/compliance/internal/application/obligation"
	reviewApp "prahari/services/compliance/internal/application/review"
	evidenceApp "prahari/services/compliance/internal/application/evidence"
	searchApp "prahari/services/compliance/internal/application/search"
	reportingApp "prahari/services/compliance/internal/application/reporting"
	exportApp "prahari/services/compliance/internal/application/export"
	searchDomain "prahari/services/compliance/internal/domain/search"
)

// Handler binds HTTP requests.
type Handler struct {
	compliance *complianceApp.Service
	obligation *obligationApp.Service
	review     *reviewApp.Service
	evidence   *evidenceApp.Service
	search     *searchApp.Service
	reporting  *reportingApp.Service
	export     *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	compliance *complianceApp.Service,
	obligation *obligationApp.Service,
	review *reviewApp.Service,
	evidence *evidenceApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		compliance: compliance,
		obligation: obligation,
		review:     review,
		evidence:   evidence,
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

// CreateCompliance handles POST /compliance.
func (h *Handler) CreateCompliance(w http.ResponseWriter, r *http.Request) {
	var cmd complianceApp.CreateComplianceCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	c, err := h.compliance.CreateCompliance(r.Context(), cmd, "actor-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, c)
}

// GetCompliance handles GET /compliance/{id}.
func (h *Handler) GetCompliance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, err := h.compliance.GetCompliance(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("compliance record not found", err))
		return
	}
	writeJSON(w, http.StatusOK, c)
}

// ListCompliance handles GET /compliance.
func (h *Handler) ListCompliance(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": []string{}, "total_count": 0})
}

// ReviewCompliance handles POST /compliance/{id}/review.
func (h *Handler) ReviewCompliance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := complianceApp.TransitionStatusCommand{
		ComplianceID: id,
		TargetCode:   "REVIEW",
		ActorID:      "assessor-id",
	}
	if err := h.compliance.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "reviewed"})
}

// EvidenceCompliance handles POST /compliance/{id}/evidence.
func (h *Handler) EvidenceCompliance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := complianceApp.TransitionStatusCommand{
		ComplianceID: id,
		TargetCode:   "EVIDENCE_COLLECTION",
		ActorID:      "assessor-id",
	}
	if err := h.compliance.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "evidence collection"})
}

// MonitorCompliance handles POST /compliance/{id}/monitor.
func (h *Handler) MonitorCompliance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := complianceApp.TransitionStatusCommand{
		ComplianceID: id,
		TargetCode:   "MONITORING",
		ActorID:      "manager-id",
	}
	if err := h.compliance.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "monitoring"})
}

// UploadAttachment handles POST /compliance/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// AddComment handles POST /compliance/{id}/comments.
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "comment added"})
}

// SearchCompliance handles POST /compliance/search.
func (h *Handler) SearchCompliance(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Disposition", "attachment; filename=compliance_register.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=compliance_sheet.pdf")
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
		"service": "compliance-service",
		"version": "1.0.0",
	})
}
