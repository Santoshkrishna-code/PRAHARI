package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/ai/internal/application/analytics"
	chatApp "prahari/services/ai/internal/application/chat"
	exportApp "prahari/services/ai/internal/application/export"
	predictionApp "prahari/services/ai/internal/application/prediction"
	ragApp "prahari/services/ai/internal/application/rag"
	recommendationApp "prahari/services/ai/internal/application/recommendation"
	searchApp "prahari/services/ai/internal/application/search"
	summarizationApp "prahari/services/ai/internal/application/summarization"
	searchDomain "prahari/services/ai/internal/domain/search"
)

type Handler struct {
	ragSvc            *ragApp.Service
	chatSvc           *chatApp.Service
	recommendationSvc *recommendationApp.Service
	summarizationSvc  *summarizationApp.Service
	predictionSvc     *predictionApp.Service
	analyticsSvc      *analyticsApp.Service
	searchSvc         *searchApp.Service
	exportSvc         *exportApp.Service
}

func NewHandler(
	ragSvc *ragApp.Service,
	chatSvc *chatApp.Service,
	recommendationSvc *recommendationApp.Service,
	summarizationSvc *summarizationApp.Service,
	predictionSvc *predictionApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		ragSvc:            ragSvc,
		chatSvc:           chatSvc,
		recommendationSvc: recommendationSvc,
		summarizationSvc:  summarizationSvc,
		predictionSvc:     predictionSvc,
		analyticsSvc:      analyticsSvc,
		searchSvc:         searchSvc,
		exportSvc:         exportSvc,
	}
}

func (h *Handler) CopilotChat(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ThreadID string `json:"thread_id,omitempty"`
		Role     string `json:"role"`
		Content  string `json:"content"`
		UserID   string `json:"user_id,omitempty"`
		PlantID  string `json:"plant_id,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	threadID := req.ThreadID
	if threadID == "" {
		t, err := h.chatSvc.CreateThread(r.Context(), req.UserID, req.PlantID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		threadID = t.ID
	}

	resp, err := h.chatSvc.SubmitMessage(r.Context(), threadID, req.Role, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]any{
		"thread_id":  threadID,
		"answer":     resp.Answer,
		"confidence": resp.Confidence,
	})
}

func (h *Handler) SearchKnowledge(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Query string `json:"query"`
		Limit int    `json:"limit"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	results, err := h.ragSvc.SearchContext(r.Context(), req.Query, req.Limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(results)
}

func (h *Handler) Summarize(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SourceID string `json:"source_id"`
		Content  string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sum, err := h.summarizationSvc.SummarizeDocument(r.Context(), req.SourceID, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(sum)
}

func (h *Handler) Recommend(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PlantID  string `json:"plant_id"`
		SourceID string `json:"source_id"`
		Type     string `json:"type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rec, err := h.recommendationSvc.GenerateSuggestions(r.Context(), req.PlantID, req.SourceID, req.Type)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(rec)
}

func (h *Handler) Predict(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PlantID string `json:"plant_id"`
		Topic   string `json:"topic"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pred, err := h.predictionSvc.PredictRisk(r.Context(), req.PlantID, req.Topic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(pred)
}

func (h *Handler) IndexDocument(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SourceID string `json:"source_id"`
		Title    string `json:"title"`
		Content  string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.ragSvc.IndexDocument(r.Context(), req.SourceID, req.Title, req.Content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "INDEXED"})
}

func (h *Handler) ListConversations(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode([]string{})
}

func (h *Handler) GetExecutiveReport(w http.ResponseWriter, r *http.Request) {
	metrics, err := h.analyticsSvc.GetModelPerformance(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(metrics)
}

func (h *Handler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	criteria := &searchDomain.Criteria{Query: r.URL.Query().Get("q")}
	data, err := h.exportSvc.ExportCSV(r.Context(), criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=docs_report.csv")
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
	w.Header().Set("Content-Disposition", "attachment; filename=document_summary.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "ai"})
}
