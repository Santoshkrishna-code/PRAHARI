package http

import (
	"net/http"

	incidentApp "prahari/services/incident/internal/application/incident"
	assignmentApp "prahari/services/incident/internal/application/assignment"
	investigationApp "prahari/services/incident/internal/application/investigation"
	searchApp "prahari/services/incident/internal/application/search"
	reportingApp "prahari/services/incident/internal/application/reporting"
	exportApp "prahari/services/incident/internal/application/export"
)

// RegisterRoutes binds all HTTP endpoints to the provided ServeMux.
func RegisterRoutes(
	mux *http.ServeMux,
	incidentSvc *incidentApp.Service,
	assignmentSvc *assignmentApp.Service,
	investigationSvc *investigationApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(incidentSvc, assignmentSvc, investigationSvc, searchSvc, reportingSvc, exportSvc)

	// Incident CRUD
	mux.HandleFunc("POST /incidents", handler.CreateIncident)
	mux.HandleFunc("GET /incidents", handler.ListIncidents)
	mux.HandleFunc("GET /incidents/{id}", handler.GetIncident)
	mux.HandleFunc("PUT /incidents/{id}", handler.UpdateIncident)
	mux.HandleFunc("DELETE /incidents/{id}", handler.DeleteIncident)

	// Incident lifecycle actions
	mux.HandleFunc("POST /incidents/{id}/assign", handler.AssignIncident)
	mux.HandleFunc("POST /incidents/{id}/investigate", handler.StartInvestigation)
	mux.HandleFunc("POST /incidents/{id}/resolve", handler.ResolveIncident)
	mux.HandleFunc("POST /incidents/{id}/close", handler.CloseIncident)

	// Sub-resources
	mux.HandleFunc("POST /incidents/{id}/comments", handler.AddComment)
	mux.HandleFunc("POST /incidents/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("GET /incidents/{id}/timeline", handler.GetTimeline)
	mux.HandleFunc("POST /incidents/{id}/capa", handler.CreateCAPA)

	// Search
	mux.HandleFunc("POST /incidents/search", handler.SearchIncidents)

	// Reporting & Export
	mux.HandleFunc("GET /reports/summary", handler.GetSummaryReport)
	mux.HandleFunc("GET /reports/metrics", handler.GetMetricsReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	// Health endpoints
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
