package grpc

import (
	"context"
	"time"

	playbackApp "prahari/services/digitaltwin/internal/application/playback"
	searchApp "prahari/services/digitaltwin/internal/application/search"
	simulationApp "prahari/services/digitaltwin/internal/application/simulation"
	syncApp "prahari/services/digitaltwin/internal/application/synchronization"
	"prahari/services/digitaltwin/internal/domain/playback"
	"prahari/services/digitaltwin/internal/domain/search"
	"prahari/services/digitaltwin/internal/domain/simulation"
	"prahari/services/digitaltwin/internal/domain/twin"
)

type Server struct {
	syncSvc       *syncApp.Service
	simulationSvc *simulationApp.Service
	playbackSvc   *playbackApp.Service
	searchSvc     *searchApp.Service
}

func NewServer(
	syncSvc *syncApp.Service,
	simulationSvc *simulationApp.Service,
	playbackSvc *playbackApp.Service,
	searchSvc *searchApp.Service,
) *Server {
	return &Server{
		syncSvc:       syncSvc,
		simulationSvc: simulationSvc,
		playbackSvc:   playbackSvc,
		searchSvc:     searchSvc,
	}
}

func (s *Server) CreateTwin(ctx context.Context, name, plantID string) (*twin.DigitalTwin, error) {
	return &twin.DigitalTwin{
		ID:        "twin-grpc-01",
		PlantID:   plantID,
		Name:      name,
		Status:    "DRAFT",
		Version:   1,
		CreatedAt: time.Now(),
	}, nil
}

func (s *Server) UpdateTwinState(ctx context.Context, twinID, equipmentID string, val float64, quality string) error {
	return s.syncSvc.SyncTelemetry(ctx, twinID, equipmentID, val, quality)
}

func (s *Server) RunSimulation(ctx context.Context, twinID, name, params string) (*simulation.Scenario, error) {
	return s.simulationSvc.RunScenario(ctx, twinID, name, params)
}

func (s *Server) ReplayTimeline(ctx context.Context, twinID string, start, end time.Time, speed float64) (*playback.Session, error) {
	return s.playbackSvc.StartPlayback(ctx, twinID, start, end, speed)
}

func (s *Server) SearchTwin(ctx context.Context, criteria *search.Criteria) ([]*twin.DigitalTwin, int64, error) {
	return s.searchSvc.ExecuteSearch(ctx, criteria)
}
