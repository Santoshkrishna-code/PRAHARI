package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/bcm/internal/application/analytics"
	biaApp "prahari/services/bcm/internal/application/bia"
	exercisesApp "prahari/services/bcm/internal/application/exercises"
	exportApp "prahari/services/bcm/internal/application/export"
	planningApp "prahari/services/bcm/internal/application/planning"
	recoveryApp "prahari/services/bcm/internal/application/recovery"
	reportingApp "prahari/services/bcm/internal/application/reporting"
	resilienceApp "prahari/services/bcm/internal/application/resilience"
	searchApp "prahari/services/bcm/internal/application/search"
	"prahari/services/bcm/internal/domain/businessimpactanalysis"
	"prahari/services/bcm/internal/domain/continuityexercise"
	"prahari/services/bcm/internal/domain/continuityplan"
	searchDomain "prahari/services/bcm/internal/domain/search"
)

type Handler struct {
	biaSvc        *biaApp.Service
	planningSvc   *planningApp.Service
	recoverySvc   *recoveryApp.Service
	exercisesSvc  *exercisesApp.Service
	resilienceSvc *resilienceApp.Service
	reportingSvc  *reportingApp.Service
	analyticsSvc  *analyticsApp.Service
	searchSvc     *searchApp.Service
	exportSvc     *exportApp.Service
}

func NewHandler(
	biaSvc *biaApp.Service,
	planningSvc *planningApp.Service,
	recoverySvc *recoveryApp.Service,
	exercisesSvc *exercisesApp.Service,
	resilienceSvc *resilienceApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		biaSvc:        biaSvc,
		planningSvc:   planningSvc,
		recoverySvc:   recoverySvc,
		exercisesSvc:  exercisesSvc,
		resilienceSvc: resilienceSvc,
		reportingSvc:  reportingSvc,
		analyticsSvc:  analyticsSvc,
		searchSvc:     searchSvc,
		exportSvc:     exportSvc,
	}
}

func (h *Handler) CreateBCM(w http.ResponseWriter, r *http.Request) {
	var plan continuityplan.Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.planningSvc.CreateContinuityPlan(r.Context(), &plan); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(plan)
}

func (h *Handler) ListBCMs(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	plans, err := h.reportingSvc.ListPlans(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(plans)
}

func (h *Handler) GetBCM(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	plan, err := h.reportingSvc.GetPlan(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(plan)
}

func (h *Handler) ExecuteBIA(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var bia businessimpactanalysis.Analysis
	if err := json.NewDecoder(r.Body).Decode(&bia); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bia.PlanID = id
	result, err := h.biaSvc.ExecuteBIA(r.Context(), &bia)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(result)
}

func (h *Handler) ScheduleExercise(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var ex continuityexercise.Exercise
	if err := json.NewDecoder(r.Body).Decode(&ex); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ex.PlanID = id
	if err := h.exercisesSvc.ScheduleExercise(r.Context(), &ex); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(ex)
}

func (h *Handler) ActivatePlan(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.planningSvc.ActivatePlan(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ACTIVATION", "plan_id": id})
}

func (h *Handler) CompleteRecovery(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.recoverySvc.CompleteRecovery(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "RECOVERY", "plan_id": id})
}

func (h *Handler) SearchBCMs(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	plans, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": plans,
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
	w.Header().Set("Content-Disposition", "attachment; filename=bcm_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=bcm_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "bcm"})
}
