package http

import (
	"net/http"
)

// RegisterRoutes maps REST pathways.
func RegisterRoutes(
	mux *http.ServeMux,
	h *Handler,
) {
	mux.HandleFunc("/energy", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateEnergyProfile(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/energy/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetEnergyProfile(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/energy/meter-reading", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.RecordMeterReading(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/energy/forecast", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ForecastEnergyDemand(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/energy/optimization", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.RecommendOptimization(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/energy/target", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.DefineTarget(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/energy/search", func(w http.ResponseWriter, r *http.Request) {
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
}
