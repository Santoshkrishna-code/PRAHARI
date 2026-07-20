package reporting_test

import (
	"context"
	"testing"

	"prahari/services/analytics/internal/application/reporting"
	"prahari/services/analytics/internal/domain/report"
)

type mockRepo struct {
	savedReport *report.Report
}

func (m *mockRepo) SaveReport(ctx context.Context, r *report.Report) error {
	m.savedReport = r
	return nil
}

func (m *mockRepo) GetReportByID(ctx context.Context, id string) (*report.Report, error) {
	return m.savedReport, nil
}

type mockPublisher struct{}

func (m *mockPublisher) Publish(ctx context.Context, eventType string, payload any) error {
	return nil
}

func TestGenerateExecutiveReport(t *testing.T) {
	repo := &mockRepo{}
	svc := reporting.NewService(repo, &mockPublisher{})

	rep, err := svc.GenerateExecutiveReport(context.Background(), "P01", "Q3 ESG Performance", "QUARTERLY", "u-123")
	if err != nil {
		t.Fatalf("unexpected error during report generation: %v", err)
	}

	if rep.ID == "" {
		t.Error("expected generated report ID to be non-empty")
	}

	if rep.FileURL == "" {
		t.Error("expected generated PDF S3 FileURL link to be populated")
	}
}
