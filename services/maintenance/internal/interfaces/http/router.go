package http

import (
	"net/http"

	maintenanceApp "prahari/services/maintenance/internal/application/maintenance"
	workorderApp "prahari/services/maintenance/internal/application/workorder"
	schedulingApp "prahari/services/maintenance/internal/application/scheduling"
	planningApp "prahari/services/maintenance/internal/application/planning"
	searchApp "prahari/services/maintenance/internal/application/search"
	reportingApp "prahari/services/maintenance/internal/application/reporting"
	exportApp "prahari/services/maintenance/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints matching path targets.
func RegisterRoutes(
	mux *http.ServeMux,
	maintenanceSvc *maintenanceApp.Service,
	workorderSvc *workorderApp.Service,
	schedulingSvc *schedulingApp.Service,
	planningSvc *planningApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(maintenanceSvc, workorderSvc, schedulingSvc, planningSvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /maintenance", handler.CreateMaintenance)
	mux.HandleFunc("GET /maintenance", handler.ListMaintenance)
	mux.HandleFunc("GET /maintenance/{id}", handler.GetMaintenance)

	mux.HandleFunc("POST /maintenance/{id}/approve", handler.ApproveMaintenance)
	mux.HandleFunc("POST /maintenance/{id}/schedule", handler.ScheduleMaintenance)
	mux.HandleFunc("POST /maintenance/{id}/assign", handler.AssignMaintenance)
	mux.HandleFunc("POST /maintenance/{id}/start", handler.StartMaintenance)
	mux.HandleFunc("POST /maintenance/{id}/complete", handler.CompleteMaintenance)
	mux.HandleFunc("POST /maintenance/{id}/verify", handler.VerifyMaintenance)
	mux.HandleFunc("POST /maintenance/{id}/close", handler.CloseMaintenance)

	mux.HandleFunc("POST /maintenance/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /maintenance/{id}/comments", handler.AddComment)
	mux.HandleFunc("POST /maintenance/search", handler.SearchMaintenance)

	mux.HandleFunc("GET /reports", handler.GetDashboardReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
