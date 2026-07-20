package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"go.uber.org/zap"

	"prahari/services/auth-service/internal/api/v1/router"
	"prahari/services/auth-service/internal/config"
	"prahari/services/auth-service/internal/events"
	"prahari/services/auth-service/internal/middleware"
	"prahari/services/auth-service/internal/repository"
	"prahari/services/auth-service/internal/service"
	"prahari/services/auth-service/internal/utils"
	"prahari/services/auth-service/pkg/logger"
)

// AppContainer aggregates all components required to run the Auth Service.
type AppContainer struct {
	Config        *config.Config
	Logger        *zap.Logger
	CognitoClient *cognitoidentityprovider.Client
	Router        http.Handler
}

// BuildApplication sets up dependencies and returns the application container.
func BuildApplication(ctx context.Context) (*AppContainer, error) {
	// 1. Load Configurations
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("config build error: %w", err)
	}

	// 2. Initialize Logger
	log, err := logger.InitLogger(cfg.Env, cfg.Logging.Level)
	if err != nil {
		return nil, fmt.Errorf("logger init error: %w", err)
	}

	// 3. Initialize Cognito Client
	cognitoClient, err := InitCognitoClient(ctx, &cfg.AWS)
	if err != nil {
		return nil, fmt.Errorf("cognito client init error: %w", err)
	}

	// 4. Initialize Repositories (Cognito + Cache)
	userRepo := repository.NewCognitoUserRepository(cognitoClient, cfg.Cognito.UserPoolID, cfg.Cognito.ClientID)
	userCache := repository.NewInMemoryUserCache(10 * time.Minute) // 10m profile TTL cache

	// 5. Initialize Outbound Event Publisher
	eventPublisher := events.NewLogEventPublisher(log)

	// 6. Initialize UseCase / Services Layer
	authUseCase := service.NewAuthUseCase(userRepo, userCache, eventPublisher, cfg.Env)

	// 7. Initialize Token Verifier (For Cognito JWT signature validation checks)
	issuerURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s", cfg.AWS.Region, cfg.Cognito.UserPoolID)
	jwksURL := cfg.Cognito.GetJWKSURL(cfg.AWS.Region)
	_ = utils.NewTokenVerifier(jwksURL, issuerURL) // Instantiated for JWT signature validations helper

	// 8. Compile Routing & Middleware chains
	logMiddleware := middleware.StructuredLogging(log)
	corsMiddleware := middleware.ConfigureCORS(cfg.Security.AllowedOrigins)
	recoveryMiddleware := middleware.PanicRecovery(log)
	
	rateLimiter := middleware.NewIPRateLimiter(cfg.Security.RateLimitRPS, cfg.Security.RateLimitRPS*2)
	rateLimitMiddleware := rateLimiter.Limit

	appRouter := router.SetupRouter(
		authUseCase,
		cognitoClient,
		logMiddleware,
		corsMiddleware,
		recoveryMiddleware,
		rateLimitMiddleware,
	)

	// Apply global security headers middleware
	finalHandler := middleware.SecurityHeaders(appRouter)

	return &AppContainer{
		Config:        cfg,
		Logger:        log,
		CognitoClient: cognitoClient,
		Router:        finalHandler,
	}, nil
}
