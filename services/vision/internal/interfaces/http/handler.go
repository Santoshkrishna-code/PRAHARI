package http

import (
	"encoding/json"
	"net/http"

	alertingApp "prahari/services/vision/internal/application/alerting"
	analyticsApp "prahari/services/vision/internal/application/analytics"
	detectionApp "prahari/services/vision/internal/application/detection"
	exportApp "prahari/services/vision/internal/application/export"
	inferenceApp "prahari/services/vision/internal/application/inference"
	searchApp "prahari/services/vision/internal/application/search"
	streamingApp "prahari/services/vision/internal/application/streaming"
	"prahari/services/vision/internal/domain/camera"
	searchDomain "prahari/services/vision/internal/domain/search"
)

type Handler struct {
	streamingSvc *streamingApp.Service
	inferenceSvc *inferenceApp.Service
	detectionSvc *detectionApp.Service
	alertingSvc  *alertingApp.Service
	analyticsSvc *analyticsApp.Service
	searchSvc    *searchApp.Service
	exportSvc    *exportApp.Service
}

func NewHandler(
	streamingSvc *streamingApp.Service,
	inferenceSvc *inferenceApp.Service,
	detectionSvc *detectionApp.Service,
	alertingSvc *alertingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		streamingSvc: streamingSvc,
		inferenceSvc: inferenceSvc,
		detectionSvc: detectionSvc,
		alertingSvc:  alertingSvc,
		analyticsSvc: analyticsSvc,
		searchSvc:    searchSvc,
		exportSvc:    exportSvc,
	}
}

func (h *Handler) RegisterCamera(w http.ResponseWriter, r *http.Request) {
	var c camera.Camera
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.streamingSvc.RegisterCamera(r.Context(), &c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(c)
}

func (h *Handler) ListCameras(w http.ResponseWriter, r *http.Request) {
	// mock output
	_ = json.NewEncoder(w).Encode([]string{})
}

func (h *Handler) RegisterModel(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "REGISTERED"})
}

func (h *Handler) RunInference(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CameraID string `json:"camera_id"`
		ModelID  string `json:"model_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	job, err := h.inferenceSvc.StartJob(r.Context(), req.CameraID, req.ModelID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(job)
}

func (h *Handler) GetDetections(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode([]string{})
}

func (h *Handler) CreateAlert(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CameraID string  `json:"camera_id"`
		Label    string  `json:"label"`
		BX       float64 `json:"bx"`
		BY       float64 `json:"by"`
		BW       float64 `json:"bw"`
		BH       float64 `json:"bh"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.alertingSvc.ProcessDetectionAlert(r.Context(), req.CameraID, req.Label, req.BX, req.BY, req.BW, req.BH); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ALERTED"})
}

func (h *Handler) SearchVisionDetections(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	detections, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{"items": detections, "total": total})
}

func (h *Handler) GetPerceptionReport(w http.ResponseWriter, r *http.Request) {
	metrics, err := h.analyticsSvc.GetPerceptionMetrics(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(metrics)
}

func (h *Handler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	criteria := &searchDomain.Criteria{CameraID: r.URL.Query().Get("camera_id")}
	data, err := h.exportSvc.ExportCSV(r.Context(), criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=vision_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=detection_summary.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "vision"})
}
