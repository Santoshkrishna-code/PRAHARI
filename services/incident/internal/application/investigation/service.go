package investigation

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	investigationDomain "prahari/services/incident/internal/domain/investigation"
	rootcauseDomain "prahari/services/incident/internal/domain/rootcause"
	auditDomain "prahari/services/incident/internal/domain/audit"
	timelineDomain "prahari/services/incident/internal/domain/timeline"
)

// InvestigationRepository defines the persistence port for investigation operations.
type InvestigationRepository interface {
	Create(ctx context.Context, inv *investigationDomain.Investigation) error
	FindByID(ctx context.Context, id string) (*investigationDomain.Investigation, error)
	FindByIncidentID(ctx context.Context, incidentID string) (*investigationDomain.Investigation, error)
	Update(ctx context.Context, inv *investigationDomain.Investigation) error
}

// RootCauseRepository defines the persistence port for root cause operations.
type RootCauseRepository interface {
	Create(ctx context.Context, rc *rootcauseDomain.RootCause) error
	FindByInvestigationID(ctx context.Context, investigationID string) ([]*rootcauseDomain.RootCause, error)
}

// EventPublisher defines the port for publishing domain events.
type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

// AuditLogger defines the port for persisting audit entries.
type AuditLogger interface {
	Log(ctx context.Context, entry *auditDomain.Entry) error
}

// TimelineRecorder defines the port for recording timeline events.
type TimelineRecorder interface {
	Record(ctx context.Context, event *timelineDomain.Event) error
}

// Service orchestrates investigation and root cause analysis operations.
type Service struct {
	invRepo   InvestigationRepository
	rcRepo    RootCauseRepository
	events    EventPublisher
	audit     AuditLogger
	timeline  TimelineRecorder
}

// NewService constructs a Service with all required dependencies injected.
func NewService(
	invRepo InvestigationRepository,
	rcRepo RootCauseRepository,
	events EventPublisher,
	audit AuditLogger,
	timeline TimelineRecorder,
) *Service {
	return &Service{
		invRepo:  invRepo,
		rcRepo:   rcRepo,
		events:   events,
		audit:    audit,
		timeline: timeline,
	}
}

// StartInvestigation initiates a formal investigation for an incident.
func (s *Service) StartInvestigation(ctx context.Context, incidentID, investigatorID string, methodology investigationDomain.Methodology) (*investigationDomain.Investigation, error) {
	inv := &investigationDomain.Investigation{
		ID:             uuid.New().String(),
		IncidentID:     incidentID,
		InvestigatorID: investigatorID,
		Methodology:    methodology,
		Status:         investigationDomain.StatusInProgress,
		StartedAt:      time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := inv.Validate(); err != nil {
		return nil, fmt.Errorf("investigation validation failed: %w", err)
	}

	if err := s.invRepo.Create(ctx, inv); err != nil {
		return nil, fmt.Errorf("failed to persist investigation: %w", err)
	}

	// Timeline
	evt := timelineDomain.NewEvent(incidentID, timelineDomain.EventInvestigationStarted, investigatorID,
		fmt.Sprintf("Investigation started using %s methodology", methodology),
		map[string]string{"methodology": string(methodology)})
	if err := s.timeline.Record(ctx, evt); err != nil {
		prahariLogger.Error(ctx, "Failed to record timeline event", prahariLogger.Err(err))
	}

	// Audit
	entry := auditDomain.NewEntry("investigation", inv.ID, auditDomain.ActionInvestigated, investigatorID, nil, inv)
	if err := s.audit.Log(ctx, entry); err != nil {
		prahariLogger.Error(ctx, "Failed to write audit log", prahariLogger.Err(err))
	}

	// Publish event
	if err := s.events.Publish(ctx, "incident.investigation.started", inv); err != nil {
		prahariLogger.Error(ctx, "Failed to publish investigation event", prahariLogger.Err(err))
	}

	return inv, nil
}

// RecordFindings captures investigation findings and recommendations.
func (s *Service) RecordFindings(ctx context.Context, investigationID, findings, recommendations, actorID string) error {
	inv, err := s.invRepo.FindByID(ctx, investigationID)
	if err != nil {
		return fmt.Errorf("failed to retrieve investigation: %w", err)
	}

	inv.RecordFindings(findings, recommendations)

	if err := s.invRepo.Update(ctx, inv); err != nil {
		return fmt.Errorf("failed to persist findings: %w", err)
	}

	prahariLogger.Info(ctx, "Investigation findings recorded",
		prahariLogger.String("investigation_id", investigationID))

	return nil
}

// AddRootCause links a root cause analysis to an investigation.
func (s *Service) AddRootCause(ctx context.Context, rc *rootcauseDomain.RootCause, actorID string) error {
	rc.ID = uuid.New().String()
	rc.IdentifiedAt = time.Now()

	if err := rc.Validate(); err != nil {
		return fmt.Errorf("root cause validation failed: %w", err)
	}

	if err := s.rcRepo.Create(ctx, rc); err != nil {
		return fmt.Errorf("failed to persist root cause: %w", err)
	}

	// Timeline
	evt := timelineDomain.NewEvent(rc.IncidentID, timelineDomain.EventRootCauseIdentified, actorID,
		fmt.Sprintf("Root cause identified: %s (%s)", rc.Description, rc.Category), nil)
	if err := s.timeline.Record(ctx, evt); err != nil {
		prahariLogger.Error(ctx, "Failed to record timeline event", prahariLogger.Err(err))
	}

	return nil
}

// CompleteInvestigation marks an investigation as completed.
func (s *Service) CompleteInvestigation(ctx context.Context, investigationID, actorID string) error {
	inv, err := s.invRepo.FindByID(ctx, investigationID)
	if err != nil {
		return fmt.Errorf("failed to retrieve investigation: %w", err)
	}

	inv.Complete()

	if err := s.invRepo.Update(ctx, inv); err != nil {
		return fmt.Errorf("failed to persist investigation completion: %w", err)
	}

	prahariLogger.Info(ctx, "Investigation completed",
		prahariLogger.String("investigation_id", investigationID))

	return nil
}

// GetInvestigation retrieves the investigation for an incident.
func (s *Service) GetInvestigation(ctx context.Context, incidentID string) (*investigationDomain.Investigation, error) {
	return s.invRepo.FindByIncidentID(ctx, incidentID)
}

// GetRootCauses retrieves all root causes for an investigation.
func (s *Service) GetRootCauses(ctx context.Context, investigationID string) ([]*rootcauseDomain.RootCause, error) {
	return s.rcRepo.FindByInvestigationID(ctx, investigationID)
}
