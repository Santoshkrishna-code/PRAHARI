package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	assessmentApp "prahari/services/risk/internal/application/assessment"
	reviewApp "prahari/services/risk/internal/application/review"
	approvalApp "prahari/services/risk/internal/application/approval"
	residualApp "prahari/services/risk/internal/application/residual"
	searchApp "prahari/services/risk/internal/application/search"
	reportingApp "prahari/services/risk/internal/application/reporting"
	exportApp "prahari/services/risk/internal/application/export"
	searchDomain "prahari/services/risk/internal/domain/search"
)

// Handler binds HTTP requests.
type Handler struct {
	assessment *assessmentApp.Service
	review     *reviewApp.Service
	approval   *approvalApp.Service
	residual   *residualApp.Service
	search     *searchApp.Service
	reporting  *reportingApp.Service
	export     *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	assessment *assessmentApp.Service,
	review *reviewApp.Service,
	approval *approvalApp.Service,
	residual *residualApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		assessment: assessment,
		review:     review,
		approval:   approval,
		residual:   residual,
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

// CreateRiskAssessment handles POST /risk-assessments.
func (h *Handler) CreateRiskAssessment(w http.ResponseWriter, r *http.Request) {
	var cmd assessmentApp.CreateRiskCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	c, err := h.assessment.CreateRiskAssessment(r.Context(), cmd, "actor-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, c)
}

// GetRiskAssessment handles GET /risk-assessments/{id}.
func (h *Handler) GetRiskAssessment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, err := h.assessment.GetRiskAssessment(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("risk assessment not found", err))
		return
	}
	writeJSON(w, http.StatusOK, c)
}

// ListRiskAssessments handles GET /risk-assessments.
func (h *Handler) ListRiskAssessments(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": []string{}, "total_count": 0})
}

// AssessRisk handles POST /risk-assessments/{id}/assess.
func (h *Handler) AssessRisk(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := assessmentApp.TransitionStatusCommand{
		RiskID:     id,
		TargetCode: "ASSESSMENT",
		ActorID:    "assessor-id",
	}
	if err := h.assessment.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "assessed"})
}

// ReviewRisk handles POST /risk-assessments/{id}/review.
func (h *Handler) ReviewRisk(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := assessmentApp.TransitionStatusCommand{
		RiskID:     id,
		TargetCode: "REVIEW",
		ActorID:    "assessor-id",
	}
	if err := h.assessment.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "reviewed"})
}

// ApproveRisk handles POST /risk-assessments/{id}/approve.
func (h *Handler) ApproveRisk(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := assessmentApp.TransitionStatusCommand{
		RiskID:     id,
		TargetCode: "APPROVAL",
		ActorID:    "manager-id",
	}
	if err := h.assessment.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "approved"})
}

// ActivateRisk handles POST /risk-assessments/{id}/activate.
func (h *Handler) ActivateRisk(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := assessmentApp.TransitionStatusCommand{
		RiskID:     id,
		TargetCode: "ACTIVE",
		ActorID:    "manager-id",
	}
	if err := h.assessment.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "active"})
}

// ReassessRisk handles POST /risk-assessments/{id}/reassess.
func (h *Handler) ReassessRisk(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := assessmentApp.TransitionStatusCommand{
		RiskID:     id,
		TargetCode: "REASSESSMENT",
		ActorID:    "assessor-id",
	}
	if err := h.assessment.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "reassessed"})
}

// UploadAttachment handles POST /risk-assessments/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// AddComment handles POST /risk-assessments/{id}/comments.
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "comment added"})
}

// SearchRiskAssessments handles POST /risk-assessments/search.
func (h *Handler) SearchRiskAssessments(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Disposition", "attachment; filename=risk_register.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=risk_assessment_sheet.pdf")
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
		"service": "risk-service",
		"version": "1.0.0",
	})
}
