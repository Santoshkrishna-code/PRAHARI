package http

import (
	"net/http"

	inspectionApp "prahari/services/inspection/internal/application/inspection"
	checklistApp "prahari/services/inspection/internal/application/checklist"
	actionApp "prahari/services/inspection/internal/application/action"
	scheduleApp "prahari/services/inspection/internal/application/schedule"
	searchApp "prahari/services/inspection/internal/application/search"
	reportingApp "prahari/services/inspection/internal/application/reporting"
	exportApp "prahari/services/inspection/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints matching path targets.
func RegisterRoutes(
	mux *http.ServeMux,
	inspectionSvc *inspectionApp.Service,
	checklistSvc *checklistApp.Service,
	actionSvc *actionApp.Service,
	scheduleSvc *scheduleApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(inspectionSvc, checklistSvc, actionSvc, scheduleSvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /inspections", handler.CreateInspection)
	mux.HandleFunc("GET /inspections", handler.ListInspections)
	mux.HandleFunc("GET /inspections/{id}", handler.GetInspection)
	mux.HandleFunc("PUT /inspections/{id}", handler.UpdateInspection)

	mux.HandleFunc("POST /inspections/{id}/schedule", handler.ScheduleInspection)
	mux.HandleFunc("POST /inspections/{id}/assign", handler.AssignInspection)
	mux.HandleFunc("POST /inspections/{id}/start", handler.StartInspection)
	mux.HandleFunc("POST /inspections/{id}/complete", handler.CompleteInspection)
	mux.HandleFunc("POST /inspections/{id}/approve", handler.ApproveInspection)

	mux.HandleFunc("POST /inspections/{id}/findings", handler.AddFinding)
	mux.HandleFunc("POST /inspections/{id}/actions", handler.AddAction)
	mux.HandleFunc("POST /inspections/{id}/comments", handler.AddComment)
	mux.HandleFunc("POST /inspections/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /inspections/search", handler.SearchInspections)

	mux.HandleFunc("GET /reports", handler.GetMetricsReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
