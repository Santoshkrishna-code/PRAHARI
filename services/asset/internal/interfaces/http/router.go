package http

import (
	"net/http"

	assetApp "prahari/services/asset/internal/application/asset"
	lifecycleApp "prahari/services/asset/internal/application/lifecycle"
	searchApp "prahari/services/asset/internal/application/search"
	reportingApp "prahari/services/asset/internal/application/reporting"
	exportApp "prahari/services/asset/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints matching path targets.
func RegisterRoutes(
	mux *http.ServeMux,
	assetSvc *assetApp.Service,
	lifecycleSvc *lifecycleApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(assetSvc, lifecycleSvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /assets", handler.CreateAsset)
	mux.HandleFunc("GET /assets", handler.ListAssets)
	mux.HandleFunc("GET /assets/{id}", handler.GetAsset)

	mux.HandleFunc("POST /assets/{id}/commission", handler.CommissionAsset)
	mux.HandleFunc("POST /assets/{id}/activate", handler.ActivateAsset)
	mux.HandleFunc("POST /assets/{id}/maintenance", handler.SendToMaintenance)
	mux.HandleFunc("POST /assets/{id}/retire", handler.RetireAsset)

	mux.HandleFunc("POST /assets/{id}/documents", handler.AddDocument)
	mux.HandleFunc("POST /assets/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /assets/search", handler.SearchAssets)

	mux.HandleFunc("GET /reports", handler.GetDashboardReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
