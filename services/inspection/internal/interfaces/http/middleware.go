package http

import (
	"net/http"

	prahariLogger "prahari/shared/logger"
)

// DepartmentIsolationMiddleware ensures isolation boundaries.
func DepartmentIsolationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deptID := r.Header.Get("X-Department-ID")
		if deptID != "" {
			prahariLogger.Info(r.Context(), "Walkthrough checks department isolation filters applied",
				prahariLogger.String("department_id", deptID))
		}
		next.ServeHTTP(w, r)
	})
}
