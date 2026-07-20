package http

import (
	"encoding/json"
	"net/http"

	analyticsApp "prahari/services/administration/internal/application/analytics"
	configurationApp "prahari/services/administration/internal/application/configuration"
	exportApp "prahari/services/administration/internal/application/export"
	hierarchyApp "prahari/services/administration/internal/application/hierarchy"
	licensingApp "prahari/services/administration/internal/application/licensing"
	reportingApp "prahari/services/administration/internal/application/reporting"
	searchApp "prahari/services/administration/internal/application/search"
	tenantApp "prahari/services/administration/internal/application/tenant"
	"prahari/services/administration/internal/domain/configuration"
	"prahari/services/administration/internal/domain/featureflag"
	"prahari/services/administration/internal/domain/license"
	"prahari/services/administration/internal/domain/organization"
	"prahari/services/administration/internal/domain/plant"
	searchDomain "prahari/services/administration/internal/domain/search"
	"prahari/services/administration/internal/domain/tenant"
)

type Handler struct {
	tenantSvc        *tenantApp.Service
	hierarchySvc     *hierarchyApp.Service
	configurationSvc *configurationApp.Service
	licensingSvc     *licensingApp.Service
	reportingSvc     *reportingApp.Service
	analyticsSvc     *analyticsApp.Service
	searchSvc        *searchApp.Service
	exportSvc        *exportApp.Service
}

func NewHandler(
	tenantSvc *tenantApp.Service,
	hierarchySvc *hierarchyApp.Service,
	configurationSvc *configurationApp.Service,
	licensingSvc *licensingApp.Service,
	reportingSvc *reportingApp.Service,
	analyticsSvc *analyticsApp.Service,
	searchSvc *searchApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		tenantSvc:        tenantSvc,
		hierarchySvc:     hierarchySvc,
		configurationSvc: configurationSvc,
		licensingSvc:     licensingSvc,
		reportingSvc:     reportingSvc,
		analyticsSvc:     analyticsSvc,
		searchSvc:        searchSvc,
		exportSvc:        exportSvc,
	}
}

func (h *Handler) CreateTenant(w http.ResponseWriter, r *http.Request) {
	var t tenant.Tenant
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.tenantSvc.CreateTenant(r.Context(), &t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(t)
}

func (h *Handler) ListTenants(w http.ResponseWriter, r *http.Request) {
	tenants, err := h.reportingSvc.ListTenants(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(tenants)
}

func (h *Handler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var org organization.Organization
	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.hierarchySvc.CreateOrganization(r.Context(), &org); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(org)
}

func (h *Handler) CreatePlant(w http.ResponseWriter, r *http.Request) {
	var plt plant.Plant
	if err := json.NewDecoder(r.Body).Decode(&plt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.hierarchySvc.CreatePlant(r.Context(), &plt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(plt)
}

func (h *Handler) SetFeatureFlag(w http.ResponseWriter, r *http.Request) {
	var flag featureflag.Flag
	if err := json.NewDecoder(r.Body).Decode(&flag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.configurationSvc.SetFeatureFlag(r.Context(), &flag); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(flag)
}

func (h *Handler) UpdateConfiguration(w http.ResponseWriter, r *http.Request) {
	var param configuration.Param
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.configurationSvc.UpdateConfiguration(r.Context(), &param); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(param)
}

func (h *Handler) AssignLicense(w http.ResponseWriter, r *http.Request) {
	var lic license.License
	if err := json.NewDecoder(r.Body).Decode(&lic); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.licensingSvc.AssignLicense(r.Context(), &lic); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(lic)
}

func (h *Handler) SearchAdministration(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tenants, total, err := h.searchSvc.ExecuteSearch(r.Context(), &criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]any{"items": tenants, "total": total})
}

func (h *Handler) GetExecutiveReport(w http.ResponseWriter, r *http.Request) {
	metrics, err := h.reportingSvc.GetExecutiveMetrics(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(metrics)
}

func (h *Handler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	criteria := &searchDomain.Criteria{TenantID: r.URL.Query().Get("tenant_id")}
	data, err := h.exportSvc.ExportCSV(r.Context(), criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=tenants_report.csv")
	_, _ = w.Write(data)
}

func (h *Handler) ExportPDF(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	data, err := h.exportSvc.ExportPDF(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=tenant_provision.pdf")
	_, _ = w.Write(data)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "UP", "service": "administration"})
}
