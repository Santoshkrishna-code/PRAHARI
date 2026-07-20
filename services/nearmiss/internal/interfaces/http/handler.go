package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	nearmissApp "prahari/services/nearmiss/internal/application/nearmiss"
	investigationApp "prahari/services/nearmiss/internal/application/investigation"
	correctiveApp "prahari/services/nearmiss/internal/application/corrective"
	verifyApp "prahari/services/nearmiss/internal/application/verification"
	searchApp "prahari/services/nearmiss/internal/application/search"
	reportingApp "prahari/services/nearmiss/internal/application/reporting"
	exportApp "prahari/services/nearmiss/internal/application/export"
	searchDomain "prahari/services/nearmiss/internal/domain/search"
)

// Handler binds HTTP requests.
type Handler struct {
	nearmiss      *nearmissApp.Service
	investigation *investigationApp.Service
	corrective    *correctiveApp.Service
	verify        *verifyApp.Service
	search        *searchApp.Service
	reporting     *reportingApp.Service
	export        *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	nearmiss *nearmissApp.Service,
	investigation *investigationApp.Service,
	corrective *correctiveApp.Service,
	verify *verifyApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		nearmiss:      nearmiss,
		investigation: investigation,
		corrective:    corrective,
		verify:        verify,
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

// CreateNearMiss handles POST /near-misses.
func (h *Handler) CreateNearMiss(w http.ResponseWriter, r *http.Request) {
	var cmd nearmissApp.CreateNearMissCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	c, err := h.nearmiss.CreateNearMiss(r.Context(), cmd, "actor-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, c)
}

// GetNearMiss handles GET /near-misses/{id}.
func (h *Handler) GetNearMiss(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, err := h.nearmiss.GetNearMiss(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("near miss not found", err))
		return
	}
	writeJSON(w, http.StatusOK, c)
}

// ListNearMisses handles GET /near-misses.
func (h *Handler) ListNearMisses(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": []string{}, "total_count": 0})
}

// InvestigateNearMiss handles POST /near-misses/{id}/investigate.
func (h *Handler) InvestigateNearMiss(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := nearmissApp.TransitionStatusCommand{
		NearMissID: id,
		TargetCode: "INVESTIGATION",
		ActorID:    "investigator-id",
	}
	if err := h.nearmiss.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "investigation started"})
}

// AddCorrectiveAction handles POST /near-misses/{id}/corrective-actions.
func (h *Handler) AddCorrectiveAction(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := nearmissApp.TransitionStatusCommand{
		NearMissID: id,
		TargetCode: "CORRECTIVE_ACTIONS",
		ActorID:    "investigator-id",
	}
	if err := h.nearmiss.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "corrective actions planned"})
}

// VerifyNearMiss handles POST /near-misses/{id}/verify.
func (h *Handler) VerifyNearMiss(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := nearmissApp.TransitionStatusCommand{
		NearMissID: id,
		TargetCode: "VERIFICATION",
		ActorID:    "investigator-id",
	}
	if err := h.nearmiss.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "verified"})
}

// CloseNearMiss handles POST /near-misses/{id}/close.
func (h *Handler) CloseNearMiss(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := nearmissApp.TransitionStatusCommand{
		NearMissID: id,
		TargetCode: "CLOSED",
		ActorID:    "investigator-id",
	}
	if err := h.nearmiss.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "closed"})
}

// EscalateNearMiss handles POST /near-misses/{id}/escalate.
func (h *Handler) EscalateNearMiss(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := nearmissApp.TransitionStatusCommand{
		NearMissID: id,
		TargetCode: "ESCALATED",
		ActorID:    "investigator-id",
	}
	if err := h.nearmiss.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "escalated"})
}

// UploadAttachment handles POST /near-misses/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// AddComment handles POST /near-misses/{id}/comments.
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "comment added"})
}

// SearchNearMisses handles POST /near-misses/search.
func (h *Handler) SearchNearMisses(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Disposition", "attachment; filename=nearmiss.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=nearmiss_report.pdf")
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
		"service": "nearmiss-service",
		"version": "1.0.0",
	})
}
