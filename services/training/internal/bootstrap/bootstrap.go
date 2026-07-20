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
	trainingApp "prahari/services/training/internal/application/training"
	enrollmentApp "prahari/services/training/internal/application/enrollment"
	competencyApp "prahari/services/training/internal/application/competency"
	assessmentApp "prahari/services/training/internal/application/assessment"
	certificationApp "prahari/services/training/internal/application/certification"
	searchApp "prahari/services/training/internal/application/search"
	reportingApp "prahari/services/training/internal/application/reporting"
	exportApp "prahari/services/training/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/training/internal/infrastructure/postgres"
	redisInfra "prahari/services/training/internal/infrastructure/redis"
	kafkaInfra "prahari/services/training/internal/infrastructure/kafka"
	workflowInfra "prahari/services/training/internal/infrastructure/workflow"
	identityInfra "prahari/services/training/internal/infrastructure/identity"
	eventsIface "prahari/services/training/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/training/internal/interfaces/http"
)

// RunApp handles bootstrap coordination sequence.
func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "training-service", cfg.Environment)
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

	workflowClient := workflowInfra.NewClient(cfg.WorkflowGrpcAddr)
	_ = identityInfra.NewClient(cfg.IdentityGrpcAddr)

	trainingStore := pgInfra.NewTrainingStore(db)
	_ = pgInfra.NewProgramStore(db)
	_ = pgInfra.NewCourseStore(db)
	_ = pgInfra.NewMatrixStore(db)
	_ = pgInfra.NewSkillStore(db)
	_ = pgInfra.NewCertStore(db)
	_ = pgInfra.NewEnrollStore(db)
	_ = pgInfra.NewAttendanceStore(db)
	_ = pgInfra.NewAssessStore(db)
	_ = pgInfra.NewResultStore(db)
	_ = pgInfra.NewRenewalStore(db)
	_ = pgInfra.NewCommentStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	trailStore := pgInfra.NewTrailStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	trainingSvc := trainingApp.NewService(trainingStore, kafkaPublisher, trailStore, timelineStore, workflowClient)
	enrollmentSvc := enrollmentApp.NewService(nil)
	competencySvc := competencyApp.NewService(nil)
	assessmentSvc := assessmentApp.NewService(nil)
	certificationSvc := certificationApp.NewService(nil)
	searchSvc := searchApp.NewService(trainingStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(trainingStore)

	kafkaConsumer := kafkaInfra.NewConsumer(trainingSvc)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, trainingSvc, enrollmentSvc, competencySvc, assessmentSvc, certificationSvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Training Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Training Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Training Service exited gracefully.")
}
