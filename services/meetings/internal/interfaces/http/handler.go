package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/meetings/internal/application/analytics"
	attendanceApp "prahari/services/meetings/internal/application/attendance"
	closureApp "prahari/services/meetings/internal/application/closure"
	conductApp "prahari/services/meetings/internal/application/conduct"
	exportApp "prahari/services/meetings/internal/application/export"
	reportingApp "prahari/services/meetings/internal/application/reporting"
	schedulingApp "prahari/services/meetings/internal/application/scheduling"
	searchApp "prahari/services/meetings/internal/application/search"
	attendanceDomain "prahari/services/meetings/internal/domain/attendance"
	"prahari/services/meetings/internal/domain/meeting"
	"prahari/services/meetings/internal/domain/minutes"
	searchDomain "prahari/services/meetings/internal/domain/search"
)

type Handler struct {
	schedulingSvc  *schedulingApp.Service
	conductSvc     *conductApp.Service
	attendanceSvc  *attendanceApp.Service
	closureSvc     *closureApp.Service
	reportingSvc   *reportingApp.Service
	analyticsSvc   *analyticsApp.Service
	searchSvc      *searchApp.Service
	exportSvc      *exportApp.Service
}

func NewHandler(
	schedulingSvc *schedulingApp.Service,
	conductSvc *conductApp.Service,
	attendanceSvc *attendanceApp.Service,
	closureSvc *closureApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		schedulingSvc:  schedulingSvc,
		conductSvc:     conductSvc,
		attendanceSvc:  attendanceSvc,
		closureSvc:     closureSvc,
		reportingSvc:   reportingSvc,
		analyticsSvc:   analyticsSvc,
		searchSvc:      searchSvc,
		exportSvc:      exportSvc,
	}
}

func (h *Handler) CreateMeeting(w http.ResponseWriter, r *http.Request) {
	var mtg meeting.Meeting
	if err := json.NewDecoder(r.Body).Decode(&mtg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.schedulingSvc.ScheduleMeeting(r.Context(), &mtg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(mtg)
}

func (h *Handler) ListMeetings(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	meetings, err := h.reportingSvc.ListMeetings(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(meetings)
}

func (h *Handler) GetMeeting(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	mtg, err := h.reportingSvc.GetMeeting(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(mtg)
}

func (h *Handler) StartMeeting(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.conductSvc.StartMeeting(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "IN_PROGRESS", "meeting_id": id})
}

func (h *Handler) RecordAttendance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rec attendanceDomain.Record
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.attendanceSvc.RecordAttendance(r.Context(), id, &rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) ApproveMinutes(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var m minutes.Minutes
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.closureSvc.ApproveMinutes(r.Context(), id, &m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(m)
}

func (h *Handler) CloseMeeting(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.closureSvc.CloseMeeting(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "CLOSED", "meeting_id": id})
}

func (h *Handler) SearchMeetings(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	meetings, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{"items": meetings, "total": total})
}

func (h *Handler) GetExecutiveReport(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	metrics, err := h.reportingSvc.GetExecutiveMetrics(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(metrics)
}

func (h *Handler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	criteria := &searchDomain.Criteria{PlantID: r.URL.Query().Get("plant_id")}
	data, err := h.exportSvc.ExportCSV(r.Context(), criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=meetings_report.csv")
	_, _ = w.Write(data)
}

func (h *Handler) ExportPDF(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	data, err := h.exportSvc.ExportPDF(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=meeting_minutes.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "meetings"})
}
