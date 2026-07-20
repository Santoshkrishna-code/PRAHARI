package http

import (
	"encoding/json"
	"net/http"
	"strings"

	prahariErrors "prahari/shared/errors"

	analyticsApp "prahari/services/energy/internal/application/analytics"
	consumptionApp "prahari/services/energy/internal/application/consumption"
	exportApp "prahari/services/energy/internal/application/export"
	forecastingApp "prahari/services/energy/internal/application/forecasting"
	optimizationApp "prahari/services/energy/internal/application/optimization"
	reportingApp "prahari/services/energy/internal/application/reporting"
	searchApp "prahari/services/energy/internal/application/search"
	"prahari/services/energy/internal/domain/energyforecast"
	"prahari/services/energy/internal/domain/energyprofile"
	"prahari/services/energy/internal/domain/energytarget"
	"prahari/services/energy/internal/domain/meterreading"
	"prahari/services/energy/internal/domain/optimization"
	"prahari/services/energy/internal/domain/search"
)

type Handler struct {
	sustSvc    *reportingApp.Service
	monSvc     *consumptionApp.Service
	carbonSvc  *forecastingApp.Service
	disclSvc   *optimizationApp.Service
	searchSvc  *searchApp.Service
	reportSvc  *analyticsApp.Service
	exportSvc  *exportApp.Service
}

func NewHandler(
	sustSvc *reportingApp.Service,
	monSvc *consumptionApp.Service,
	carbonSvc *forecastingApp.Service,
	disclSvc *optimizationApp.Service,
	searchSvc *searchApp.Service,
	reportSvc *analyticsApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		sustSvc:    sustSvc,
		monSvc:     monSvc,
		carbonSvc:  carbonSvc,
		disclSvc:   disclSvc,
		searchSvc:  searchSvc,
		reportSvc:  reportSvc,
		exportSvc:  exportSvc,
	}
}

func getID(path, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}

func (h *Handler) CreateEnergyProfile(w http.ResponseWriter, r *http.Request) {
	var req energyprofile.Profile
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.sustSvc.CreateProfile(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) GetEnergyProfile(w http.ResponseWriter, r *http.Request) {
	profileID := getID(r.URL.Path, "/energy/")
	if profileID == "" {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("profile ID is required", nil))
		return
	}

	resp := map[string]string{
		"id":            profileID,
		"plant_id":      "plant-3001",
		"facility_name": "Grid Transformer Facility",
		"status":        "ACTIVE",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *Handler) RecordMeterReading(w http.ResponseWriter, r *http.Request) {
	var req meterreading.Reading
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.monSvc.RecordMeterReading(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) ForecastEnergyDemand(w http.ResponseWriter, r *http.Request) {
	var req energyforecast.Forecast
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.carbonSvc.PredictDemand(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) RecommendOptimization(w http.ResponseWriter, r *http.Request) {
	var req optimization.Recommendation
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.disclSvc.RecommendOptimization(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) DefineTarget(w http.ResponseWriter, r *http.Request) {
	var req energytarget.Target
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) SearchProfiles(w http.ResponseWriter, r *http.Request) {
	var req search.Criteria
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid search criteria", err))
		return
	}

	results, err := h.searchSvc.Search(r.Context(), req)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(results)
}

func (h *Handler) GetDashboardReport(w http.ResponseWriter, r *http.Request) {
	results, err := h.reportSvc.GetDashboardReport(r.Context())
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(results)
}

func (h *Handler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=energy_profiles.csv")
	if err := h.exportSvc.WriteCSVReport(r.Context(), w); err != nil {
		prahariErrors.WriteHTTP(w, err)
	}
}

func (h *Handler) ExportPDF(w http.ResponseWriter, r *http.Request) {
	recordID := getID(r.URL.Path, "/export/pdf/")
	if recordID == "" {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("record ID is required", nil))
		return
	}

	pdfData, err := h.exportSvc.GetPDFReport(r.Context(), recordID)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment;filename=energy_audit_report.pdf")
	_, _ = w.Write(pdfData)
}
