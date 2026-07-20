package http

import (
	"net/http"
)

// RegisterRoutes maps REST endpoints to target request handlers.
func RegisterRoutes(
	mux *http.ServeMux,
	h *Handler,
) {
	mux.HandleFunc("/health-profiles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateHealthProfile(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/health-profiles/medical-examination", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.RecordMedicalExamination(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/health-profiles/fitness-assessment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.AssessFitness(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/health-profiles/restriction", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ApplyRestriction(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/health/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.SearchProfiles(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/reports", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.GetDashboardReport(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/export/csv", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.ExportCSV(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/export/pdf/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.ExportPDF(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/medical-record", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.AddMedicalRecord(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/appointment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ScheduleAppointment(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
