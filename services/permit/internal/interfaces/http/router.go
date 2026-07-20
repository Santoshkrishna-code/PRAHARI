package http

import (
	"net/http"

	permitApp "prahari/services/permit/internal/application/permit"
	approvalApp "prahari/services/permit/internal/application/approval"
	riskApp "prahari/services/permit/internal/application/riskassessment"
	searchApp "prahari/services/permit/internal/application/search"
	reportingApp "prahari/services/permit/internal/application/reporting"
	exportApp "prahari/services/permit/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints.
func RegisterRoutes(
	mux *http.ServeMux,
	permitSvc *permitApp.Service,
	approvalSvc *approvalApp.Service,
	riskSvc *riskApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(permitSvc, approvalSvc, riskSvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /permits", handler.CreatePermit)
	mux.HandleFunc("GET /permits", handler.ListPermits)
	mux.HandleFunc("GET /permits/{id}", handler.GetPermit)
	mux.HandleFunc("PUT /permits/{id}", handler.UpdatePermit)

	mux.HandleFunc("POST /permits/{id}/submit", handler.SubmitPermit)
	mux.HandleFunc("POST /permits/{id}/approve", handler.ApprovePermit)
	mux.HandleFunc("POST /permits/{id}/reject", handler.RejectPermit)
	mux.HandleFunc("POST /permits/{id}/issue", handler.IssuePermit)
	mux.HandleFunc("POST /permits/{id}/activate", handler.ActivatePermit)
	mux.HandleFunc("POST /permits/{id}/suspend", handler.SuspendPermit)
	mux.HandleFunc("POST /permits/{id}/complete", handler.CompletePermit)
	mux.HandleFunc("POST /permits/{id}/close", handler.ClosePermit)
	mux.HandleFunc("POST /permits/{id}/extend", handler.ExtendPermit)

	mux.HandleFunc("POST /permits/{id}/comments", handler.AddComment)
	mux.HandleFunc("POST /permits/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /permits/search", handler.SearchPermits)

	mux.HandleFunc("GET /reports", handler.GetMetricsReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
