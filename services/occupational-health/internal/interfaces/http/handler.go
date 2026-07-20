package http

import (
	"encoding/json"
	"net/http"
	"strings"

	prahariErrors "prahari/shared/errors"

	clearanceApp "prahari/services/occupational-health/internal/application/clearance"
	exportApp "prahari/services/occupational-health/internal/application/export"
	medicalApp "prahari/services/occupational-health/internal/application/medical"
	reportingApp "prahari/services/occupational-health/internal/application/reporting"
	searchApp "prahari/services/occupational-health/internal/application/search"
	"prahari/services/occupational-health/internal/domain/appointment"
	"prahari/services/occupational-health/internal/domain/fitnessassessment"
	"prahari/services/occupational-health/internal/domain/healthprofile"
	"prahari/services/occupational-health/internal/domain/medicalexamination"
	"prahari/services/occupational-health/internal/domain/medicalrecord"
	"prahari/services/occupational-health/internal/domain/restriction"
	"prahari/services/occupational-health/internal/domain/search"
)

type Handler struct {
	clearance *clearanceApp.Service
	medical   *medicalApp.Service
	searchSvc *searchApp.Service
	reportSvc *reportingApp.Service
	exportSvc *exportApp.Service
}

func NewHandler(
	clearance *clearanceApp.Service,
	medical *medicalApp.Service,
	searchSvc *searchApp.Service,
	reportSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) *Handler {
	return &Handler{
		clearance: clearance,
		medical:   medical,
		searchSvc: searchSvc,
		reportSvc: reportSvc,
		exportSvc: exportSvc,
	}
}

// helper to strip URL prefix and get ID.
func getID(path, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}

func (h *Handler) CreateHealthProfile(w http.ResponseWriter, r *http.Request) {
	var req healthprofile.HealthProfile
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.clearance.CreateProfile(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) RecordMedicalExamination(w http.ResponseWriter, r *http.Request) {
	var req medicalexamination.MedicalExamination
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.medical.RecordExamination(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	// Trigger fitness assessment transition check
	_ = h.clearance.TransitionStatus(r.Context(), clearanceApp.TransitionStatusCommand{
		ProfileID:  req.HealthProfileID,
		TargetCode: "MEDICAL_EXAMINATION",
		ActorID:    req.PhysicianID,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) AssessFitness(w http.ResponseWriter, r *http.Request) {
	var req fitnessassessment.FitnessAssessment
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.clearance.TransitionStatus(r.Context(), clearanceApp.TransitionStatusCommand{
		ProfileID:  req.HealthProfileID,
		TargetCode: "FITNESS_ASSESSMENT",
		ActorID:    req.EvaluatorID,
	}) ; err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) ApplyRestriction(w http.ResponseWriter, r *http.Request) {
	var req restriction.MedicalRestriction
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.clearance.AddRestriction(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) SearchProfiles(w http.ResponseWriter, r *http.Request) {
	var req search.Criteria
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid search criteria", err))
		return
	}

	results, err := h.searchSvc.SearchProfiles(r.Context(), req)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(results)
}

func (h *Handler) GetDashboardReport(w http.ResponseWriter, r *http.Request) {
	results, err := h.reportSvc.GetDashboardReport(r.Context())
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(results)
}

func (h *Handler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=health_profiles.csv")
	if err := h.exportSvc.WriteCSVReport(r.Context(), w); err != nil {
		prahariErrors.WriteHTTP(w, err)
	}
}

func (h *Handler) ExportPDF(w http.ResponseWriter, r *http.Request) {
	profileID := getID(r.URL.Path, "/export/pdf/")
	if profileID == "" {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("profile ID is required", nil))
		return
	}

	pdfData, err := h.exportSvc.GetPDFReport(r.Context(), profileID)
	if err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment;filename=fitness_card.pdf")
	_, _ = w.Write(pdfData)
}

func (h *Handler) AddMedicalRecord(w http.ResponseWriter, r *http.Request) {
	var req medicalrecord.MedicalRecord
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.medical.CreateMedicalRecord(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}

func (h *Handler) ScheduleAppointment(w http.ResponseWriter, r *http.Request) {
	var req appointment.Appointment
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewBadRequestError("invalid body payload", err))
		return
	}

	if err := h.medical.ScheduleAppointment(r.Context(), &req); err != nil {
		prahariErrors.WriteHTTP(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(req)
}
