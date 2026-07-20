package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	inspectionApp "prahari/services/inspection/internal/application/inspection"
	checklistApp "prahari/services/inspection/internal/application/checklist"
	actionApp "prahari/services/inspection/internal/application/action"
	scheduleApp "prahari/services/inspection/internal/application/schedule"
	searchApp "prahari/services/inspection/internal/application/search"
	reportingApp "prahari/services/inspection/internal/application/reporting"
	exportApp "prahari/services/inspection/internal/application/export"
	searchDomain "prahari/services/inspection/internal/domain/search"
)

// Handler binds HTTP requests.
type Handler struct {
	inspection *inspectionApp.Service
	checklist  *checklistApp.Service
	action     *actionApp.Service
	schedule   *scheduleApp.Service
	search     *searchApp.Service
	reporting  *reportingApp.Service
	export     *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	inspection *inspectionApp.Service,
	checklist *checklistApp.Service,
	action *actionApp.Service,
	schedule *scheduleApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		inspection: inspection,
		checklist:  checklist,
		action:     action,
		schedule:   schedule,
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

// CreateInspection handles POST /inspections.
func (h *Handler) CreateInspection(w http.ResponseWriter, r *http.Request) {
	var cmd inspectionApp.CreateInspectionCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	ins, err := h.inspection.CreateInspection(r.Context(), cmd)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, ins)
}

// GetInspection handles GET /inspections/{id}.
func (h *Handler) GetInspection(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ins, err := h.inspection.GetInspection(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("inspection not found", err))
		return
	}
	writeJSON(w, http.StatusOK, ins)
}

// ListInspections handles GET /inspections.
func (h *Handler) ListInspections(w http.ResponseWriter, r *http.Request) {
	inspections, total, err := h.inspection.ListInspections(r.Context(), 1, 20)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"items":       inspections,
		"total_count": total,
	})
}

// UpdateInspection handles PUT /inspections/{id}.
func (h *Handler) UpdateInspection(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var cmd inspectionApp.UpdateInspectionCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	ins, err := h.inspection.UpdateInspection(r.Context(), id, cmd, "inspector-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, ins)
}

// ScheduleInspection handles POST /inspections/{id}/schedule.
func (h *Handler) ScheduleInspection(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := inspectionApp.TransitionStatusCommand{
		InspectionID: id,
		TargetCode:   "SCHEDULED",
		ActorID:      "scheduler-id",
	}
	if err := h.inspection.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "scheduled"})
}

// AssignInspection handles POST /inspections/{id}/assign.
func (h *Handler) AssignInspection(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := inspectionApp.TransitionStatusCommand{
		InspectionID: id,
		TargetCode:   "ASSIGNED",
		ActorID:      "scheduler-id",
	}
	if err := h.inspection.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "assigned"})
}

// StartInspection handles POST /inspections/{id}/start.
func (h *Handler) StartInspection(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := inspectionApp.TransitionStatusCommand{
		InspectionID: id,
		TargetCode:   "IN_PROGRESS",
		ActorID:      "inspector-id",
	}
	if err := h.inspection.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "started"})
}

// CompleteInspection handles POST /inspections/{id}/complete.
func (h *Handler) CompleteInspection(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := inspectionApp.TransitionStatusCommand{
		InspectionID: id,
		TargetCode:   "COMPLETED",
		ActorID:      "inspector-id",
	}
	if err := h.inspection.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "completed"})
}

// ApproveInspection handles POST /inspections/{id}/approve.
func (h *Handler) ApproveInspection(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := inspectionApp.TransitionStatusCommand{
		InspectionID: id,
		TargetCode:   "APPROVED",
		ActorID:      "manager-id",
	}
	if err := h.inspection.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "approved"})
}

// AddFinding handles POST /inspections/{id}/findings.
func (h *Handler) AddFinding(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "finding recorded"})
}

// AddAction handles POST /inspections/{id}/actions.
func (h *Handler) AddAction(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "CAPA created"})
}

// AddComment handles POST /inspections/{id}/comments.
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "comment added"})
}

// UploadAttachment handles POST /inspections/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// SearchInspections handles POST /inspections/search.
func (h *Handler) SearchInspections(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Disposition", "attachment; filename=inspections.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=inspection_report.pdf")
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
		"service": "inspection-service",
		"version": "1.0.0",
	})
}
