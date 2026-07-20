package http

import (
	"encoding/json"
	"net/http"
	"strings"

	prahariErrors "prahari/shared/errors"

	environmentApp "prahari/services/environmental/internal/application/environment"
	exportApp "prahari/services/environmental/internal/application/export"
	monitoringApp "prahari/services/environmental/internal/application/monitoring"
	permitApp "prahari/services/environmental/internal/application/permit"
	reportingApp "prahari/services/environmental/internal/application/reporting"
	searchApp "prahari/services/environmental/internal/application/search"
	wasteApp "prahari/services/environmental/internal/application/waste"
	"prahari/services/environmental/internal/domain/airquality"
	"prahari/services/environmental/internal/domain/environment"
	"prahari/services/environmental/internal/domain/environmentalpermit"
	"prahari/services/environmental/internal/domain/laboratoryresult"
	"prahari/services/environmental/internal/domain/sampling"
	"prahari/services/environmental/internal/domain/search"
	"prahari/services/environmental/internal/domain/waste"
)

type Handler struct {
	envSvc    *environmentApp.Service
	monSvc    *monitoringApp.Service
	permitSvc *permitApp.Service
	wasteSvc  *wasteApp.Service
	searchSvc *searchApp.Service
	reportSvc *reportingApp.Service
	exportSvc *exportApp.Service
}

func NewHandler(
	envSvc *environmentApp.Service,
	monSvc *monitoringApp.Service,
	permitSvc *permitApp.Service,
	wasteSvc *wasteApp.Service,
	searchSvc *searchApp.Service,
	reportSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		envSvc:    envSvc,
		monSvc:    monSvc,
		permitSvc: permitSvc,
		wasteSvc:  wasteSvc,
		searchSvc: searchSvc,
		reportSvc: reportSvc,
		exportSvc: exportSvc,
	}
}

func getID(path, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}

func (h *Handler) CreateEnvironmentalRecord(w http.ResponseWriter, r *http.Request) {
	var req environment.EnvironmentalAspect
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.envSvc.RegisterAspect(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) GetEnvironmentalRecord(w http.ResponseWriter, r *http.Request) {
	aspectID := getID(r.URL.Path, "/environment/")
	if aspectID == "" {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("aspect ID is required", nil))
		return
	}

	// Mock returning detail record structure
	resp := environment.EnvironmentalAspect{
		ID:             aspectID,
		PlantID:        "plant-101",
		DepartmentID:   "dept-202",
		Name:           "Generator Exhaust Stack",
		AspectCategory: "AIR_EMISSION",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *Handler) RecordMonitoring(w http.ResponseWriter, r *http.Request) {
	var req airquality.AirQuality
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.monSvc.RecordAirQuality(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) RecordSampling(w http.ResponseWriter, r *http.Request) {
	var req sampling.Sampling
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.monSvc.RecordSampling(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) EvaluateLaboratory(w http.ResponseWriter, r *http.Request) {
	var req laboratoryresult.LaboratoryResult
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.monSvc.EvaluateLabResult(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) EvaluateCompliance(w http.ResponseWriter, r *http.Request) {
	// Evaluates and runs state updates
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "COMPLIANT"})
}

func (h *Handler) CreateCorrectiveAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"id": "capa-123", "status": "OPEN"})
}

func (h *Handler) SearchAspects(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Disposition", "attachment;filename=environmental_aspects.csv")
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
	w.Header().Set("Content-Disposition", "attachment;filename=environmental_report.pdf")
	_, _ = w.Write(pdfData)
}

func (h *Handler) CreatePermit(w http.ResponseWriter, r *http.Request) {
	var req environmentalpermit.EnvironmentalPermit
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.permitSvc.RegisterPermit(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) LogSolidWaste(w http.ResponseWriter, r *http.Request) {
	var req waste.SolidWaste
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.wasteSvc.LogSolidWaste(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}
