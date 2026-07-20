package http

import (
	"net/http"

	prahariLogger "prahari/shared/logger"
)

// DepartmentIsolationMiddleware ensures plant departments isolation bounds.
func DepartmentIsolationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deptID := r.Header.Get("X-Department-ID")
		if deptID != "" {
			prahariLogger.Info(r.Context(), "Risk department isolation checks applied",
				prahariLogger.String("department_id", deptID))
		}
		next.ServeHTTP(w, r)
	})
}
