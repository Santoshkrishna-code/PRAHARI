package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/shift/internal/application/analytics"
	exportApp "prahari/services/shift/internal/application/export"
	handoverApp "prahari/services/shift/internal/application/handover"
	logbookApp "prahari/services/shift/internal/application/logbook"
	schedulingApp "prahari/services/shift/internal/application/scheduling"
	reportingApp "prahari/services/shift/internal/application/reporting"
	searchApp "prahari/services/shift/internal/application/search"
	"prahari/services/shift/internal/domain/handover"
	"prahari/services/shift/internal/domain/operatorjournal"
	"prahari/services/shift/internal/domain/shift"
	"prahari/services/shift/internal/domain/shiftlog"
	searchDomain "prahari/services/shift/internal/domain/search"
)

type Handler struct {
	schedulingSvc *schedulingApp.Service
	handoverSvc   *handoverApp.Service
	logbookSvc    *logbookApp.Service
	reportingSvc  *reportingApp.Service
	analyticsSvc  *analyticsApp.Service
	searchSvc     *searchApp.Service
	exportSvc     *exportApp.Service
}

func NewHandler(
	schedulingSvc *schedulingApp.Service,
	handoverSvc *handoverApp.Service,
	logbookSvc *logbookApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		schedulingSvc: schedulingSvc,
		handoverSvc:   handoverSvc,
		logbookSvc:    logbookSvc,
		reportingSvc:  reportingSvc,
		analyticsSvc:  analyticsSvc,
		searchSvc:     searchSvc,
		exportSvc:     exportSvc,
	}
}

func (h *Handler) CreateShift(w http.ResponseWriter, r *http.Request) {
	var sh shift.Shift
	if err := json.NewDecoder(r.Body).Decode(&sh); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.schedulingSvc.CreateShift(r.Context(), &sh); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(sh)
}

func (h *Handler) ListShifts(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	shifts, err := h.reportingSvc.ListShifts(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(shifts)
}

func (h *Handler) GetShift(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	sh, err := h.reportingSvc.GetShift(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(sh)
}

func (h *Handler) StartShift(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.schedulingSvc.StartShift(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "SHIFT_STARTED", "shift_id": id})
}

func (h *Handler) InitiateHandover(w http.ResponseWriter, r *http.Request) {
	var ho handover.Handover
	if err := json.NewDecoder(r.Body).Decode(&ho); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.handoverSvc.InitiateHandover(r.Context(), &ho); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(ho)
}

func (h *Handler) AcceptHandover(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.handoverSvc.AcceptHandover(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ACCEPTED", "handover_id": id})
}

func (h *Handler) CloseShift(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "SHIFT_CLOSED", "shift_id": id})
}

func (h *Handler) LogActivity(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var log shiftlog.Log
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.ShiftID = id
	if err := h.logbookSvc.LogActivity(r.Context(), &log); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(log)
}

func (h *Handler) WriteJournal(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var jr operatorjournal.Journal
	if err := json.NewDecoder(r.Body).Decode(&jr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jr.ShiftID = id
	if err := h.logbookSvc.WriteJournal(r.Context(), &jr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(jr)
}

func (h *Handler) SearchShifts(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shifts, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": shifts,
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
	w.Header().Set("Content-Disposition", "attachment; filename=shift_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=shift_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "shift"})
}
