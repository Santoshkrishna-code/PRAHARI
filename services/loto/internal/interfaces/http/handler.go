package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/loto/internal/application/analytics"
	executionApp "prahari/services/loto/internal/application/execution"
	exportApp "prahari/services/loto/internal/application/export"
	planningApp "prahari/services/loto/internal/application/planning"
	reportingApp "prahari/services/loto/internal/application/reporting"
	restorationApp "prahari/services/loto/internal/application/restoration"
	searchApp "prahari/services/loto/internal/application/search"
	verificationApp "prahari/services/loto/internal/application/verification"
	"prahari/services/loto/internal/domain/isolationcertificate"
	"prahari/services/loto/internal/domain/isolationplan"
	"prahari/services/loto/internal/domain/restoration"
	searchDomain "prahari/services/loto/internal/domain/search"
	"prahari/services/loto/internal/domain/verification"
)

type Handler struct {
	planningSvc     *planningApp.Service
	executionSvc    *executionApp.Service
	verificationSvc *verificationApp.Service
	restorationSvc  *restorationApp.Service
	reportingSvc    *reportingApp.Service
	analyticsSvc    *analyticsApp.Service
	searchSvc       *searchApp.Service
	exportSvc       *exportApp.Service
}

func NewHandler(
	planningSvc *planningApp.Service,
	executionSvc *executionApp.Service,
	verificationSvc *verificationApp.Service,
	restorationSvc *restorationApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		planningSvc:     planningSvc,
		executionSvc:    executionSvc,
		verificationSvc: verificationSvc,
		restorationSvc:  restorationSvc,
		reportingSvc:    reportingSvc,
		analyticsSvc:    analyticsSvc,
		searchSvc:       searchSvc,
		exportSvc:       exportSvc,
	}
}

func (h *Handler) CreateLOTO(w http.ResponseWriter, r *http.Request) {
	var plan isolationplan.Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.planningSvc.CreateIsolationPlan(r.Context(), &plan); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(plan)
}

func (h *Handler) ListLOTOs(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	certs, err := h.reportingSvc.ListCertificates(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(certs)
}

func (h *Handler) GetLOTO(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, err := h.reportingSvc.GetCertificate(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(c)
}

func (h *Handler) ApproveIsolation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var cert isolationcertificate.Certificate
	if err := json.NewDecoder(r.Body).Decode(&cert); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cert.PlanID = id
	if err := h.executionSvc.ApproveIsolation(r.Context(), &cert); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(cert)
}

func (h *Handler) ApplyLocks(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.executionSvc.ApplyLocksAndTags(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "LOCKS_APPLIED", "certificate_id": id})
}

func (h *Handler) VerifyZeroEnergy(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var v verification.ZeroEnergy
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.verificationSvc.VerifyZeroEnergy(r.Context(), id, &v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handler) RestoreSystem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rec restoration.Record
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.restorationSvc.RestoreSystem(r.Context(), id, &rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) SearchLOTOs(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	certs, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": certs,
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
	w.Header().Set("Content-Disposition", "attachment; filename=loto_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=loto_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "loto"})
}
