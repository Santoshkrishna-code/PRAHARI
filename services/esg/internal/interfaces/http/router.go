package http

import (
	"net/http"
)

// RegisterRoutes maps REST pathways.
func RegisterRoutes(
	mux *http.ServeMux,
	h *Handler,
) {
	mux.HandleFunc("/esg", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateESGProfile(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/esg/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetESGProfile(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/esg/carbon", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CalculateCarbon(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/esg/disclosure", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.PublishDisclosure(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/esg/objective", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreateObjective(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/esg/report", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreateReport(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/esg/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.SearchObjectives(w, r)
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
}
