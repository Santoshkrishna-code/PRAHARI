package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /tenants", h.CreateTenant)
	mux.HandleFunc("GET /tenants", h.ListTenants)
	mux.HandleFunc("POST /organizations", h.CreateOrganization)
	mux.HandleFunc("POST /plants", h.CreatePlant)
	mux.HandleFunc("POST /feature-flags", h.SetFeatureFlag)
	mux.HandleFunc("POST /configurations", h.UpdateConfiguration)
	mux.HandleFunc("POST /licenses", h.AssignLicense)
	mux.HandleFunc("POST /administration/search", h.SearchAdministration)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
