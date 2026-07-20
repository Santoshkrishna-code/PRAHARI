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
	exportApp "prahari/services/moc/internal/application/export"
	implApp "prahari/services/moc/internal/application/implementation"
	reportingApp "prahari/services/moc/internal/application/reporting"
	reqApp "prahari/services/moc/internal/application/request"
	reviewApp "prahari/services/moc/internal/application/review"
	searchApp "prahari/services/moc/internal/application/search"
	verifApp "prahari/services/moc/internal/application/verification"

	// Infrastructure adapters
	kafkaInfra "prahari/services/moc/internal/infrastructure/kafka"
	maintInfra "prahari/services/moc/internal/infrastructure/maintenance"
	pgInfra "prahari/services/moc/internal/infrastructure/postgres"
	redisInfra "prahari/services/moc/internal/infrastructure/redis"
	riskInfra "prahari/services/moc/internal/infrastructure/risk"

	// Interfaces
	eventsIface "prahari/services/moc/internal/interfaces/events"
	httpIface "prahari/services/moc/internal/interfaces/http"
)

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "moc-service", cfg.Environment)
	if err == nil && tp != nil {
		defer func() { _ = tp.Shutdown(ctx) }()
	}

	db, err := InitDatabase(ctx, cfg.DatabaseURL)
	if err != nil {
		prahariLogger.Error(ctx, "Database initialization failed", prahariLogger.Err(err))
	}
	if db != nil {
		defer db.Close()
	}

	redisClient, err := InitRedis(ctx, cfg.RedisAddr)
	if err == nil && redisClient != nil {
		defer redisClient.Close()
	}

	_ = InitKafka(ctx, cfg.KafkaBrokers)

	store := pgInfra.NewStore(db)
	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher()

	riskClient := riskInfra.NewClient(cfg.RiskGrpcAddr)
	maintClient := maintInfra.NewClient(cfg.MaintenanceGrpcAddr)

	// DI Wiring
	reqSvc := reqApp.NewService(store, kafkaPublisher)
	reviewSvc := reviewApp.NewService(store, kafkaPublisher, riskClient)
	implSvc := implApp.NewService(store, kafkaPublisher, maintClient)
	verifSvc := verifApp.NewService(store, kafkaPublisher)
	reportingSvc := reportingApp.NewService(store)
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(reqSvc, reviewSvc, implSvc, verifSvc, reportingSvc, searchSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Management of Change (MOC) Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down MOC Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "MOC Service exited gracefully.")
}
