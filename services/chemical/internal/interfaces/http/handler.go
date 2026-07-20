package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/chemical/internal/application/analytics"
	exportApp "prahari/services/chemical/internal/application/export"
	inventoryApp "prahari/services/chemical/internal/application/inventory"
	reportingApp "prahari/services/chemical/internal/application/reporting"
	sdsApp "prahari/services/chemical/internal/application/sds"
	storageApp "prahari/services/chemical/internal/application/storage"
	searchApp "prahari/services/chemical/internal/application/search"
	"prahari/services/chemical/internal/domain/container"
	sdsDomain "prahari/services/chemical/internal/domain/sds"
	searchDomain "prahari/services/chemical/internal/domain/search"
)

type Handler struct {
	inventorySvc *inventoryApp.Service
	storageSvc   *storageApp.Service
	sdsSvc       *sdsApp.Service
	reportingSvc *reportingApp.Service
	analyticsSvc *analyticsApp.Service
	searchSvc    *searchApp.Service
	exportSvc    *exportApp.Service
}

func NewHandler(
	inventorySvc *inventoryApp.Service,
	storageSvc *storageApp.Service,
	sdsSvc *sdsApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		inventorySvc: inventorySvc,
		storageSvc:   storageSvc,
		sdsSvc:       sdsSvc,
		reportingSvc: reportingSvc,
		analyticsSvc: analyticsSvc,
		searchSvc:    searchSvc,
		exportSvc:    exportSvc,
	}
}

func (h *Handler) ReceiveChemical(w http.ResponseWriter, r *http.Request) {
	var con container.Container
	if err := json.NewDecoder(r.Body).Decode(&con); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.inventorySvc.ReceiveContainer(r.Context(), &con); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(con)
}

func (h *Handler) IssueChemical(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req struct {
		IssuedTo string `json:"issued_to"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.inventorySvc.IssueContainer(r.Context(), id, req.IssuedTo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ISSUED", "container_id": id})
}

func (h *Handler) ReturnChemical(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.inventorySvc.ReturnContainer(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "STORED", "container_id": id})
}

func (h *Handler) TransferChemical(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req struct {
		TargetAreaID string `json:"target_area_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.inventorySvc.TransferContainer(r.Context(), id, req.TargetAreaID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "STORED", "container_id": id})
}

func (h *Handler) GetSDS(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // chemical_id
	sd, err := h.sdsSvc.GetSDS(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(sd)
}

func (h *Handler) RegisterSDS(w http.ResponseWriter, r *http.Request) {
	var sd sdsDomain.SDS
	if err := json.NewDecoder(r.Body).Decode(&sd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.sdsSvc.RegisterSDS(r.Context(), &sd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(sd)
}

func (h *Handler) GetChemical(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, err := h.reportingSvc.GetChemical(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(c)
}

func (h *Handler) ListChemicals(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	chems, err := h.reportingSvc.ListChemicals(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(chems)
}

func (h *Handler) SearchChemicals(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chems, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{"items": chems, "total": total})
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
	w.Header().Set("Content-Disposition", "attachment; filename=chemical_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=sds_summary.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "chemical"})
}
