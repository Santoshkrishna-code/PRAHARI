package http

import (
	"net/http"

	observationApp "prahari/services/observation/internal/application/observation"
	coachingApp "prahari/services/observation/internal/application/coaching"
	followupApp "prahari/services/observation/internal/application/followup"
	effectivenessApp "prahari/services/observation/internal/application/effectiveness"
	searchApp "prahari/services/observation/internal/application/search"
	reportingApp "prahari/services/observation/internal/application/reporting"
	exportApp "prahari/services/observation/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints matching path targets.
func RegisterRoutes(
	mux *http.ServeMux,
	observationSvc *observationApp.Service,
	coachingSvc *coachingApp.Service,
	followupSvc *followupApp.Service,
	effectivenessSvc *effectivenessApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(observationSvc, coachingSvc, followupSvc, effectivenessSvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /observations", handler.CreateObservation)
	mux.HandleFunc("GET /observations", handler.ListObservations)
	mux.HandleFunc("GET /observations/{id}", handler.GetObservation)

	mux.HandleFunc("POST /observations/{id}/coach", handler.CoachObservation)
	mux.HandleFunc("POST /observations/{id}/recognize", handler.RecognizeObservation)
	mux.HandleFunc("POST /observations/{id}/followup", handler.FollowUpObservation)
	mux.HandleFunc("POST /observations/{id}/verify", handler.VerifyObservation)
	mux.HandleFunc("POST /observations/{id}/close", handler.CloseObservation)

	mux.HandleFunc("POST /observations/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /observations/{id}/comments", handler.AddComment)
	mux.HandleFunc("POST /observations/search", handler.SearchObservations)

	mux.HandleFunc("GET /reports", handler.GetDashboardReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
