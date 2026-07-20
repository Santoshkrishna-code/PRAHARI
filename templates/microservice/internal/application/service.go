package application

import (
	"context"
	"fmt"
	"time"

	prahariLogger "prahari/shared/logger"
	"prahari/templates/microservice/internal/domain"
)

// IncidentService manages Incident aggregate lifecycles.
type IncidentService struct {
}

// NewIncidentService constructs an IncidentService.
func NewIncidentService() *IncidentService {
	return &IncidentService{}
}

// CreateIncident validates and registers a new aggregate.
func (s *IncidentService) CreateIncident(ctx context.Context, title string) (*domain.Incident, error) {
	incident := &domain.Incident{
		ID:        fmt.Sprintf("inc-%d", time.Now().UnixNano()),
		Title:     title,
		Status:    "OPEN",
		CreatedAt: time.Now(),
	}

	if err := incident.Validate(); err != nil {
		prahariLogger.Error(ctx, "failed incident validation check", prahariLogger.Err(err))
		return nil, err
	}

	prahariLogger.Info(ctx, "Incident created successfully", prahariLogger.String("incident_id", incident.ID))
	return incident, nil
}
