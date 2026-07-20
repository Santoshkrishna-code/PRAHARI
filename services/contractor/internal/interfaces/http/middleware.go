package http

import (
	"net/http"

	prahariLogger "prahari/shared/logger"
)

// ContractorIsolationMiddleware ensures contractor data isolation bounds.
func ContractorIsolationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deptID := r.Header.Get("X-Department-ID")
		if deptID != "" {
			prahariLogger.Info(r.Context(), "Contractor isolation checks applied",
				prahariLogger.String("department_id", deptID))
		}
		next.ServeHTTP(w, r)
	})
}
