package router

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"prahari/services/auth-service/internal/api/v1/handlers"
	"prahari/services/auth-service/internal/domain"
)

// SetupRouter binds handlers and standard middleware wrappers.
func SetupRouter(
	useCase domain.AuthUseCase,
	cognitoClient *cognitoidentityprovider.Client,
	logMiddleware func(http.Handler) http.Handler,
	corsMiddleware func(http.Handler) http.Handler,
	recoveryMiddleware func(http.Handler) http.Handler,
	rateLimitMiddleware func(http.Handler) http.Handler,
) *chi.Mux {
	r := chi.NewRouter()

	// 1. Basic Ingress Middlewares
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(recoveryMiddleware)
	r.Use(corsMiddleware)
	r.Use(logMiddleware)
	r.Use(rateLimitMiddleware)

	// 2. Health Monitoring Routes
	healthHandler := handlers.NewHealthHandler(cognitoClient)
	r.Get("/health/live", healthHandler.Live)
	r.Get("/health/ready", healthHandler.Ready)

	// 3. API V1 Routes
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Method(http.MethodPost, "/register", handlers.NewRegisterHandler(useCase))
		r.Method(http.MethodPost, "/login", handlers.NewLoginHandler(useCase))
		r.Method(http.MethodPost, "/refresh", handlers.NewRefreshHandler(useCase))
		r.Method(http.MethodGet, "/profile", handlers.NewProfileHandler(useCase))
	})

	return r
}
