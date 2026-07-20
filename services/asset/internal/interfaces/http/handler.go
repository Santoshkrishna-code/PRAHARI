package http

import (
	"encoding/json"
	"net/http"

	prahariErrors "prahari/shared/errors"

	assetApp "prahari/services/asset/internal/application/asset"
	lifecycleApp "prahari/services/asset/internal/application/lifecycle"
	searchApp "prahari/services/asset/internal/application/search"
	reportingApp "prahari/services/asset/internal/application/reporting"
	exportApp "prahari/services/asset/internal/application/export"
	searchDomain "prahari/services/asset/internal/domain/search"
)

// Handler binds HTTP requests.
type Handler struct {
	asset     *assetApp.Service
	lifecycle *lifecycleApp.Service
	search    *searchApp.Service
	reporting *reportingApp.Service
	export    *exportApp.Service
}

// NewHandler instantiates Handler.
func NewHandler(
	asset *assetApp.Service,
	lifecycle *lifecycleApp.Service,
	search *searchApp.Service,
	reporting *reportingApp.Service,
	export *exportApp.Service,
) *Handler {
	return &Handler{
		asset:     asset,
		lifecycle: lifecycle,
		search:    search,
		reporting: reporting,
		export:    export,
	}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// CreateAsset handles POST /assets.
func (h *Handler) CreateAsset(w http.ResponseWriter, r *http.Request) {
	var cmd assetApp.RegisterAssetCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid request body", err))
		return
	}

	as, err := h.asset.RegisterAsset(r.Context(), cmd, "actor-id")
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, as)
}

// GetAsset handles GET /assets/{id}.
func (h *Handler) GetAsset(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	as, err := h.asset.GetAsset(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("asset not found", err))
		return
	}
	writeJSON(w, http.StatusOK, as)
}

// ListAssets handles GET /assets.
func (h *Handler) ListAssets(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": []string{}, "total_count": 0})
}

// CommissionAsset handles POST /assets/{id}/commission.
func (h *Handler) CommissionAsset(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := assetApp.TransitionStatusCommand{
		AssetID:    id,
		TargetCode: "COMMISSIONED",
		ActorID:    "manager-id",
	}
	if err := h.asset.TransitionLifecycle(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "commissioned"})
}

// ActivateAsset handles POST /assets/{id}/activate.
func (h *Handler) ActivateAsset(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := assetApp.TransitionStatusCommand{
		AssetID:    id,
		TargetCode: "OPERATIONAL",
		ActorID:    "manager-id",
	}
	if err := h.asset.TransitionLifecycle(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "activated"})
}

// SendToMaintenance handles POST /assets/{id}/maintenance.
func (h *Handler) SendToMaintenance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := assetApp.TransitionStatusCommand{
		AssetID:    id,
		TargetCode: "MAINTENANCE",
		ActorID:    "maintenance-officer-id",
	}
	if err := h.asset.TransitionLifecycle(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "maintenance started"})
}

// RetireAsset handles POST /assets/{id}/retire.
func (h *Handler) RetireAsset(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	cmd := assetApp.TransitionStatusCommand{
		AssetID:    id,
		TargetCode: "DECOMMISSIONED",
		ActorID:    "manager-id",
	}
	if err := h.asset.TransitionLifecycle(r.Context(), cmd); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "decommissioned"})
}

// AddDocument handles POST /assets/{id}/documents.
func (h *Handler) AddDocument(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "document registered"})
}

// UploadAttachment handles POST /assets/{id}/attachments.
func (h *Handler) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]string{"status": "attachment uploaded"})
}

// SearchAssets handles POST /assets/search.
func (h *Handler) SearchAssets(w http.ResponseWriter, r *http.Request) {
	var criteria searchDomain.Criteria
	if err := json.NewDecoder(r.Body).Decode(&criteria); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	res, err := h.search.Search(r.Context(), &criteria)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, res)
}

// GetDashboardReport handles GET /reports.
func (h *Handler) GetDashboardReport(w http.ResponseWriter, r *http.Request) {
	report, err := h.reporting.GenerateDashboardReport(r.Context())
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}
	writeJSON(w, http.StatusOK, report)
}

// ExportCSV handles GET /export/csv.
func (h *Handler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	criteria := &searchDomain.Criteria{}
	data, err := h.export.ExportCSV(r.Context(), criteria)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=assets.csv")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// ExportPDF handles GET /export/pdf/{id}.
func (h *Handler) ExportPDF(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	data, err := h.export.ExportPDF(r.Context(), id)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=asset_specification.pdf")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// Health handles GET /health.
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
}

// Ready handles GET /ready.
func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

// Live handles GET /live.
func (h *Handler) Live(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "alive"})
}

// Version handles GET /version.
func (h *Handler) Version(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"service": "asset-service",
		"version": "1.0.0",
	})
}
