package http

import (
	"net/http"

	contractorApp "prahari/services/contractor/internal/application/contractor"
	onboardingApp "prahari/services/contractor/internal/application/onboarding"
	complianceApp "prahari/services/contractor/internal/application/compliance"
	siteaccessApp "prahari/services/contractor/internal/application/siteaccess"
	searchApp "prahari/services/contractor/internal/application/search"
	reportingApp "prahari/services/contractor/internal/application/reporting"
	exportApp "prahari/services/contractor/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints matching path targets.
func RegisterRoutes(
	mux *http.ServeMux,
	contractorSvc *contractorApp.Service,
	onboardingSvc *onboardingApp.Service,
	complianceSvc *complianceApp.Service,
	siteaccessSvc *siteaccessApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(contractorSvc, onboardingSvc, complianceSvc, siteaccessSvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /contractors", handler.CreateContractor)
	mux.HandleFunc("GET /contractors", handler.ListContractors)
	mux.HandleFunc("GET /contractors/{id}", handler.GetContractor)

	mux.HandleFunc("POST /contractors/{id}/approve", handler.ApproveContractor)
	mux.HandleFunc("POST /contractors/{id}/activate", handler.ActivateContractor)
	mux.HandleFunc("POST /contractors/{id}/suspend", handler.SuspendContractor)
	mux.HandleFunc("POST /contractors/{id}/offboard", handler.OffboardContractor)

	mux.HandleFunc("POST /contractors/{id}/documents", handler.AddDocument)
	mux.HandleFunc("POST /contractors/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /contractors/search", handler.SearchContractors)

	mux.HandleFunc("GET /reports", handler.GetDashboardReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
