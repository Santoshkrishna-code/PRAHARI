package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /meetings", h.CreateMeeting)
	mux.HandleFunc("GET /meetings", h.ListMeetings)
	mux.HandleFunc("GET /meetings/{id}", h.GetMeeting)
	mux.HandleFunc("POST /meetings/{id}/start", h.StartMeeting)
	mux.HandleFunc("POST /meetings/{id}/attendance", h.RecordAttendance)
	mux.HandleFunc("POST /meetings/{id}/minutes", h.ApproveMinutes)
	mux.HandleFunc("POST /meetings/{id}/close", h.CloseMeeting)
	mux.HandleFunc("POST /meetings/search", h.SearchMeetings)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
