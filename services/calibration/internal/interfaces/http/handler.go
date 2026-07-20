package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/calibration/internal/application/analytics"
	certificateApp "prahari/services/calibration/internal/application/certificate"
	executionApp "prahari/services/calibration/internal/application/execution"
	exportApp "prahari/services/calibration/internal/application/export"
	reportingApp "prahari/services/calibration/internal/application/reporting"
	schedulingApp "prahari/services/calibration/internal/application/scheduling"
	searchApp "prahari/services/calibration/internal/application/search"
	"prahari/services/calibration/internal/domain/calibration"
	"prahari/services/calibration/internal/domain/calibrationschedule"
	"prahari/services/calibration/internal/domain/measurement"
	searchDomain "prahari/services/calibration/internal/domain/search"
)

type Handler struct {
	schedulingSvc  *schedulingApp.Service
	executionSvc   *executionApp.Service
	certificateSvc *certificateApp.Service
	reportingSvc   *reportingApp.Service
	analyticsSvc   *analyticsApp.Service
	searchSvc      *searchApp.Service
	exportSvc      *exportApp.Service
}

func NewHandler(
	schedulingSvc *schedulingApp.Service,
	executionSvc *executionApp.Service,
	certificateSvc *certificateApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		schedulingSvc:  schedulingSvc,
		executionSvc:   executionSvc,
		certificateSvc: certificateSvc,
		reportingSvc:   reportingSvc,
		analyticsSvc:   analyticsSvc,
		searchSvc:      searchSvc,
		exportSvc:      exportSvc,
	}
}

func (h *Handler) CreateCalibration(w http.ResponseWriter, r *http.Request) {
	var rec calibration.Record
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.executionSvc.StartCalibration(r.Context(), &rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) ListCalibrations(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	cals, err := h.reportingSvc.ListCalibrations(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(cals)
}

func (h *Handler) GetCalibration(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, err := h.reportingSvc.GetCalibration(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(c)
}

func (h *Handler) ScheduleCalibration(w http.ResponseWriter, r *http.Request) {
	var sched calibrationschedule.Schedule
	if err := json.NewDecoder(r.Body).Decode(&sched); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.schedulingSvc.ScheduleCalibrationTask(r.Context(), &sched); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(sched)
}

func (h *Handler) ExecuteCalibration(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rec calibration.Record
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rec.InstrumentID = id
	if err := h.executionSvc.StartCalibration(r.Context(), &rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) RecordMeasurements(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var m measurement.Result
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.executionSvc.RecordMeasurements(r.Context(), id, &m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(m)
}

func (h *Handler) ApproveCalibration(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req struct {
		SupervisorID string `json:"supervisor_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.executionSvc.ApproveCalibration(r.Context(), id, req.SupervisorID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "APPROVED", "calibration_id": id})
}

func (h *Handler) SearchCalibrations(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cals, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": cals,
		"total": total,
	})
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
	w.Header().Set("Content-Disposition", "attachment; filename=calibration_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=calibration_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "calibration"})
}
