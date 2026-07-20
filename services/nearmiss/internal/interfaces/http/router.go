package http

import (
	"net/http"

	nearmissApp "prahari/services/nearmiss/internal/application/nearmiss"
	investigationApp "prahari/services/nearmiss/internal/application/investigation"
	correctiveApp "prahari/services/nearmiss/internal/application/corrective"
	verifyApp "prahari/services/nearmiss/internal/application/verification"
	searchApp "prahari/services/nearmiss/internal/application/search"
	reportingApp "prahari/services/nearmiss/internal/application/reporting"
	exportApp "prahari/services/nearmiss/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints matching path targets.
func RegisterRoutes(
	mux *http.ServeMux,
	nearmissSvc *nearmissApp.Service,
	investigationSvc *investigationApp.Service,
	correctiveSvc *correctiveApp.Service,
	verifySvc *verifyApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(nearmissSvc, investigationSvc, correctiveSvc, verifySvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /near-misses", handler.CreateNearMiss)
	mux.HandleFunc("GET /near-misses", handler.ListNearMisses)
	mux.HandleFunc("GET /near-misses/{id}", handler.GetNearMiss)

	mux.HandleFunc("POST /near-misses/{id}/investigate", handler.InvestigateNearMiss)
	mux.HandleFunc("POST /near-misses/{id}/corrective-actions", handler.AddCorrectiveAction)
	mux.HandleFunc("POST /near-misses/{id}/verify", handler.VerifyNearMiss)
	mux.HandleFunc("POST /near-misses/{id}/close", handler.CloseNearMiss)
	mux.HandleFunc("POST /near-misses/{id}/escalate", handler.EscalateNearMiss)

	mux.HandleFunc("POST /near-misses/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /near-misses/{id}/comments", handler.AddComment)
	mux.HandleFunc("POST /near-misses/search", handler.SearchNearMisses)

	mux.HandleFunc("GET /reports", handler.GetDashboardReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
