package http

import "net/http"

func TenantIsolationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantID := r.Header.Get("X-Tenant-ID")
		if tenantID != "" {
			r.Header.Set("X-Isolated-Tenant-ID", tenantID)
		}
		next.ServeHTTP(w, r)
	})
}
