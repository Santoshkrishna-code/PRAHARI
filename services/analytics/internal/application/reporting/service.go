package reporting

import (
	"context"
	"fmt"
	"time"

	"prahari/services/analytics/internal/domain/events"
	"prahari/services/analytics/internal/domain/report"
)

type Repository interface {
	SaveReport(ctx context.Context, r *report.Report) error
	GetReportByID(ctx context.Context, id string) (*report.Report, error)
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{repo: repo, publisher: pub}
}

func (s *Service) GenerateExecutiveReport(ctx context.Context, plantID, title, reportType, userID string) (*report.Report, error) {
	r := &report.Report{
		ID:         fmt.Sprintf("rep-%d", time.Now().UnixNano()),
		PlantID:    plantID,
		Title:      title,
		ReportType: reportType,
		FileURL:    fmt.Sprintf("http://s3.aws/prahari-reports/%s.pdf", title),
		CreatedBy:  userID,
		CreatedAt:  time.Now(),
	}

	if err := s.repo.SaveReport(ctx, r); err != nil {
		return nil, err
	}

	_ = s.publisher.Publish(ctx, events.EventReportGenerated, r)
	return r, nil
}

func (s *Service) GetReport(ctx context.Context, id string) (*report.Report, error) {
	return s.repo.GetReportByID(ctx, id)
}
