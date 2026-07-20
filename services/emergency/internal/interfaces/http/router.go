package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /emergencies", h.DeclareEmergency)
	mux.HandleFunc("GET /emergencies", h.ListEmergencies)
	mux.HandleFunc("GET /emergencies/{id}", h.GetEmergency)
	mux.HandleFunc("POST /emergencies/{id}/activate", h.ActivateEmergency)
	mux.HandleFunc("POST /emergencies/{id}/deploy", h.DeployResources)
	mux.HandleFunc("POST /emergencies/{id}/evacuation", h.InitiateEvacuation)
	mux.HandleFunc("POST /emergencies/{id}/recovery", h.InitiateRecovery)
	mux.HandleFunc("POST /emergencies/search", h.SearchEmergencies)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
