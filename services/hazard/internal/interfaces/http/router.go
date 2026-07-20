package http

import (
	"net/http"

	hazardApp "prahari/services/hazard/internal/application/hazard"
	assessmentApp "prahari/services/hazard/internal/application/assessment"
	mitigationApp "prahari/services/hazard/internal/application/mitigation"
	verifyApp "prahari/services/hazard/internal/application/verification"
	searchApp "prahari/services/hazard/internal/application/search"
	reportingApp "prahari/services/hazard/internal/application/reporting"
	exportApp "prahari/services/hazard/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints matching path targets.
func RegisterRoutes(
	mux *http.ServeMux,
	hazardSvc *hazardApp.Service,
	assessmentSvc *assessmentApp.Service,
	mitigationSvc *mitigationApp.Service,
	verifySvc *verifyApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(hazardSvc, assessmentSvc, mitigationSvc, verifySvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /hazards", handler.CreateHazard)
	mux.HandleFunc("GET /hazards", handler.ListHazards)
	mux.HandleFunc("GET /hazards/{id}", handler.GetHazard)

	mux.HandleFunc("POST /hazards/{id}/assess", handler.AssessHazard)
	mux.HandleFunc("POST /hazards/{id}/mitigate", handler.MitigateHazard)
	mux.HandleFunc("POST /hazards/{id}/verify", handler.VerifyHazard)
	mux.HandleFunc("POST /hazards/{id}/close", handler.CloseHazard)

	mux.HandleFunc("POST /hazards/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /hazards/{id}/comments", handler.AddComment)
	mux.HandleFunc("POST /hazards/search", handler.SearchHazards)

	mux.HandleFunc("GET /reports", handler.GetDashboardReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
