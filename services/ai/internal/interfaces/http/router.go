package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /copilot/chat", h.CopilotChat)
	mux.HandleFunc("POST /search", h.SearchKnowledge)
	mux.HandleFunc("POST /summarize", h.Summarize)
	mux.HandleFunc("POST /recommend", h.Recommend)
	mux.HandleFunc("POST /predict", h.Predict)
	mux.HandleFunc("POST /documents/index", h.IndexDocument)
	mux.HandleFunc("GET /conversations", h.ListConversations)
	mux.HandleFunc("GET /metrics", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
