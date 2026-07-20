package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	contractorApp "prahari/services/contractor/internal/application/contractor"
	onboardingApp "prahari/services/contractor/internal/application/onboarding"
	complianceApp "prahari/services/contractor/internal/application/compliance"
	siteaccessApp "prahari/services/contractor/internal/application/siteaccess"
	searchApp "prahari/services/contractor/internal/application/search"
	reportingApp "prahari/services/contractor/internal/application/reporting"
	exportApp "prahari/services/contractor/internal/application/export"
	searchDomain "prahari/services/contractor/internal/domain/search"
)

// Handler binds HTTP requests.
type Handler struct {
	contractor *contractorApp.Service
	onboarding *onboardingApp.Service
	compliance *complianceApp.Service
	siteaccess *siteaccessApp.Service
	search     *searchApp.Service
	reporting  *reportingApp.Service
	export     *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	contractor *contractorApp.Service,
	onboarding *onboardingApp.Service,
	compliance *complianceApp.Service,
	siteaccess *siteaccessApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		contractor: contractor,
		onboarding: onboarding,
		compliance: compliance,
		siteaccess: siteaccess,
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

// CreateContractor handles POST /contractors.
func (h *Handler) CreateContractor(w http.ResponseWriter, r *http.Request) {
	var cmd contractorApp.RegisterContractorCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	c, err := h.contractor.CreateContractor(r.Context(), cmd, "actor-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, c)
}

// GetContractor handles GET /contractors/{id}.
func (h *Handler) GetContractor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, err := h.contractor.GetContractor(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("contractor record not found", err))
		return
	}
	writeJSON(w, http.StatusOK, c)
}

// ListContractors handles GET /contractors.
func (h *Handler) ListContractors(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": []string{}, "total_count": 0})
}

// ApproveContractor handles POST /contractors/{id}/approve.
func (h *Handler) ApproveContractor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := contractorApp.TransitionStatusCommand{
		ContractorID: id,
		TargetCode:   "APPROVED",
		ActorID:      "supervisor-id",
	}
	if err := h.contractor.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "approved"})
}

// ActivateContractor handles POST /contractors/{id}/activate.
func (h *Handler) ActivateContractor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := contractorApp.TransitionStatusCommand{
		ContractorID: id,
		TargetCode:   "ACTIVE",
		ActorID:      "supervisor-id",
	}
	if err := h.contractor.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "active"})
}

// SuspendContractor handles POST /contractors/{id}/suspend.
func (h *Handler) SuspendContractor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := contractorApp.TransitionStatusCommand{
		ContractorID: id,
		TargetCode:   "SUSPENDED",
		ActorID:      "supervisor-id",
	}
	if err := h.contractor.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "suspended"})
}

// OffboardContractor handles POST /contractors/{id}/offboard.
func (h *Handler) OffboardContractor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := contractorApp.TransitionStatusCommand{
		ContractorID: id,
		TargetCode:   "OFFBOARDED",
		ActorID:      "supervisor-id",
	}
	if err := h.contractor.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "offboarded"})
}

// AddDocument handles POST /contractors/{id}/documents.
func (h *Handler) AddDocument(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "document registered"})
}

// UploadAttachment handles POST /contractors/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// SearchContractors handles POST /contractors/search.
func (h *Handler) SearchContractors(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Disposition", "attachment; filename=contractor.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=contractor_profile.pdf")
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
		"service": "contractor-service",
		"version": "1.0.0",
	})
}
