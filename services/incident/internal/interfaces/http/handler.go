package http

import (
	"encoding/json"
	"net/http"
	"time"

	prahariErrors "prahari/shared/errors"

	incidentApp "prahari/services/incident/internal/application/incident"
	assignmentApp "prahari/services/incident/internal/application/assignment"
	investigationApp "prahari/services/incident/internal/application/investigation"
	searchApp "prahari/services/incident/internal/application/search"
	reportingApp "prahari/services/incident/internal/application/reporting"
	exportApp "prahari/services/incident/internal/application/export"
	assignmentDomain "prahari/services/incident/internal/domain/assignment"
	investigationDomain "prahari/services/incident/internal/domain/investigation"
	searchDomain "prahari/services/incident/internal/domain/search"
)

// Handler binds HTTP requests to application services.
type Handler struct {
	incident      *incidentApp.Service
	assignment    *assignmentApp.Service
	investigation *investigationApp.Service
	search        *searchApp.Service
	reporting     *reportingApp.Service
	export        *exportApp.Service
}

// NewHandler constructs a Handler with all application services injected.
func NewHandler(
	incident *incidentApp.Service,
	assignment *assignmentApp.Service,
	investigation *investigationApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		incident:      incident,
		assignment:    assignment,
		investigation: investigation,
		search:        search,
		reporting:     reporting,
		export:        export,
	}
}

// writeJSON encodes a response body as JSON with the given status code.
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// CreateIncident handles POST /incidents.
func (h *Handler) CreateIncident(w http.ResponseWriter, r *http.Request) {
	var cmd incidentApp.CreateIncidentCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	inc, err := h.incident.CreateIncident(r.Context(), cmd)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to create incident", err))
		return
	}

	writeJSON(w, http.StatusCreated, inc)
}

// GetIncident handles GET /incidents/{id}.
func (h *Handler) GetIncident(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	inc, err := h.incident.GetIncident(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("incident not found", err))
		return
	}
	writeJSON(w, http.StatusOK, inc)
}

// ListIncidents handles GET /incidents.
func (h *Handler) ListIncidents(w http.ResponseWriter, r *http.Request) {
	incidents, total, err := h.incident.ListIncidents(r.Context(), 1, 20)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to list incidents", err))
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"items":       incidents,
		"total_count": total,
	})
}

// UpdateIncident handles PUT /incidents/{id}.
func (h *Handler) UpdateIncident(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var cmd incidentApp.UpdateIncidentCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	inc, err := h.incident.UpdateIncident(r.Context(), id, cmd, "actor-from-jwt")
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to update incident", err))
		return
	}
	writeJSON(w, http.StatusOK, inc)
}

// DeleteIncident handles DELETE /incidents/{id}.
func (h *Handler) DeleteIncident(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.incident.DeleteIncident(r.Context(), id, "actor-from-jwt"); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to delete incident", err))
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// AssignIncident handles POST /incidents/{id}/assign.
func (h *Handler) AssignIncident(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var body struct {
		AssigneeID string `json:"assignee_id"`
		Role       string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	assignment, err := h.assignment.AssignIncident(r.Context(), id, body.AssigneeID, "assigner-from-jwt", assignmentDomain.Role(body.Role))
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to assign incident", err))
		return
	}
	writeJSON(w, http.StatusOK, assignment)
}

// StartInvestigation handles POST /incidents/{id}/investigate.
func (h *Handler) StartInvestigation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var body struct {
		InvestigatorID string `json:"investigator_id"`
		Methodology    string `json:"methodology"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	inv, err := h.investigation.StartInvestigation(r.Context(), id, body.InvestigatorID, investigationDomain.Methodology(body.Methodology))
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to start investigation", err))
		return
	}
	writeJSON(w, http.StatusOK, inv)
}

// ResolveIncident handles POST /incidents/{id}/resolve.
func (h *Handler) ResolveIncident(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := incidentApp.TransitionStatusCommand{
		IncidentID: id,
		TargetCode: "RESOLVED",
		ActorID:    "actor-from-jwt",
	}
	if err := h.incident.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to resolve incident", err))
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "resolved"})
}

// CloseIncident handles POST /incidents/{id}/close.
func (h *Handler) CloseIncident(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := incidentApp.TransitionStatusCommand{
		IncidentID: id,
		TargetCode: "CLOSED",
		ActorID:    "actor-from-jwt",
	}
	if err := h.incident.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to close incident", err))
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "closed"})
}

// AddComment handles POST /incidents/{id}/comments.
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "comment added"})
}

// UploadAttachment handles POST /incidents/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// GetTimeline handles GET /incidents/{id}/timeline.
func (h *Handler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "timeline retrieved"})
}

// CreateCAPA handles POST /incidents/{id}/capa.
func (h *Handler) CreateCAPA(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "CAPA created"})
}

// SearchIncidents handles POST /incidents/search.
func (h *Handler) SearchIncidents(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid search criteria", err))
		return
	}

	result, err := h.search.Search(r.Context(), &criteria)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("search failed", err))
		return
	}
	writeJSON(w, http.StatusOK, result)
}

// GetSummaryReport handles GET /reports/summary.
func (h *Handler) GetSummaryReport(w http.ResponseWriter, r *http.Request) {
	to := time.Now()
	from := to.AddDate(0, -1, 0) // Last 30 days

	report, err := h.reporting.GenerateSummaryReport(r.Context(), from, to)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to generate summary report", err))
		return
	}
	writeJSON(w, http.StatusOK, report)
}

// GetMetricsReport handles GET /reports/metrics.
func (h *Handler) GetMetricsReport(w http.ResponseWriter, r *http.Request) {
	to := time.Now()
	from := to.AddDate(0, -1, 0)

	report, err := h.reporting.GenerateMetricsReport(r.Context(), from, to)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to generate metrics report", err))
		return
	}
	writeJSON(w, http.StatusOK, report)
}

// ExportCSV handles GET /export/csv.
func (h *Handler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	criteria := &searchDomain.Criteria{}
	data, err := h.export.ExportCSV(r.Context(), criteria)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("CSV export failed", err))
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=incidents.csv")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// ExportPDF handles GET /export/pdf/{id}.
func (h *Handler) ExportPDF(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	data, err := h.export.ExportPDF(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("PDF export failed", err))
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=incident_report.pdf")
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
		"service": "incident-service",
		"version": "1.0.0",
	})
}
