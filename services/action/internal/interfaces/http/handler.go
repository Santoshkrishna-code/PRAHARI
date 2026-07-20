package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/action/internal/application/analytics"
	executionApp "prahari/services/action/internal/application/execution"
	exportApp "prahari/services/action/internal/application/export"
	planningApp "prahari/services/action/internal/application/planning"
	reportingApp "prahari/services/action/internal/application/reporting"
	searchApp "prahari/services/action/internal/application/search"
	verificationApp "prahari/services/action/internal/application/verification"
	"prahari/services/action/internal/domain/action"
	"prahari/services/action/internal/domain/effectivenessreview"
	searchDomain "prahari/services/action/internal/domain/search"
)

type Handler struct {
	planningSvc     *planningApp.Service
	executionSvc    *executionApp.Service
	verificationSvc *verificationApp.Service
	reportingSvc    *reportingApp.Service
	analyticsSvc    *analyticsApp.Service
	searchSvc       *searchApp.Service
	exportSvc       *exportApp.Service
}

func NewHandler(
	planningSvc *planningApp.Service,
	executionSvc *executionApp.Service,
	verificationSvc *verificationApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		planningSvc:     planningSvc,
		executionSvc:    executionSvc,
		verificationSvc: verificationSvc,
		reportingSvc:    reportingSvc,
		analyticsSvc:    analyticsSvc,
		searchSvc:       searchSvc,
		exportSvc:       exportSvc,
	}
}

func (h *Handler) CreateAction(w http.ResponseWriter, r *http.Request) {
	var act action.Action
	if err := json.NewDecoder(r.Body).Decode(&act); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.planningSvc.CreateActionItem(r.Context(), &act); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(act)
}

func (h *Handler) ListActions(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	actions, err := h.reportingSvc.ListActions(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(actions)
}

func (h *Handler) GetAction(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, err := h.reportingSvc.GetAction(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(c)
}

func (h *Handler) AssignAction(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req struct {
		UserID string `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.executionSvc.AssignActionItem(r.Context(), id, req.UserID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ASSIGNED", "action_id": id})
}

func (h *Handler) SubmitEvidence(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.executionSvc.SubmitEvidence(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "EVIDENCE_SUBMITTED", "action_id": id})
}

func (h *Handler) ReviewAction(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rev effectivenessreview.Review
	if err := json.NewDecoder(r.Body).Decode(&rev); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.verificationSvc.ReviewActionItem(r.Context(), id, &rev); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(rev)
}

func (h *Handler) CloseAction(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.verificationSvc.CloseActionItem(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "CLOSED", "action_id": id})
}

func (h *Handler) SearchActions(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	actions, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": actions,
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
	w.Header().Set("Content-Disposition", "attachment; filename=action_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=action_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "action"})
}
