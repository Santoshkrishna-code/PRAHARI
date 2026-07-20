package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	trainingApp "prahari/services/training/internal/application/training"
	enrollmentApp "prahari/services/training/internal/application/enrollment"
	competencyApp "prahari/services/training/internal/application/competency"
	assessmentApp "prahari/services/training/internal/application/assessment"
	certificationApp "prahari/services/training/internal/application/certification"
	searchApp "prahari/services/training/internal/application/search"
	reportingApp "prahari/services/training/internal/application/reporting"
	exportApp "prahari/services/training/internal/application/export"
	searchDomain "prahari/services/training/internal/domain/search"
)

// Handler binds HTTP requests.
type Handler struct {
	training      *trainingApp.Service
	enrollment    *enrollmentApp.Service
	competency    *competencyApp.Service
	assessment    *assessmentApp.Service
	certification *certificationApp.Service
	search        *searchApp.Service
	reporting     *reportingApp.Service
	export        *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	training *trainingApp.Service,
	enrollment *enrollmentApp.Service,
	competency *competencyApp.Service,
	assessment *assessmentApp.Service,
	certification *certificationApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		training:      training,
		enrollment:    enrollment,
		competency:    competency,
		assessment:    assessment,
		certification: certification,
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

// CreateTraining handles POST /training.
func (h *Handler) CreateTraining(w http.ResponseWriter, r *http.Request) {
	var cmd trainingApp.CreateTrainingCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	t, err := h.training.CreateTraining(r.Context(), cmd, "actor-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, t)
}

// GetTraining handles GET /training/{id}.
func (h *Handler) GetTraining(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	t, err := h.training.GetTraining(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("training record not found", err))
		return
	}
	writeJSON(w, http.StatusOK, t)
}

// ListTrainings handles GET /training.
func (h *Handler) ListTrainings(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": []string{}, "total_count": 0})
}

// ScheduleTraining handles POST /training/{id}/schedule.
func (h *Handler) ScheduleTraining(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := trainingApp.TransitionStatusCommand{
		TrainingID: id,
		TargetCode: "SCHEDULED",
		ActorID:    "lead-instructor-id",
	}
	if err := h.training.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "scheduled"})
}

// EnrollTrainee handles POST /training/{id}/enroll.
func (h *Handler) EnrollTrainee(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := trainingApp.TransitionStatusCommand{
		TrainingID: id,
		TargetCode: "ENROLLMENT",
		ActorID:    "lead-instructor-id",
	}
	if err := h.training.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "enrollment"})
}

// RecordAttendance handles POST /training/{id}/attendance.
func (h *Handler) RecordAttendance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := trainingApp.TransitionStatusCommand{
		TrainingID: id,
		TargetCode: "IN_PROGRESS",
		ActorID:    "lead-instructor-id",
	}
	if err := h.training.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "attendance recorded"})
}

// AssessTraining handles POST /training/{id}/assessment.
func (h *Handler) AssessTraining(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := trainingApp.TransitionStatusCommand{
		TrainingID: id,
		TargetCode: "ASSESSMENT",
		ActorID:    "lead-instructor-id",
	}
	if err := h.training.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "assessment recorded"})
}

// CertifyTraining handles POST /training/{id}/certify.
func (h *Handler) CertifyTraining(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := trainingApp.TransitionStatusCommand{
		TrainingID: id,
		TargetCode: "CERTIFIED",
		ActorID:    "manager-id",
	}
	if err := h.training.TransitionStatus(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "certified"})
}

// UploadAttachment handles POST /training/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// AddComment handles POST /training/{id}/comments.
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "comment added"})
}

// SearchTraining handles POST /training/search.
func (h *Handler) SearchTraining(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Disposition", "attachment; filename=training_register.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=training_sheet.pdf")
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
		"service": "training-service",
		"version": "1.0.0",
	})
}
