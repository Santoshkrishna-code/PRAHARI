package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"prahari/templates/service-template/internal/domain"
)

type HTTPHandler struct {
	useCase domain.PermitUseCase
}

func NewHTTPHandler(useCase domain.PermitUseCase) *HTTPHandler {
	return &HTTPHandler{useCase: useCase}
}

// RequestPermitDTO maps request body for permit requests
type RequestPermitDTO struct {
	WorkerID        string `json:"worker_id"`
	ZoneID          string `json:"zone_id"`
	DurationMinutes int    `json:"duration_minutes"`
}

// Standard JSON error helper
func writeError(w http.ResponseWriter, status int, code, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": map[string]string{
			"code":    code,
			"message": msg,
		},
	})
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Basic router implementation
	if r.Method == http.MethodPost && r.URL.Path == "/api/v1/permits" {
		h.handleRequestPermit(w, r)
		return
	}
	
	if r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/api/v1/permits/") {
		h.handleApprovePermit(w, r)
		return
	}
	
	writeError(w, http.StatusNotFound, "NOT_FOUND", "API endpoint not found")
}

func (h *HTTPHandler) handleRequestPermit(w http.ResponseWriter, r *http.Request) {
	var dto RequestPermitDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_ARGUMENT", "Malformed JSON request body")
		return
	}
	
	if dto.WorkerID == "" || dto.ZoneID == "" || dto.DurationMinutes <= 0 {
		writeError(w, http.StatusUnprocessableEntity, "VALIDATION_FAILED", "worker_id, zone_id, and positive duration_minutes are required")
		return
	}
	
	permit, err := h.useCase.RequestPermit(
		r.Context(),
		dto.WorkerID,
		dto.ZoneID,
		time.Duration(dto.DurationMinutes)*time.Minute,
	)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(permit)
}

func (h *HTTPHandler) handleApprovePermit(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		writeError(w, http.StatusBadRequest, "INVALID_ARGUMENT", "Missing permit ID in URL")
		return
	}
	permitID := pathParts[4]
	
	// Assume supervisor_id is injected into context by auth middleware
	supervisorID := r.Header.Get("X-User-ID")
	if supervisorID == "" {
		writeError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Missing X-User-ID authorization header")
		return
	}
	
	err := h.useCase.ApprovePermit(r.Context(), permitID, supervisorID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_STATE", err.Error())
		return
	}
	
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"SUCCESS"}`))
}
