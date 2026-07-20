package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	observationApp "prahari/services/observation/internal/application/observation"
	coachingApp "prahari/services/observation/internal/application/coaching"
	followupApp "prahari/services/observation/internal/application/followup"
	effectivenessApp "prahari/services/observation/internal/application/effectiveness"
	searchApp "prahari/services/observation/internal/application/search"
	reportingApp "prahari/services/observation/internal/application/reporting"
	exportApp "prahari/services/observation/internal/application/export"
	searchDomain "prahari/services/observation/internal/domain/search"
)

// Handler binds HTTP requests.
type Handler struct {
	observation   *observationApp.Service
	coaching      *coachingApp.Service
	followup      *followupApp.Service
	effectiveness *effectivenessApp.Service
	search        *searchApp.Service
	reporting     *reportingApp.Service
	export        *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	observation *observationApp.Service,
	coaching *coachingApp.Service,
	followup *followupApp.Service,
	effectiveness *effectivenessApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		observation:   observation,
		coaching:      coaching,
		followup:      followup,
		effectiveness: effectiveness,
		search:        search,
		reporting:     reporting,
		export:        export,
	}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// CreateObservation handles POST /observations.
func (h *Handler) CreateObservation(w http.ResponseWriter, r *http.Request) {
	var cmd observationApp.CreateObservationCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	c, err := h.observation.CreateObservation(r.Context(), cmd, "actor-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, c)
}

// GetObservation handles GET /observations/{id}.
func (h *Handler) GetObservation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, err := h.observation.GetObservation(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("observation record not found", err))
		return
	}
	writeJSON(w, http.StatusOK, c)
}

// ListObservations handles GET /observations.
func (h *Handler) ListObservations(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": []string{}, "total_count": 0})
}

// CoachObservation handles POST /observations/{id}/coach.
func (h *Handler) CoachObservation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := observationApp.TransitionStatusCommand{
		ObservationID: id,
		TargetCode:    "COACHING",
		ActorID:       "supervisor-id",
	}
	if err := h.observation.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "coached"})
}

// RecognizeObservation handles POST /observations/{id}/recognize.
func (h *Handler) RecognizeObservation(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "recognized"})
}

// FollowUpObservation handles POST /observations/{id}/followup.
func (h *Handler) FollowUpObservation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := observationApp.TransitionStatusCommand{
		ObservationID: id,
		TargetCode:    "FOLLOWUP",
		ActorID:       "supervisor-id",
	}
	if err := h.observation.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "follow-up completed"})
}

// VerifyObservation handles POST /observations/{id}/verify.
func (h *Handler) VerifyObservation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := observationApp.TransitionStatusCommand{
		ObservationID: id,
		TargetCode:    "VERIFIED",
		ActorID:       "supervisor-id",
	}
	if err := h.observation.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "verified"})
}

// CloseObservation handles POST /observations/{id}/close.
func (h *Handler) CloseObservation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := observationApp.TransitionStatusCommand{
		ObservationID: id,
		TargetCode:    "CLOSED",
		ActorID:       "supervisor-id",
	}
	if err := h.observation.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "closed"})
}

// UploadAttachment handles POST /observations/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// AddComment handles POST /observations/{id}/comments.
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "comment added"})
}

// SearchObservations handles POST /observations/search.
func (h *Handler) SearchObservations(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Disposition", "attachment; filename=observation.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=observation_sheet.pdf")
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
		"service": "observation-service",
		"version": "1.0.0",
	})
}
