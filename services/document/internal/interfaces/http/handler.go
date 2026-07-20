package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/document/internal/application/analytics"
	approvalApp "prahari/services/document/internal/application/approval"
	creationApp "prahari/services/document/internal/application/creation"
	distributionApp "prahari/services/document/internal/application/distribution"
	exportApp "prahari/services/document/internal/application/export"
	lifecycleApp "prahari/services/document/internal/application/lifecycle"
	reportingApp "prahari/services/document/internal/application/reporting"
	searchApp "prahari/services/document/internal/application/search"
	versioningApp "prahari/services/document/internal/application/versioning"
	"prahari/services/document/internal/domain/document"

	searchDomain "prahari/services/document/internal/domain/search"
)

type Handler struct {
	creationSvc     *creationApp.Service
	lifecycleSvc    *lifecycleApp.Service
	versioningSvc   *versioningApp.Service
	approvalSvc     *approvalApp.Service
	distributionSvc *distributionApp.Service
	reportingSvc    *reportingApp.Service
	analyticsSvc    *analyticsApp.Service
	searchSvc       *searchApp.Service
	exportSvc       *exportApp.Service
}

func NewHandler(
	creationSvc *creationApp.Service,
	lifecycleSvc *lifecycleApp.Service,
	versioningSvc *versioningApp.Service,
	approvalSvc *approvalApp.Service,
	distributionSvc *distributionApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		creationSvc:     creationSvc,
		lifecycleSvc:    lifecycleSvc,
		versioningSvc:   versioningSvc,
		approvalSvc:     approvalSvc,
		distributionSvc: distributionSvc,
		reportingSvc:    reportingSvc,
		analyticsSvc:    analyticsSvc,
		searchSvc:       searchSvc,
		exportSvc:       exportSvc,
	}
}

func (h *Handler) CreateDocument(w http.ResponseWriter, r *http.Request) {
	var doc document.Document
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.creationSvc.CreateDocument(r.Context(), &doc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(doc)
}

func (h *Handler) ListDocuments(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	docs, err := h.reportingSvc.ListDocuments(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(docs)
}

func (h *Handler) GetDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	doc, err := h.reportingSvc.GetDocument(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(doc)
}

func (h *Handler) ApproveDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req struct {
		ApproverID string `json:"approver_id"`
		Comments   string `json:"comments"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.approvalSvc.ApproveDocument(r.Context(), id, req.ApproverID, req.Comments); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "APPROVED", "document_id": id})
}

func (h *Handler) PublishDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.lifecycleSvc.PublishDocument(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "PUBLISHED", "document_id": id})
}

func (h *Handler) CheckoutDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req struct {
		UserID string `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.versioningSvc.Checkout(r.Context(), id, req.UserID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "CHECKED_OUT", "document_id": id})
}

func (h *Handler) CheckinDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req struct {
		UserID        string `json:"user_id"`
		FileURL       string `json:"file_url"`
		FileHash      string `json:"file_hash"`
		ChangeSummary string `json:"change_summary"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ver, err := h.versioningSvc.Checkin(r.Context(), id, req.UserID, req.FileURL, req.FileHash, req.ChangeSummary)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(ver)
}

func (h *Handler) SearchDocuments(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	docs, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": docs,
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
	w.Header().Set("Content-Disposition", "attachment; filename=document_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=document_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "document"})
}
