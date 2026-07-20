package riskassessment

import (
	"context"
	"time"

	"github.com/google/uuid"

	prahariLogger "prahari/shared/logger"

	riskDomain "prahari/services/permit/internal/domain/riskassessment"
	hazardDomain "prahari/services/permit/internal/domain/hazard"
	isolationDomain "prahari/services/permit/internal/domain/isolation"
	gasDomain "prahari/services/permit/internal/domain/gastest"
	auditDomain "prahari/services/permit/internal/domain/audit"
	timelineDomain "prahari/services/permit/internal/domain/timeline"
)

// RiskRepository defines the persistence port for risk assessments.
type RiskRepository interface {
	Create(ctx context.Context, ra *riskDomain.RiskAssessment) error
	FindByPermitID(ctx context.Context, permitID string) (*riskDomain.RiskAssessment, error)
}

// HazardRepository defines the persistence port for hazards.
type HazardRepository interface {
	Create(ctx context.Context, h *hazardDomain.Hazard) error
	FindByPermitID(ctx context.Context, permitID string) ([]*hazardDomain.Hazard, error)
}

// IsolationRepository defines the persistence port for LOTO locks.
type IsolationRepository interface {
	Create(ctx context.Context, i *isolationDomain.Isolation) error
	FindByPermitID(ctx context.Context, permitID string) ([]*isolationDomain.Isolation, error)
	Update(ctx context.Context, i *isolationDomain.Isolation) error
}

// GasRepository defines the persistence port for gas tests.
type GasRepository interface {
	Create(ctx context.Context, gt *gasDomain.GasTest) error
	FindByPermitID(ctx context.Context, permitID string) ([]*gasDomain.GasTest, error)
}

// AuditLogger defines the port for logging change entries.
type AuditLogger interface {
	Log(ctx context.Context, entry *auditDomain.Entry) error
}

// TimelineRecorder defines the port for registering milestones.
type TimelineRecorder interface {
	Record(ctx context.Context, event *timelineDomain.Event) error
}

// Service manages pre-work safety evaluations.
type Service struct {
	riskRepo  RiskRepository
	hazRepo   HazardRepository
	isoRepo   IsolationRepository
	gasRepo   GasRepository
	audit     AuditLogger
	timeline  TimelineRecorder
}

// NewService instantiates a RiskAssessment Application Service.
func NewService(
	riskRepo RiskRepository,
	hazRepo HazardRepository,
	isoRepo IsolationRepository,
	gasRepo GasRepository,
	audit AuditLogger,
	timeline TimelineRecorder,
) *Service {
	return &Service{
		riskRepo:  riskRepo,
		hazRepo:   hazRepo,
		isoRepo:   isoRepo,
		gasRepo:   gasRepo,
		audit:     audit,
		timeline:  timeline,
	}
}

// PerformAssessment records a Likelihood x Consequence score card.
func (s *Service) PerformAssessment(ctx context.Context, ra *riskDomain.RiskAssessment) (*riskDomain.RiskAssessment, error) {
	ra.ID = uuid.New().String()
	ra.AssessedAt = time.Now()
	ra.CalculateRiskScore()

	if err := ra.Validate(); err != nil {
		return nil, err
	}

	if err := s.riskRepo.Create(ctx, ra); err != nil {
		return nil, err
	}

	// Timeline
	evt := timelineDomain.NewEvent(ra.PermitID, timelineDomain.EventRiskAssessed, ra.AssessorID, "Risk assessment completed", nil)
	_ = s.timeline.Record(ctx, evt)

	auditLog := auditDomain.NewEntry("risk_assessment", ra.ID, auditDomain.ActionRiskAssessed, ra.AssessorID, nil, ra)
	_ = s.audit.Log(ctx, auditLog)

	return ra, nil
}

// RecordHazard documents an identified safety hazard.
func (s *Service) RecordHazard(ctx context.Context, h *hazardDomain.Hazard) (*hazardDomain.Hazard, error) {
	h.ID = uuid.New().String()
	h.IdentifiedAt = time.Now()

	if err := h.Validate(); err != nil {
		return nil, err
	}

	if err := s.hazRepo.Create(ctx, h); err != nil {
		return nil, err
	}

	return h, nil
}

// RecordIsolation registers LOTO point verification.
func (s *Service) RecordIsolation(ctx context.Context, i *isolationDomain.Isolation) (*isolationDomain.Isolation, error) {
	i.ID = uuid.New().String()
	i.IsolatedAt = time.Now()
	i.Status = isolationDomain.StatusApplied

	if err := i.Validate(); err != nil {
		return nil, err
	}

	if err := s.isoRepo.Create(ctx, i); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(i.PermitID, timelineDomain.EventIsolationApplied, i.IsolatedBy, "Lock-out tag-out applied", nil)
	_ = s.timeline.Record(ctx, evt)

	return i, nil
}

// RecordGasTest registers gas reading.
func (s *Service) RecordGasTest(ctx context.Context, gt *gasDomain.GasTest) (*gasDomain.GasTest, error) {
	gt.ID = uuid.New().String()
	gt.TestedAt = time.Now()
	gt.EvaluateResult()

	if err := gt.Validate(); err != nil {
		return nil, err
	}

	if err := s.gasRepo.Create(ctx, gt); err != nil {
		return nil, err
	}

	evt := timelineDomain.NewEvent(gt.PermitID, timelineDomain.EventGasTestRecorded, gt.TestedBy, "Atmospheric gas test reading saved", nil)
	_ = s.timeline.Record(ctx, evt)

	return gt, nil
}
