package grpc

import (
	"context"
	"errors"

	trainingApp "prahari/services/training/internal/application/training"
	searchApp "prahari/services/training/internal/application/search"
	trainingDomain "prahari/services/training/internal/domain/training"
	searchDomain "prahari/services/training/internal/domain/search"
)

// Server exposes gRPC endpoints.
type Server struct {
	training *trainingApp.Service
	search   *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	training *trainingApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		training: training,
		search:   search,
	}
}

// CreateTraining registers new entry.
func (s *Server) CreateTraining(ctx context.Context, cmd trainingApp.CreateTrainingCommand) (*trainingDomain.Training, error) {
	return s.training.CreateTraining(ctx, cmd, "grpc-actor")
}

// GetTraining returns registers details.
func (s *Server) GetTraining(ctx context.Context, id string) (*trainingDomain.Training, error) {
	if id == "" {
		return nil, errors.New("training ID is required")
	}
	return s.training.GetTraining(ctx, id)
}

// SearchTraining queries matches.
func (s *Server) SearchTraining(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
