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
	analyticsApp "prahari/services/visitor/internal/application/analytics"
	approvalApp "prahari/services/visitor/internal/application/approval"
	checkinApp "prahari/services/visitor/internal/application/checkin"
	checkoutApp "prahari/services/visitor/internal/application/checkout"
	exportApp "prahari/services/visitor/internal/application/export"
	musterApp "prahari/services/visitor/internal/application/muster"
	registrationApp "prahari/services/visitor/internal/application/registration"
	reportingApp "prahari/services/visitor/internal/application/reporting"
	searchApp "prahari/services/visitor/internal/application/search"

	// Infrastructure adapters
	kafkaInfra "prahari/services/visitor/internal/infrastructure/kafka"
	pgInfra "prahari/services/visitor/internal/infrastructure/postgres"
	redisInfra "prahari/services/visitor/internal/infrastructure/redis"

	// Interfaces
	eventsIface "prahari/services/visitor/internal/interfaces/events"
	httpIface "prahari/services/visitor/internal/interfaces/http"
)

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "visitor-service", cfg.Environment)
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

	// App Services DI
	registrationSvc := registrationApp.NewService(store, kafkaPublisher)
	approvalSvc := approvalApp.NewService(store, kafkaPublisher)
	checkinSvc := checkinApp.NewService(store, kafkaPublisher)
	checkoutSvc := checkoutApp.NewService(store, kafkaPublisher)
	musterSvc := musterApp.NewService(store, kafkaPublisher)
	reportingSvc := reportingApp.NewService(store)
	analyticsSvc := analyticsApp.NewService(store)
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(registrationSvc, approvalSvc, checkinSvc, checkoutSvc, musterSvc, reportingSvc, analyticsSvc, searchSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Visitor Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Visitor Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "Visitor Service exited gracefully.")
}
