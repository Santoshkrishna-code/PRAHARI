package http

import (
	"encoding/json"
	"net/http"

	alertsApp "prahari/services/analytics/internal/application/alerts"
	analyticsApp "prahari/services/analytics/internal/application/analytics"
	dashboardsApp "prahari/services/analytics/internal/application/dashboards"
	exportApp "prahari/services/analytics/internal/application/export"
	reportingApp "prahari/services/analytics/internal/application/reporting"
	searchApp "prahari/services/analytics/internal/application/search"
	"prahari/services/analytics/internal/domain/dashboard"
	searchDomain "prahari/services/analytics/internal/domain/search"
)

type Handler struct {
	dashboardsSvc *dashboardsApp.Service
	reportingSvc  *reportingApp.Service
	alertsSvc     *alertsApp.Service
	analyticsSvc  *analyticsApp.Service
	searchSvc     *searchApp.Service
	exportSvc     *exportApp.Service
}

func NewHandler(
	dashboardsSvc *dashboardsApp.Service,
	reportingSvc *reportingApp.Service,
	alertsSvc *alertsApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		dashboardsSvc: dashboardsSvc,
		reportingSvc:  reportingSvc,
		alertsSvc:     alertsSvc,
		analyticsSvc:  analyticsSvc,
		searchSvc:     searchSvc,
		exportSvc:     exportSvc,
	}
}

func (h *Handler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	d, err := h.dashboardsSvc.GetDashboard(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(d)
}

func (h *Handler) CreateDashboard(w http.ResponseWriter, r *http.Request) {
	var d dashboard.Dashboard
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.dashboardsSvc.CreateDashboard(r.Context(), &d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(d)
}

func (h *Handler) GetKPIs(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	kpis, err := h.analyticsSvc.GetKPIs(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(kpis)
}

func (h *Handler) GetMetrics(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	res, err := h.analyticsSvc.AnalyzeHSETrends(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(res)
}

func (h *Handler) CreateReport(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PlantID    string `json:"plant_id"`
		Title      string `json:"title"`
		ReportType string `json:"report_type"`
		UserID     string `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rep, err := h.reportingSvc.GenerateExecutiveReport(r.Context(), req.PlantID, req.Title, req.ReportType, req.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(rep)
}

func (h *Handler) GetReport(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	rep, err := h.reportingSvc.GetReport(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(rep)
}

func (h *Handler) SearchAnalytics(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	items, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{"items": items, "total": total})
}

func (h *Handler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	criteria := &searchDomain.Criteria{PlantID: r.URL.Query().Get("plant_id")}
	data, err := h.exportSvc.ExportCSV(r.Context(), criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=analytics_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=metric_summary.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "analytics"})
}
