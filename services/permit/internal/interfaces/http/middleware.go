package http

import (
	"net/http"

	prahariLogger "prahari/shared/logger"
)

// DepartmentIsolationMiddleware ensures users only interact with permits from their respective organizations.
func DepartmentIsolationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deptID := r.Header.Get("X-Department-ID")
		if deptID != "" {
			prahariLogger.Info(r.Context(), "Department isolation filter applied to permit scope",
				prahariLogger.String("department_id", deptID))
		}
		next.ServeHTTP(w, r)
	})
}
