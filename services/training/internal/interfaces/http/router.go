package http

import (
	"net/http"

	trainingApp "prahari/services/training/internal/application/training"
	enrollmentApp "prahari/services/training/internal/application/enrollment"
	competencyApp "prahari/services/training/internal/application/competency"
	assessmentApp "prahari/services/training/internal/application/assessment"
	certificationApp "prahari/services/training/internal/application/certification"
	searchApp "prahari/services/training/internal/application/search"
	reportingApp "prahari/services/training/internal/application/reporting"
	exportApp "prahari/services/training/internal/application/export"
)

// RegisterRoutes binds HTTP endpoints matching path targets.
func RegisterRoutes(
	mux *http.ServeMux,
	trainingSvc *trainingApp.Service,
	enrollmentSvc *enrollmentApp.Service,
	competencySvc *competencyApp.Service,
	assessmentSvc *assessmentApp.Service,
	certificationSvc *certificationApp.Service,
	searchSvc *searchApp.Service,
	reportingSvc *reportingApp.Service,
	exportSvc *exportApp.Service,
) {
	handler := NewHandler(trainingSvc, enrollmentSvc, competencySvc, assessmentSvc, certificationSvc, searchSvc, reportingSvc, exportSvc)

	mux.HandleFunc("POST /training", handler.CreateTraining)
	mux.HandleFunc("GET /training", handler.ListTrainings)
	mux.HandleFunc("GET /training/{id}", handler.GetTraining)

	mux.HandleFunc("POST /training/{id}/schedule", handler.ScheduleTraining)
	mux.HandleFunc("POST /training/{id}/enroll", handler.EnrollTrainee)
	mux.HandleFunc("POST /training/{id}/attendance", handler.RecordAttendance)
	mux.HandleFunc("POST /training/{id}/assessment", handler.AssessTraining)
	mux.HandleFunc("POST /training/{id}/certify", handler.CertifyTraining)

	mux.HandleFunc("POST /training/{id}/attachments", handler.UploadAttachment)
	mux.HandleFunc("POST /training/{id}/comments", handler.AddComment)
	mux.HandleFunc("POST /training/search", handler.SearchTraining)

	mux.HandleFunc("GET /reports", handler.GetDashboardReport)
	mux.HandleFunc("GET /export/csv", handler.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", handler.ExportPDF)

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("GET /live", handler.Live)
	mux.HandleFunc("GET /version", handler.Version)
}
