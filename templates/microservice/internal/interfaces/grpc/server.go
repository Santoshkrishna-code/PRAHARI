package grpc

import (
	"context"
	"errors"

	"prahari/templates/microservice/internal/application"
)

// Server implements gRPC service endpoints contracts.
type Server struct {
	svc *application.IncidentService
}

// NewServer constructs a Server.
func NewServer(svc *application.IncidentService) *Server {
	return &Server{svc: svc}
}

// CreateIncident provides gRPC incident creation.
func (s *Server) CreateIncident(ctx context.Context, title string) (string, error) {
	if title == "" {
		return "", errors.New("title cannot be empty")
	}

	inc, err := s.svc.CreateIncident(ctx, title)
	if err != nil {
		return "", err
	}

	return inc.ID, nil
}
