package study

import (
	"context"
	"fmt"
	"time"

	"prahari/services/pha/internal/domain/events"
	"prahari/services/pha/internal/domain/phastudy"
	"prahari/services/pha/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveStudy(ctx context.Context, s *phastudy.Study) error
	GetStudyByID(ctx context.Context, id string) (*phastudy.Study, error)
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{
		repo:      repo,
		publisher: pub,
	}
}

func (s *Service) CreateStudy(ctx context.Context, study *phastudy.Study) error {
	study.ID = fmt.Sprintf("pha-%d", time.Now().UnixNano())
	study.StudyNumber = fmt.Sprintf("PHA-%s-%d", study.PlantID, time.Now().Unix()%100000)
	study.Status = string(status.CodeDraft)
	study.CreatedAt = time.Now()
	study.UpdatedAt = time.Now()

	// Default 5-year revalidation due date
	revalDate := time.Now().AddDate(5, 0, 0)
	study.RevalidationDueAt = &revalDate

	if err := s.repo.SaveStudy(ctx, study); err != nil {
		return fmt.Errorf("failed to save PHA study: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventPHACreated, study)
	prahariLogger.Info(ctx, "PHA study created", prahariLogger.String("study_number", study.StudyNumber))
	return nil
}

func (s *Service) ApproveStudy(ctx context.Context, studyID string) error {
	st, err := s.repo.GetStudyByID(ctx, studyID)
	if err != nil {
		return err
	}
	if err := status.ValidateTransition(status.Code(st.Status), status.CodeApproval); err != nil {
		return err
	}
	st.Status = string(status.CodeApproval)
	st.UpdatedAt = time.Now()

	if err := s.repo.SaveStudy(ctx, st); err != nil {
		return err
	}
	_ = s.publisher.Publish(ctx, events.EventPHAApproved, st)
	prahariLogger.Info(ctx, "PHA study approved", prahariLogger.String("study_id", studyID))
	return nil
}
