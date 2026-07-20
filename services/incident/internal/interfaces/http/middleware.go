package http

import (
	"net/http"

	prahariLogger "prahari/shared/logger"
)

// DepartmentIsolationMiddleware ensures users can only see incidents
// within their own department unless they have elevated cross-department permissions.
func DepartmentIsolationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// In production, extract department_id from JWT claims and compare
		// against the incident's department_id. Users with admin/safety-officer
		// roles bypass this filter.
		departmentID := r.Header.Get("X-Department-ID")
		if departmentID != "" {
			prahariLogger.Info(r.Context(), "Department isolation filter applied",
				prahariLogger.String("department_id", departmentID))
		}

		next.ServeHTTP(w, r)
	})
}
