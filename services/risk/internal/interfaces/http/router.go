package http

import (
	"net/http"

	assessmentApp "prahari/services/risk/internal/application/assessment"
	reviewApp "prahari/services/risk/internal/application/review"
	approvalApp "prahari/services/risk/internal/application/approval"
	residualApp "prahari/services/risk/internal/application/residual"
	searchApp "prahari/services/risk/internal/application/search"
	reportingApp "prahari/services/risk/internal/application/reporting"
	exportApp "prahari/services/risk/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints matching path targets.
func RegisterRoutes(
	mux *http.ServeMux,
	assessmentSvc *assessmentApp.Service,
	reviewSvc *reviewApp.Service,
	approvalSvc *approvalApp.Service,
	residualSvc *residualApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(assessmentSvc, reviewSvc, approvalSvc, residualSvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /risk-assessments", handler.CreateRiskAssessment)
	mux.HandleFunc("GET /risk-assessments", handler.ListRiskAssessments)
	mux.HandleFunc("GET /risk-assessments/{id}", handler.GetRiskAssessment)

	mux.HandleFunc("POST /risk-assessments/{id}/assess", handler.AssessRisk)
	mux.HandleFunc("POST /risk-assessments/{id}/review", handler.ReviewRisk)
	mux.HandleFunc("POST /risk-assessments/{id}/approve", handler.ApproveRisk)
	mux.HandleFunc("POST /risk-assessments/{id}/activate", handler.ActivateRisk)
	mux.HandleFunc("POST /risk-assessments/{id}/reassess", handler.ReassessRisk)

	mux.HandleFunc("POST /risk-assessments/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /risk-assessments/{id}/comments", handler.AddComment)
	mux.HandleFunc("POST /risk-assessments/search", handler.SearchRiskAssessments)

	mux.HandleFunc("GET /reports", handler.GetDashboardReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
