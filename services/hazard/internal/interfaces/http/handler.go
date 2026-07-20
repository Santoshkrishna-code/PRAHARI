package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	hazardApp "prahari/services/hazard/internal/application/hazard"
	assessmentApp "prahari/services/hazard/internal/application/assessment"
	mitigationApp "prahari/services/hazard/internal/application/mitigation"
	verifyApp "prahari/services/hazard/internal/application/verification"
	searchApp "prahari/services/hazard/internal/application/search"
	reportingApp "prahari/services/hazard/internal/application/reporting"
	exportApp "prahari/services/hazard/internal/application/export"
	searchDomain "prahari/services/hazard/internal/domain/search"
)

// Handler binds HTTP requests.
type Handler struct {
	hazard     *hazardApp.Service
	assessment *assessmentApp.Service
	mitigation *mitigationApp.Service
	verify     *verifyApp.Service
	search     *searchApp.Service
	reporting  *reportingApp.Service
	export     *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	hazard *hazardApp.Service,
	assessment *assessmentApp.Service,
	mitigation *mitigationApp.Service,
	verify *verifyApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		hazard:     hazard,
		assessment: assessment,
		mitigation: mitigation,
		verify:     verify,
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

// CreateHazard handles POST /hazards.
func (h *Handler) CreateHazard(w http.ResponseWriter, r *http.Request) {
	var cmd hazardApp.CreateHazardCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	c, err := h.hazard.CreateHazard(r.Context(), cmd, "actor-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, c)
}

// GetHazard handles GET /hazards/{id}.
func (h *Handler) GetHazard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, err := h.hazard.GetHazard(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("hazard record not found", err))
		return
	}
	writeJSON(w, http.StatusOK, c)
}

// ListHazards handles GET /hazards.
func (h *Handler) ListHazards(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": []string{}, "total_count": 0})
}

// AssessHazard handles POST /hazards/{id}/assess.
func (h *Handler) AssessHazard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := hazardApp.TransitionStatusCommand{
		HazardID:   id,
		TargetCode: "RISK_ASSESSMENT",
		ActorID:    "supervisor-id",
	}
	if err := h.hazard.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "assessed"})
}

// MitigateHazard handles POST /hazards/{id}/mitigate.
func (h *Handler) MitigateHazard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := hazardApp.TransitionStatusCommand{
		HazardID:   id,
		TargetCode: "MITIGATION_PLANNING",
		ActorID:    "supervisor-id",
	}
	if err := h.hazard.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "mitigation planned"})
}

// VerifyHazard handles POST /hazards/{id}/verify.
func (h *Handler) VerifyHazard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := hazardApp.TransitionStatusCommand{
		HazardID:   id,
		TargetCode: "VERIFICATION",
		ActorID:    "supervisor-id",
	}
	if err := h.hazard.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "verified"})
}

// CloseHazard handles POST /hazards/{id}/close.
func (h *Handler) CloseHazard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := hazardApp.TransitionStatusCommand{
		HazardID:   id,
		TargetCode: "CLOSED",
		ActorID:    "supervisor-id",
	}
	if err := h.hazard.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "closed"})
}

// UploadAttachment handles POST /hazards/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// AddComment handles POST /hazards/{id}/comments.
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "comment added"})
}

// SearchHazards handles POST /hazards/search.
func (h *Handler) SearchHazards(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Disposition", "attachment; filename=hazard.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=hazard_assessment.pdf")
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
		"service": "hazard-service",
		"version": "1.0.0",
	})
}
