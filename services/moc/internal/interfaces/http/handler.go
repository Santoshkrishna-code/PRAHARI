package http

import (
	"encoding/json"
	"net/http"

	exportApp "prahari/services/moc/internal/application/export"
	implApp "prahari/services/moc/internal/application/implementation"
	reportingApp "prahari/services/moc/internal/application/reporting"
	reqApp "prahari/services/moc/internal/application/request"
	reviewApp "prahari/services/moc/internal/application/review"
	searchApp "prahari/services/moc/internal/application/search"
	verifApp "prahari/services/moc/internal/application/verification"
	"prahari/services/moc/internal/domain/approval"
	"prahari/services/moc/internal/domain/changerequest"
	"prahari/services/moc/internal/domain/impactassessment"
	"prahari/services/moc/internal/domain/implementation"
	"prahari/services/moc/internal/domain/riskreview"
	searchDomain "prahari/services/moc/internal/domain/search"
	"prahari/services/moc/internal/domain/technicalreview"
	"prahari/services/moc/internal/domain/verification"
)

type Handler struct {
	reqSvc       *reqApp.Service
	reviewSvc    *reviewApp.Service
	implSvc      *implApp.Service
	verifSvc     *verifApp.Service
	reportingSvc *reportingApp.Service
	searchSvc    *searchApp.Service
	exportSvc    *exportApp.Service
}

func NewHandler(
	reqSvc *reqApp.Service,
	reviewSvc *reviewApp.Service,
	implSvc *implApp.Service,
	verifSvc *verifApp.Service,
	reportingSvc *reportingApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		reqSvc:       reqSvc,
		reviewSvc:    reviewSvc,
		implSvc:      implSvc,
		verifSvc:     verifSvc,
		reportingSvc: reportingSvc,
		searchSvc:    searchSvc,
		exportSvc:    exportSvc,
	}
}

func (h *Handler) CreateMOC(w http.ResponseWriter, r *http.Request) {
	var req changerequest.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.reqSvc.CreateRequest(r.Context(), &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) ListMOCs(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	reqs, err := h.reportingSvc.ListRequests(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(reqs)
}

func (h *Handler) GetMOC(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	req, err := h.reportingSvc.GetRequest(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) SubmitImpactAssessment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var ia impactassessment.Assessment
	if err := json.NewDecoder(r.Body).Decode(&ia); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ia.ChangeRequestID = id
	if err := h.reviewSvc.SubmitImpactAssessment(r.Context(), &ia); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(ia)
}

func (h *Handler) SubmitTechnicalReview(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var tr technicalreview.Review
	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tr.ChangeRequestID = id
	if err := h.reviewSvc.SubmitTechnicalReview(r.Context(), &tr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(tr)
}

func (h *Handler) SubmitRiskReview(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rr riskreview.Review
	if err := json.NewDecoder(r.Body).Decode(&rr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rr.ChangeRequestID = id
	if err := h.reviewSvc.SubmitRiskReview(r.Context(), &rr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(rr)
}

func (h *Handler) SubmitApproval(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var app approval.Record
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	app.ChangeRequestID = id
	if err := h.reviewSvc.ApproveChange(r.Context(), &app); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(app)
}

func (h *Handler) StartImplementation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var plan implementation.Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	plan.ChangeRequestID = id
	if err := h.implSvc.StartImplementation(r.Context(), &plan); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(plan)
}

func (h *Handler) SubmitVerification(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var v verification.Record
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	v.ChangeRequestID = id
	if err := h.verifSvc.VerifyChange(r.Context(), &v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handler) SearchMOCs(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reqs, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": reqs,
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
	w.Header().Set("Content-Disposition", "attachment; filename=moc_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=moc_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "moc"})
}
