package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	maintenanceApp "prahari/services/maintenance/internal/application/maintenance"
	workorderApp "prahari/services/maintenance/internal/application/workorder"
	schedulingApp "prahari/services/maintenance/internal/application/scheduling"
	planningApp "prahari/services/maintenance/internal/application/planning"
	searchApp "prahari/services/maintenance/internal/application/search"
	reportingApp "prahari/services/maintenance/internal/application/reporting"
	exportApp "prahari/services/maintenance/internal/application/export"
	searchDomain "prahari/services/maintenance/internal/domain/search"
)

// Handler binds HTTP requests.
type Handler struct {
	maintenance *maintenanceApp.Service
	workorder   *workorderApp.Service
	scheduling  *schedulingApp.Service
	planning    *planningApp.Service
	search      *searchApp.Service
	reporting   *reportingApp.Service
	export      *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	maintenance *maintenanceApp.Service,
	workorder *workorderApp.Service,
	scheduling *schedulingApp.Service,
	planning *planningApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		maintenance: maintenance,
		workorder:   workorder,
		scheduling:  scheduling,
		planning:    planning,
		search:      search,
		reporting:   reporting,
		export:      export,
	}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// CreateMaintenance handles POST /maintenance.
func (h *Handler) CreateMaintenance(w http.ResponseWriter, r *http.Request) {
	var cmd maintenanceApp.CreateMaintenanceCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	mnt, err := h.maintenance.CreateMaintenance(r.Context(), cmd, "actor-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, mnt)
}

// GetMaintenance handles GET /maintenance/{id}.
func (h *Handler) GetMaintenance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	mnt, err := h.maintenance.GetMaintenance(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("maintenance record not found", err))
		return
	}
	writeJSON(w, http.StatusOK, mnt)
}

// ListMaintenance handles GET /maintenance.
func (h *Handler) ListMaintenance(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": []string{}, "total_count": 0})
}

// ApproveMaintenance handles POST /maintenance/{id}/approve.
func (h *Handler) ApproveMaintenance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := maintenanceApp.TransitionStatusCommand{
		MaintenanceID: id,
		TargetCode:    "APPROVED",
		ActorID:       "manager-id",
	}
	if err := h.maintenance.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "approved"})
}

// ScheduleMaintenance handles POST /maintenance/{id}/schedule.
func (h *Handler) ScheduleMaintenance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := maintenanceApp.TransitionStatusCommand{
		MaintenanceID: id,
		TargetCode:    "SCHEDULED",
		ActorID:       "scheduler-id",
	}
	if err := h.maintenance.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "scheduled"})
}

// AssignMaintenance handles POST /maintenance/{id}/assign.
func (h *Handler) AssignMaintenance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := maintenanceApp.TransitionStatusCommand{
		MaintenanceID: id,
		TargetCode:    "ASSIGNED",
		ActorID:       "scheduler-id",
	}
	if err := h.maintenance.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "assigned"})
}

// StartMaintenance handles POST /maintenance/{id}/start.
func (h *Handler) StartMaintenance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := maintenanceApp.TransitionStatusCommand{
		MaintenanceID: id,
		TargetCode:    "IN_PROGRESS",
		ActorID:       "technician-id",
	}
	if err := h.maintenance.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "started"})
}

// CompleteMaintenance handles POST /maintenance/{id}/complete.
func (h *Handler) CompleteMaintenance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := maintenanceApp.TransitionStatusCommand{
		MaintenanceID: id,
		TargetCode:    "COMPLETED",
		ActorID:       "technician-id",
	}
	if err := h.maintenance.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "completed"})
}

// VerifyMaintenance handles POST /maintenance/{id}/verify.
func (h *Handler) VerifyMaintenance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := maintenanceApp.TransitionStatusCommand{
		MaintenanceID: id,
		TargetCode:    "VERIFIED",
		ActorID:       "inspector-id",
	}
	if err := h.maintenance.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "verified"})
}

// CloseMaintenance handles POST /maintenance/{id}/close.
func (h *Handler) CloseMaintenance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := maintenanceApp.TransitionStatusCommand{
		MaintenanceID: id,
		TargetCode:    "CLOSED",
		ActorID:       "manager-id",
	}
	if err := h.maintenance.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "closed"})
}

// UploadAttachment handles POST /maintenance/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// AddComment handles POST /maintenance/{id}/comments.
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "comment added"})
}

// SearchMaintenance handles POST /maintenance/search.
func (h *Handler) SearchMaintenance(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Disposition", "attachment; filename=maintenance.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=maintenance_workorder.pdf")
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
		"service": "maintenance-service",
		"version": "1.0.0",
	})
}
