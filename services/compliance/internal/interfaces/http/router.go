package http

import (
	"net/http"

	complianceApp "prahari/services/compliance/internal/application/compliance"
	obligationApp "prahari/services/compliance/internal/application/obligation"
	reviewApp "prahari/services/compliance/internal/application/review"
	evidenceApp "prahari/services/compliance/internal/application/evidence"
	searchApp "prahari/services/compliance/internal/application/search"
	reportingApp "prahari/services/compliance/internal/application/reporting"
	exportApp "prahari/services/compliance/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints matching path targets.
func RegisterRoutes(
	mux *http.ServeMux,
	complianceSvc *complianceApp.Service,
	obligationSvc *obligationApp.Service,
	reviewSvc *reviewApp.Service,
	evidenceSvc *evidenceApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(complianceSvc, obligationSvc, reviewSvc, evidenceSvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /compliance", handler.CreateCompliance)
	mux.HandleFunc("GET /compliance", handler.ListCompliance)
	mux.HandleFunc("GET /compliance/{id}", handler.GetCompliance)

	mux.HandleFunc("POST /compliance/{id}/review", handler.ReviewCompliance)
	mux.HandleFunc("POST /compliance/{id}/evidence", handler.EvidenceCompliance)
	mux.HandleFunc("POST /compliance/{id}/monitor", handler.MonitorCompliance)

	mux.HandleFunc("POST /compliance/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /compliance/{id}/comments", handler.AddComment)
	mux.HandleFunc("POST /compliance/search", handler.SearchCompliance)

	mux.HandleFunc("GET /reports", handler.GetDashboardReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
