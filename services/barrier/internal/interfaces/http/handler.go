package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/barrier/internal/application/analytics"
	barrierApp "prahari/services/barrier/internal/application/barrier"
	exportApp "prahari/services/barrier/internal/application/export"
	impairmentApp "prahari/services/barrier/internal/application/impairment"
	integrityApp "prahari/services/barrier/internal/application/integrity"
	prooftestApp "prahari/services/barrier/internal/application/prooftest"
	reportingApp "prahari/services/barrier/internal/application/reporting"
	searchApp "prahari/services/barrier/internal/application/search"
	"prahari/services/barrier/internal/domain/barrier"
	"prahari/services/barrier/internal/domain/bypass"
	"prahari/services/barrier/internal/domain/impairment"
	"prahari/services/barrier/internal/domain/integrityassessment"
	"prahari/services/barrier/internal/domain/prooftest"
	searchDomain "prahari/services/barrier/internal/domain/search"
)

type Handler struct {
	barrierSvc    *barrierApp.Service
	integritySvc  *integrityApp.Service
	prooftestSvc  *prooftestApp.Service
	impairmentSvc *impairmentApp.Service
	reportingSvc  *reportingApp.Service
	analyticsSvc  *analyticsApp.Service
	searchSvc     *searchApp.Service
	exportSvc     *exportApp.Service
}

func NewHandler(
	barrierSvc *barrierApp.Service,
	integritySvc *integrityApp.Service,
	prooftestSvc *prooftestApp.Service,
	impairmentSvc *impairmentApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		barrierSvc:    barrierSvc,
		integritySvc:  integritySvc,
		prooftestSvc:  prooftestSvc,
		impairmentSvc: impairmentSvc,
		reportingSvc:  reportingSvc,
		analyticsSvc:  analyticsSvc,
		searchSvc:     searchSvc,
		exportSvc:     exportSvc,
	}
}

func (h *Handler) CreateBarrier(w http.ResponseWriter, r *http.Request) {
	var b barrier.Barrier
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.barrierSvc.CreateBarrier(r.Context(), &b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(b)
}

func (h *Handler) ListBarriers(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	barriers, err := h.reportingSvc.ListBarriers(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(barriers)
}

func (h *Handler) GetBarrier(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	b, err := h.reportingSvc.GetBarrier(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(b)
}

func (h *Handler) RecordProofTest(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var pt prooftest.Test
	if err := json.NewDecoder(r.Body).Decode(&pt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pt.BarrierID = id
	if err := h.prooftestSvc.RecordProofTest(r.Context(), &pt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(pt)
}

func (h *Handler) AssessIntegrity(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var ia integrityassessment.Assessment
	if err := json.NewDecoder(r.Body).Decode(&ia); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ia.BarrierID = id
	if err := h.integritySvc.AssessIntegrity(r.Context(), &ia); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(ia)
}

func (h *Handler) RegisterImpairment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var imp impairment.Record
	if err := json.NewDecoder(r.Body).Decode(&imp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	imp.BarrierID = id
	if err := h.impairmentSvc.RegisterImpairment(r.Context(), &imp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(imp)
}

func (h *Handler) RegisterBypass(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var bp bypass.Record
	if err := json.NewDecoder(r.Body).Decode(&bp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bp.BarrierID = id
	if err := h.impairmentSvc.RegisterBypass(r.Context(), &bp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(bp)
}

func (h *Handler) SearchBarriers(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	barriers, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": barriers,
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
	w.Header().Set("Content-Disposition", "attachment; filename=barrier_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=barrier_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "barrier"})
}
