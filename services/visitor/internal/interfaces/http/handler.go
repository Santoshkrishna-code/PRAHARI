package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/visitor/internal/application/analytics"
	approvalApp "prahari/services/visitor/internal/application/approval"
	checkinApp "prahari/services/visitor/internal/application/checkin"
	checkoutApp "prahari/services/visitor/internal/application/checkout"
	exportApp "prahari/services/visitor/internal/application/export"
	musterApp "prahari/services/visitor/internal/application/muster"
	registrationApp "prahari/services/visitor/internal/application/registration"
	reportingApp "prahari/services/visitor/internal/application/reporting"
	searchApp "prahari/services/visitor/internal/application/search"
	"prahari/services/visitor/internal/domain/emergencymuster"
	searchDomain "prahari/services/visitor/internal/domain/search"
	"prahari/services/visitor/internal/domain/visitor"
)

type Handler struct {
	registrationSvc *registrationApp.Service
	approvalSvc     *approvalApp.Service
	checkinSvc      *checkinApp.Service
	checkoutSvc     *checkoutApp.Service
	musterSvc       *musterApp.Service
	reportingSvc    *reportingApp.Service
	analyticsSvc    *analyticsApp.Service
	searchSvc       *searchApp.Service
	exportSvc       *exportApp.Service
}

func NewHandler(
	registrationSvc *registrationApp.Service,
	approvalSvc *approvalApp.Service,
	checkinSvc *checkinApp.Service,
	checkoutSvc *checkoutApp.Service,
	musterSvc *musterApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		registrationSvc: registrationSvc,
		approvalSvc:     approvalSvc,
		checkinSvc:      checkinSvc,
		checkoutSvc:     checkoutSvc,
		musterSvc:       musterSvc,
		reportingSvc:    reportingSvc,
		analyticsSvc:    analyticsSvc,
		searchSvc:       searchSvc,
		exportSvc:       exportSvc,
	}
}

func (h *Handler) CreateVisitor(w http.ResponseWriter, r *http.Request) {
	var vis visitor.Visitor
	if err := json.NewDecoder(r.Body).Decode(&vis); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.registrationSvc.RegisterVisitor(r.Context(), &vis); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(vis)
}

func (h *Handler) ListVisitors(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	visits, err := h.reportingSvc.ListVisits(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(visits)
}

func (h *Handler) GetVisitor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	vis, err := h.reportingSvc.GetVisitor(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(vis)
}

func (h *Handler) ApproveVisitor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.approvalSvc.ApproveVisit(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "HOST_APPROVED", "visit_id": id})
}

func (h *Handler) CheckInVisitor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req struct {
		SecurityCheckpoint string `json:"security_checkpoint"`
		OperatorID         string `json:"operator_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.checkinSvc.CheckIn(r.Context(), id, req.SecurityCheckpoint, req.OperatorID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "CHECKED_IN", "visit_id": id})
}

func (h *Handler) CheckOutVisitor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req struct {
		OperatorID string `json:"operator_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.checkoutSvc.CheckOut(r.Context(), id, req.OperatorID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "CHECKED_OUT", "visit_id": id})
}

func (h *Handler) EvacuateVisitor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rec emergencymuster.Record
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rec.VisitorID = id
	if err := h.musterSvc.AccountForVisitor(r.Context(), &rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) SearchVisitors(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	visits, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": visits,
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
	w.Header().Set("Content-Disposition", "attachment; filename=visitor_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=visitor_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "visitor"})
}
