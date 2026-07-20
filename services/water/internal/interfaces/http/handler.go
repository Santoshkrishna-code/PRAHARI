package http

import (
	"encoding/json"
	"net/http"
	"time"

	analyticsApp "prahari/services/water/internal/application/analytics"
	consumptionApp "prahari/services/water/internal/application/consumption"
	exportApp "prahari/services/water/internal/application/export"
	forecastingApp "prahari/services/water/internal/application/forecasting"
	optimizationApp "prahari/services/water/internal/application/optimization"
	recyclingApp "prahari/services/water/internal/application/recycling"
	reportingApp "prahari/services/water/internal/application/reporting"
	searchApp "prahari/services/water/internal/application/search"
	"prahari/services/water/internal/domain/meterreading"
	"prahari/services/water/internal/domain/optimization"
	"prahari/services/water/internal/domain/recycling"
	searchDomain "prahari/services/water/internal/domain/search"
	"prahari/services/water/internal/domain/waterprofile"
)

type Handler struct {
	reportingSvc    *reportingApp.Service
	consumptionSvc  *consumptionApp.Service
	recyclingSvc    *recyclingApp.Service
	forecastingSvc  *forecastingApp.Service
	optimizationSvc *optimizationApp.Service
	searchSvc       *searchApp.Service
	analyticsSvc    *analyticsApp.Service
	exportSvc       *exportApp.Service
}

func NewHandler(
	reportingSvc *reportingApp.Service,
	consumptionSvc *consumptionApp.Service,
	recyclingSvc *recyclingApp.Service,
	forecastingSvc *forecastingApp.Service,
	optimizationSvc *optimizationApp.Service,
	searchSvc *searchApp.Service,
	analyticsSvc *analyticsApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		reportingSvc:    reportingSvc,
		consumptionSvc:  consumptionSvc,
		recyclingSvc:    recyclingSvc,
		forecastingSvc:  forecastingSvc,
		optimizationSvc: optimizationSvc,
		searchSvc:       searchSvc,
		analyticsSvc:    analyticsSvc,
		exportSvc:       exportSvc,
	}
}

func (h *Handler) CreateWaterProfile(w http.ResponseWriter, r *http.Request) {
	var profile waterprofile.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()
	if err := h.reportingSvc.CreateProfile(r.Context(), &profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(profile)
}

func (h *Handler) ListWaterProfiles(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	profiles, err := h.reportingSvc.ListProfiles(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(profiles)
}

func (h *Handler) GetWaterProfile(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	profile, err := h.reportingSvc.GetProfile(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(profile)
}

func (h *Handler) RecordMeterReading(w http.ResponseWriter, r *http.Request) {
	var reading meterreading.Reading
	if err := json.NewDecoder(r.Body).Decode(&reading); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.consumptionSvc.RecordReading(r.Context(), &reading); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(reading)
}

func (h *Handler) RegisterRecycling(w http.ResponseWriter, r *http.Request) {
	var prog recycling.Program
	if err := json.NewDecoder(r.Body).Decode(&prog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.recyclingSvc.RegisterRecyclingProgram(r.Context(), &prog); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(prog)
}

func (h *Handler) ForecastWater(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PlantID    string  `json:"plant_id"`
		Period     string  `json:"period"`
		BaselineKL float64 `json:"baseline_kl"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fc, err := h.forecastingSvc.GenerateForecast(r.Context(), req.PlantID, req.Period, req.BaselineKL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(fc)
}

func (h *Handler) CreateOptimization(w http.ResponseWriter, r *http.Request) {
	var rec optimization.Recommendation
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.optimizationSvc.CreateRecommendation(r.Context(), &rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) SearchProfiles(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	profiles, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": profiles,
		"total": total,
	})
}

func (h *Handler) GetExecutiveReport(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	metrics, err := h.analyticsSvc.GetExecutiveMetrics(r.Context(), plantID)
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
	w.Header().Set("Content-Disposition", "attachment; filename=water_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=water_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "water"})
}
