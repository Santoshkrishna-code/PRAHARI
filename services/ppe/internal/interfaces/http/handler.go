package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/ppe/internal/application/analytics"
	catalogApp "prahari/services/ppe/internal/application/catalog"
	exportApp "prahari/services/ppe/internal/application/export"
	inspectionApp "prahari/services/ppe/internal/application/inspection"
	issuanceApp "prahari/services/ppe/internal/application/issuance"
	maintenanceApp "prahari/services/ppe/internal/application/maintenance"
	reportingApp "prahari/services/ppe/internal/application/reporting"
	searchApp "prahari/services/ppe/internal/application/search"
	"prahari/services/ppe/internal/domain/ppe"
	"prahari/services/ppe/internal/domain/ppeinspection"
	"prahari/services/ppe/internal/domain/ppeissue"
	"prahari/services/ppe/internal/domain/ppereturn"
	"prahari/services/ppe/internal/domain/ppemaintenance"
	searchDomain "prahari/services/ppe/internal/domain/search"
)

type Handler struct {
	catalogSvc     *catalogApp.Service
	issuanceSvc    *issuanceApp.Service
	inspectionSvc  *inspectionApp.Service
	maintenanceSvc *maintenanceApp.Service
	reportingSvc   *reportingApp.Service
	analyticsSvc   *analyticsApp.Service
	searchSvc      *searchApp.Service
	exportSvc      *exportApp.Service
}

func NewHandler(
	catalogSvc *catalogApp.Service,
	issuanceSvc *issuanceApp.Service,
	inspectionSvc *inspectionApp.Service,
	maintenanceSvc *maintenanceApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		catalogSvc:     catalogSvc,
		issuanceSvc:    issuanceSvc,
		inspectionSvc:  inspectionSvc,
		maintenanceSvc: maintenanceSvc,
		reportingSvc:   reportingSvc,
		analyticsSvc:   analyticsSvc,
		searchSvc:      searchSvc,
		exportSvc:      exportSvc,
	}
}

func (h *Handler) CreatePPE(w http.ResponseWriter, r *http.Request) {
	var p ppe.PPE
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.catalogSvc.CreateCatalogPPE(r.Context(), &p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(p)
}

func (h *Handler) ListPPEs(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	ppeList, err := h.reportingSvc.ListPPEs(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(ppeList)
}

func (h *Handler) GetPPE(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	p, err := h.reportingSvc.GetPPE(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(p)
}

func (h *Handler) IssuePPE(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rec ppeissue.Record
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rec.ItemID = id
	if err := h.issuanceSvc.IssuePPEItem(r.Context(), &rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) ReturnPPE(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rec ppereturn.Record
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rec.ItemID = id
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "RETURNED", "item_id": id})
}

func (h *Handler) InspectPPE(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rec ppeinspection.Record
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rec.ItemID = id
	if err := h.inspectionSvc.InspectPPEItem(r.Context(), &rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) MaintainPPE(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rec ppemaintenance.Record
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rec.ItemID = id
	if err := h.maintenanceSvc.MaintainPPEItem(r.Context(), &rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) SearchPPEs(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ppeList, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": ppeList,
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
	w.Header().Set("Content-Disposition", "attachment; filename=ppe_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=ppe_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "ppe"})
}
