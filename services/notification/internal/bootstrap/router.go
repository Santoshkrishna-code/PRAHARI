package bootstrap

import (
	"net/http"
	"time"

	prahariMid "prahari/shared/middleware"
	prahariCORS "prahari/shared/middleware/cors"
	prahariLogMid "prahari/shared/middleware/logging"
	prahariRecovery "prahari/shared/middleware/recovery"
	prahariRequest "prahari/shared/middleware/requestid"
	prahariTimeout "prahari/shared/middleware/timeout"
)

// InitRouter registers standard execution filters.
func InitRouter(handler http.Handler) http.Handler {
	opts := prahariCORS.DefaultOptions()

	pipeline := prahariMid.New(
		prahariRequest.Middleware,
		prahariCORS.Middleware(opts),
		prahariLogMid.Middleware,
		prahariRecovery.Middleware,
		prahariTimeout.Middleware(15*time.Second),
	)

	return pipeline.Then(handler)
}
