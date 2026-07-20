package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /moc", h.CreateMOC)
	mux.HandleFunc("GET /moc", h.ListMOCs)
	mux.HandleFunc("GET /moc/{id}", h.GetMOC)
	mux.HandleFunc("POST /moc/{id}/impact-assessment", h.SubmitImpactAssessment)
	mux.HandleFunc("POST /moc/{id}/technical-review", h.SubmitTechnicalReview)
	mux.HandleFunc("POST /moc/{id}/risk-review", h.SubmitRiskReview)
	mux.HandleFunc("POST /moc/{id}/approval", h.SubmitApproval)
	mux.HandleFunc("POST /moc/{id}/implementation", h.StartImplementation)
	mux.HandleFunc("POST /moc/{id}/verification", h.SubmitVerification)
	mux.HandleFunc("POST /moc/search", h.SearchMOCs)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
