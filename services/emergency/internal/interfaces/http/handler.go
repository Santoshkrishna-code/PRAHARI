package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/emergency/internal/application/analytics"
	drillApp "prahari/services/emergency/internal/application/drill"
	evacuationApp "prahari/services/emergency/internal/application/evacuation"
	exportApp "prahari/services/emergency/internal/application/export"
	planningApp "prahari/services/emergency/internal/application/planning"
	recoveryApp "prahari/services/emergency/internal/application/recovery"
	reportingApp "prahari/services/emergency/internal/application/reporting"
	responseApp "prahari/services/emergency/internal/application/response"
	searchApp "prahari/services/emergency/internal/application/search"
	"prahari/services/emergency/internal/domain/emergency"
	"prahari/services/emergency/internal/domain/evacuation"
	"prahari/services/emergency/internal/domain/recovery"
	searchDomain "prahari/services/emergency/internal/domain/search"
)

type Handler struct {
	planningSvc   *planningApp.Service
	responseSvc   *responseApp.Service
	evacuationSvc *evacuationApp.Service
	drillSvc      *drillApp.Service
	recoverySvc   *recoveryApp.Service
	reportingSvc  *reportingApp.Service
	analyticsSvc  *analyticsApp.Service
	searchSvc     *searchApp.Service
	exportSvc     *exportApp.Service
}

func NewHandler(
	planningSvc *planningApp.Service,
	responseSvc *responseApp.Service,
	evacuationSvc *evacuationApp.Service,
	drillSvc *drillApp.Service,
	recoverySvc *recoveryApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		planningSvc:   planningSvc,
		responseSvc:   responseSvc,
		evacuationSvc: evacuationSvc,
		drillSvc:      drillSvc,
		recoverySvc:   recoverySvc,
		reportingSvc:  reportingSvc,
		analyticsSvc:  analyticsSvc,
		searchSvc:     searchSvc,
		exportSvc:     exportSvc,
	}
}

func (h *Handler) DeclareEmergency(w http.ResponseWriter, r *http.Request) {
	var em emergency.Emergency
	if err := json.NewDecoder(r.Body).Decode(&em); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.responseSvc.DeclareEmergency(r.Context(), &em); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(em)
}

func (h *Handler) ListEmergencies(w http.ResponseWriter, r *http.Request) {
	plantID := r.URL.Query().Get("plant_id")
	emergencies, err := h.reportingSvc.ListEmergencies(r.Context(), plantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(emergencies)
}

func (h *Handler) GetEmergency(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	em, err := h.reportingSvc.GetEmergency(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(em)
}

func (h *Handler) ActivateEmergency(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.responseSvc.ActivateResponse(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "RESPONSE_ACTIVATED", "emergency_id": id})
}

func (h *Handler) DeployResources(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req struct {
		ResourceID string `json:"resource_id"`
		Quantity   int    `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.responseSvc.DeployResource(r.Context(), id, req.ResourceID, req.Quantity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "RESOURCE_DEPLOYED", "emergency_id": id})
}

func (h *Handler) InitiateEvacuation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var rec evacuation.Record
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rec.EmergencyID = id
	if err := h.evacuationSvc.InitiateEvacuation(r.Context(), &rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) InitiateRecovery(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var plan recovery.Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	plan.EmergencyID = id
	if err := h.recoverySvc.InitiateRecovery(r.Context(), &plan); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(plan)
}

func (h *Handler) SearchEmergencies(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	emergencies, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{
		"items": emergencies,
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
	w.Header().Set("Content-Disposition", "attachment; filename=emergency_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=emergency_report.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "emergency"})
}
