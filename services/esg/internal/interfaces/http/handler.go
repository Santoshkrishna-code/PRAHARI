package http

import (
	"encoding/json"
	"net/http"
	"strings"

	prahariErrors "prahari/shared/errors"

	analyticsApp "prahari/services/esg/internal/application/analytics"
	carbonApp "prahari/services/esg/internal/application/carbon"
	disclosureApp "prahari/services/esg/internal/application/disclosure"
	exportApp "prahari/services/esg/internal/application/export"
	reportingApp "prahari/services/esg/internal/application/reporting"
	searchApp "prahari/services/esg/internal/application/search"
	sustainabilityApp "prahari/services/esg/internal/application/sustainability"
	"prahari/services/esg/internal/domain/carboninventory"
	"prahari/services/esg/internal/domain/disclosure"
	"prahari/services/esg/internal/domain/esgobjective"
	"prahari/services/esg/internal/domain/search"
	"prahari/services/esg/internal/domain/sustainabilityreport"
)

type Handler struct {
	sustSvc   *sustainabilityApp.Service
	monSvc    *reportingApp.Service
	carbonSvc *carbonApp.Service
	disclSvc  *disclosureApp.Service
	searchSvc *searchApp.Service
	reportSvc *analyticsApp.Service
	exportSvc *exportApp.Service
}

func NewHandler(
	sustSvc *sustainabilityApp.Service,
	monSvc *reportingApp.Service,
	carbonSvc *carbonApp.Service,
	disclSvc *disclosureApp.Service,
	searchSvc *searchApp.Service,
	reportSvc *analyticsApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		sustSvc:   sustSvc,
		monSvc:    monSvc,
		carbonSvc: carbonSvc,
		disclSvc:  disclSvc,
		searchSvc: searchSvc,
		reportSvc: reportSvc,
		exportSvc: exportSvc,
	}
}

func getID(path, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}

func (h *Handler) CreateESGProfile(w http.ResponseWriter, r *http.Request) {
	// Simple mock mapping ESG organizational scope profile values
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "CREATED"})
}

func (h *Handler) GetESGProfile(w http.ResponseWriter, r *http.Request) {
	profileID := getID(r.URL.Path, "/esg/")
	if profileID == "" {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("profile ID is required", nil))
		return
	}

	resp := map[string]string{
		"id":               profileID,
		"business_unit_id": "bu-4001",
		"status":           "ACTIVE",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *Handler) CalculateCarbon(w http.ResponseWriter, r *http.Request) {
	var req carboninventory.Inventory
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.carbonSvc.CalculateCarbon(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) PublishDisclosure(w http.ResponseWriter, r *http.Request) {
	var req disclosure.Disclosure
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.disclSvc.PublishDisclosure(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) CreateObjective(w http.ResponseWriter, r *http.Request) {
	var req esgobjective.Objective
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.sustSvc.CreateObjective(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) CreateReport(w http.ResponseWriter, r *http.Request) {
	var req sustainabilityreport.Report
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.monSvc.GenerateReport(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) SearchObjectives(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Disposition", "attachment;filename=sustainability_objectives.csv")
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
	w.Header().Set("Content-Disposition", "attachment;filename=sustainability_report.pdf")
	_, _ = w.Write(pdfData)
}
