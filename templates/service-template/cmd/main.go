package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"prahari/templates/service-template/internal/api"
	"prahari/templates/service-template/internal/config"
	"prahari/templates/service-template/internal/middleware"
	"prahari/templates/service-template/internal/repository"
	"prahari/templates/service-template/internal/service"
)

func main() {
	// 1. Load Configurations
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	
	fmt.Printf("Starting service in %s mode on port %s...\n", cfg.Env, cfg.Port)
	
	// 2. Initialize Database connection (Mocked DB handle for template representation)
	var db *sql.DB // db, err := sql.Open("postgres", cfg.DatabaseURL)
	
	// 3. Initialize Repositories and Services
	permitRepo := repository.NewPostgresPermitRepository(db)
	permitUseCase := service.NewPermitUseCase(permitRepo)
	
	// 4. Initialize HTTP Handlers & Middlewares
	httpHandler := api.NewHTTPHandler(permitUseCase)
	authMiddleware := middleware.NewAuthMiddleware(cfg.CognitoUserPool, cfg.CognitoClientID)
	
	// Apply middleware chain
	mux := http.NewServeMux()
	mux.Handle("/api/v1/permits", authMiddleware.Authenticate(httpHandler))
	mux.Handle("/api/v1/permits/", authMiddleware.Authenticate(httpHandler))
	
	// 5. Start Server
	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
