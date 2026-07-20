package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/pha/internal/application/analytics"
	exportApp "prahari/services/pha/internal/application/export"
	hazopApp "prahari/services/pha/internal/application/hazop"
	lopaApp "prahari/services/pha/internal/application/lopa"
	recApp "prahari/services/pha/internal/application/recommendation"
	reportingApp "prahari/services/pha/internal/application/reporting"
	searchApp "prahari/services/pha/internal/application/search"
	studyApp "prahari/services/pha/internal/application/study"
	"prahari/services/pha/internal/domain/hazardscenario"
	"prahari/services/pha/internal/domain/lopa"
	"prahari/services/pha/internal/domain/phastudy"
	"prahari/services/pha/internal/domain/recommendation"
	searchDomain "prahari/services/pha/internal/domain/search"
)

type Handler struct {
	studySvc     *studyApp.Service
	hazopSvc     *hazopApp.Service
	lopaSvc      *lopaApp.Service
	recSvc       *recApp.Service
	reportingSvc *reportingApp.Service
	analyticsSvc *analyticsApp.Service
	searchSvc    *searchApp.Service
	exportSvc    *exportApp.Service
}

func NewHandler(
	studySvc *studyApp.Service,
	hazopSvc *hazopApp.Service,
	lopaSvc *lopaApp.Service,
	recSvc *recApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		studySvc:     studySvc,
		hazopSvc:     hazopSvc,
		lopaSvc:      lopaSvc,
		recSvc:       recSvc,
		reportingSvc: reportingSvc,
		analyticsSvc: analyticsSvc,
		searchSvc:    searchSvc,
		exportSvc:    exportSvc,
	}
}

func (h *Handler) CreatePHA(w http.ResponseWriter, r *http.Request) {
	var st phastudy.Study
	if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.studySvc.CreateStudy(r.Context(), &st); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(st)
}

func (h *Handler) ListPHAs(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	studies, err := h.reportingSvc.ListStudies(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(studies)
}

func (h *Handler) GetPHA(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	st, err := h.reportingSvc.GetStudy(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(st)
}

func (h *Handler) ExecuteHAZOP(w http.ResponseWriter, r *http.Request) {
	var sc hazardscenario.Scenario
	if err := json.NewDecoder(r.Body).Decode(&sc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.hazopSvc.RecordScenario(r.Context(), &sc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(sc)
}

func (h *Handler) ExecuteLOPA(w http.ResponseWriter, r *http.Request) {
	var analysis lopa.Analysis
	if err := json.NewDecoder(r.Body).Decode(&analysis); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := h.lopaSvc.ExecuteLOPA(r.Context(), &analysis)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(result)
}

func (h *Handler) CreateRecommendation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rec recommendation.Recommendation
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rec.StudyID = id
	if err := h.recSvc.CreateRecommendation(r.Context(), &rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) ApprovePHA(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.studySvc.ApproveStudy(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "APPROVED", "study_id": id})
}

func (h *Handler) SearchPHAs(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	studies, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": studies,
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
	w.Header().Set("Content-Disposition", "attachment; filename=pha_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=pha_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "pha"})
}
