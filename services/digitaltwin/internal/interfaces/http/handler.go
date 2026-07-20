package http

import (
	"encoding/json"
	"net/http"
	"time"

	analyticsApp "prahari/services/digitaltwin/internal/application/analytics"
	exportApp "prahari/services/digitaltwin/internal/application/export"
	playbackApp "prahari/services/digitaltwin/internal/application/playback"
	searchApp "prahari/services/digitaltwin/internal/application/search"
	simulationApp "prahari/services/digitaltwin/internal/application/simulation"
	syncApp "prahari/services/digitaltwin/internal/application/synchronization"
	visualizationApp "prahari/services/digitaltwin/internal/application/visualization"
	searchDomain "prahari/services/digitaltwin/internal/domain/search"
	"prahari/services/digitaltwin/internal/domain/twin"
)

type Handler struct {
	syncSvc          *syncApp.Service
	visualizationSvc *visualizationApp.Service
	simulationSvc    *simulationApp.Service
	playbackSvc      *playbackApp.Service
	analyticsSvc     *analyticsApp.Service
	searchSvc        *searchApp.Service
	exportSvc        *exportApp.Service
}

func NewHandler(
	syncSvc *syncApp.Service,
	visualizationSvc *visualizationApp.Service,
	simulationSvc *simulationApp.Service,
	playbackSvc *playbackApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		syncSvc:          syncSvc,
		visualizationSvc: visualizationSvc,
		simulationSvc:    simulationSvc,
		playbackSvc:      playbackSvc,
		analyticsSvc:     analyticsSvc,
		searchSvc:        searchSvc,
		exportSvc:        exportSvc,
	}
}

func (h *Handler) CreateTwin(w http.ResponseWriter, r *http.Request) {
	var t twin.DigitalTwin
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.ID = "twin-" + time.Now().Format("20060102150405")
	t.Version = 1
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(t)
}

func (h *Handler) GetTwin(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	// return mock twin
	_ = json.NewEncoder(w).Encode(&twin.DigitalTwin{
		ID:        id,
		PlantID:   "P01",
		Name:      "Plant Reactor Twin",
		Status:    "ACTIVE",
		Version:   1,
		CreatedAt: time.Now(),
	})
}

func (h *Handler) ListTwins(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode([]string{})
}

func (h *Handler) RunSimulation(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TwinID     string `json:"twin_id"`
		Name       string `json:"name"`
		Parameters string `json:"parameters"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sc, err := h.simulationSvc.RunScenario(r.Context(), req.TwinID, req.Name, req.Parameters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(sc)
}

func (h *Handler) StartPlayback(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TwinID    string    `json:"twin_id"`
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
		Speed     float64   `json:"speed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ps, err := h.playbackSvc.StartPlayback(r.Context(), req.TwinID, req.StartTime, req.EndTime, req.Speed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(ps)
}

func (h *Handler) GetOverlays(w http.ResponseWriter, r *http.Request) {
	twinID := r.URL.Query().Get("twin_id")
	overlays, err := h.visualizationSvc.RenderTwinCanvas(r.Context(), twinID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(overlays)
}

func (h *Handler) SearchTwin(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	twins, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{"items": twins, "total": total})
}

func (h *Handler) GetPerceptionReport(w http.ResponseWriter, r *http.Request) {
	metrics, err := h.analyticsSvc.GetDigitalTwinMetrics(r.Context())
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
	w.Header().Set("Content-Disposition", "attachment; filename=twin_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=twin_summary.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "digitaltwin"})
}
