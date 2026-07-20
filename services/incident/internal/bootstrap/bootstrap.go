package bootstrap

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	prahariLogger "prahari/shared/logger"

	// Application services
	incidentApp "prahari/services/incident/internal/application/incident"
	assignmentApp "prahari/services/incident/internal/application/assignment"
	investigationApp "prahari/services/incident/internal/application/investigation"
	searchApp "prahari/services/incident/internal/application/search"
	reportingApp "prahari/services/incident/internal/application/reporting"
	exportApp "prahari/services/incident/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/incident/internal/infrastructure/postgres"
	redisInfra "prahari/services/incident/internal/infrastructure/redis"
	kafkaInfra "prahari/services/incident/internal/infrastructure/kafka"
	workflowInfra "prahari/services/incident/internal/infrastructure/workflow"
	identityInfra "prahari/services/incident/internal/infrastructure/identity"
	eventsIface "prahari/services/incident/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/incident/internal/interfaces/http"
)

// RunApp is the master startup coordinator. It wires all dependencies using
// constructor injection and starts the HTTP server with graceful shutdown.
func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ──────────────────────────────────────────────
	// 1. Configuration
	// ──────────────────────────────────────────────
	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	// ──────────────────────────────────────────────
	// 2. Logger
	// ──────────────────────────────────────────────
	_ = InitLogger(cfg.Environment)

	// ──────────────────────────────────────────────
	// 3. Telemetry (OpenTelemetry)
	// ──────────────────────────────────────────────
	tp, err := InitTelemetry(ctx, "incident-service", cfg.Environment)
	if err == nil && tp != nil {
		defer func() { _ = tp.Shutdown(ctx) }()
	}

	// ──────────────────────────────────────────────
	// 4. PostgreSQL
	// ──────────────────────────────────────────────
	db, err := InitDatabase(ctx, cfg.DatabaseURL)
	if err != nil {
		prahariLogger.Error(ctx, "Database initialization failed", prahariLogger.Err(err))
	}
	if db != nil {
		defer db.Close()
	}

	// ──────────────────────────────────────────────
	// 5. Redis
	// ──────────────────────────────────────────────
	redisClient, err := InitRedis(ctx, cfg.RedisAddr)
	if err == nil && redisClient != nil {
		defer redisClient.Close()
	}

	// ──────────────────────────────────────────────
	// 6. Kafka
	// ──────────────────────────────────────────────
	_ = InitKafka(ctx, cfg.KafkaBrokers)

	// ──────────────────────────────────────────────
	// 7. Platform service clients
	// ──────────────────────────────────────────────
	workflowClient := workflowInfra.NewClient(cfg.WorkflowGrpcAddr)
	identityClient := identityInfra.NewClient(cfg.IdentityGrpcAddr)

	// ──────────────────────────────────────────────
	// 8. Repository / Store adapters
	// ──────────────────────────────────────────────
	incidentStore := pgInfra.NewIncidentStore(db)
	assignmentStore := pgInfra.NewAssignmentStore(db)
	investigationStore := pgInfra.NewInvestigationStore(db)
	rootCauseStore := pgInfra.NewRootCauseStore(db)
	auditStore := pgInfra.NewAuditStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	searchStore := pgInfra.NewSearchStore(db)

	// Redis cache layer
	_ = redisInfra.NewCache(redisClient)

	// Kafka event publisher
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// ──────────────────────────────────────────────
	// 9. Application services (dependency injection)
	// ──────────────────────────────────────────────
	incidentSvc := incidentApp.NewService(incidentStore, kafkaPublisher, auditStore, timelineStore, workflowClient)
	assignmentSvc := assignmentApp.NewService(assignmentStore, incidentStore, identityClient, kafkaPublisher, auditStore, timelineStore)
	investigationSvc := investigationApp.NewService(investigationStore, rootCauseStore, kafkaPublisher, auditStore, timelineStore)
	searchSvc := searchApp.NewService(searchStore)
	reportingSvc := reportingApp.NewService(nil) // Reporting counter adapter
	exportSvc := exportApp.NewService(searchStore)

	// ──────────────────────────────────────────────
	// 10. Kafka consumer (inbound events from Workflow Engine)
	// ──────────────────────────────────────────────
	kafkaConsumer := kafkaInfra.NewConsumer(nil)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	// ──────────────────────────────────────────────
	// 11. HTTP route registration
	// ──────────────────────────────────────────────
	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, incidentSvc, assignmentSvc, investigationSvc, searchSvc, reportingSvc, exportSvc)

	// ──────────────────────────────────────────────
	// 12. Middleware pipeline
	// ──────────────────────────────────────────────
	pipelineHandler := InitRouter(mux)

	// ──────────────────────────────────────────────
	// 13. HTTP server
	// ──────────────────────────────────────────────
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Incident Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	// ──────────────────────────────────────────────
	// 14. Graceful shutdown
	// ──────────────────────────────────────────────
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Incident Management Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)

	if err := srv.Shutdown(shutdownCtx); err != nil {
		prahariLogger.Error(ctx, "HTTP server shutdown failed", prahariLogger.Err(err))
	}

	prahariLogger.Info(ctx, "Incident Management Service exited gracefully.")
}
