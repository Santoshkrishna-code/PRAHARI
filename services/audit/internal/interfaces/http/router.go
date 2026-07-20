package http

import (
	"net/http"

	auditApp "prahari/services/audit/internal/application/audit"
	planningApp "prahari/services/audit/internal/application/planning"
	executionApp "prahari/services/audit/internal/application/execution"
	reviewApp "prahari/services/audit/internal/application/review"
	findingsApp "prahari/services/audit/internal/application/findings"
	searchApp "prahari/services/audit/internal/application/search"
	reportingApp "prahari/services/audit/internal/application/reporting"
	exportApp "prahari/services/audit/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints matching path targets.
func RegisterRoutes(
	mux *http.ServeMux,
	auditSvc *auditApp.Service,
	planningSvc *planningApp.Service,
	executionSvc *executionApp.Service,
	reviewSvc *reviewApp.Service,
	findingsSvc *findingsApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(auditSvc, planningSvc, executionSvc, reviewSvc, findingsSvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /audits", handler.CreateAudit)
	mux.HandleFunc("GET /audits", handler.ListAudits)
	mux.HandleFunc("GET /audits/{id}", handler.GetAudit)

	mux.HandleFunc("POST /audits/{id}/schedule", handler.ScheduleAudit)
	mux.HandleFunc("POST /audits/{id}/execute", handler.ExecuteAudit)
	mux.HandleFunc("POST /audits/{id}/review", handler.ReviewAudit)
	mux.HandleFunc("POST /audits/{id}/approve", handler.ApproveAudit)
	mux.HandleFunc("POST /audits/{id}/close", handler.CloseAudit)

	mux.HandleFunc("POST /audits/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /audits/{id}/comments", handler.AddComment)
	mux.HandleFunc("POST /audits/search", handler.SearchAudits)

	mux.HandleFunc("GET /reports", handler.GetDashboardReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
