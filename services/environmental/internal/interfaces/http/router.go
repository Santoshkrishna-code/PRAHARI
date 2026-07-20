package http

import (
	"net/http"
)

// RegisterRoutes registers environmental HTTP pathways.
func RegisterRoutes(
	mux *http.ServeMux,
	h *Handler,
) {
	mux.HandleFunc("/environment", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateEnvironmentalRecord(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/environment/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetEnvironmentalRecord(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/environment/monitor", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.RecordMonitoring(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/environment/sample", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.RecordSampling(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/environment/laboratory", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.EvaluateLaboratory(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/environment/evaluate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.EvaluateCompliance(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/environment/corrective-action", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreateCorrectiveAction(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/environment/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.SearchAspects(w, r)
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

	mux.HandleFunc("/permit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreatePermit(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/waste/solid", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.LogSolidWaste(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
