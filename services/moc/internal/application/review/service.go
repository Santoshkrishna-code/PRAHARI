package review

import (
	"context"
	"fmt"
	"time"

	"prahari/services/moc/internal/domain/approval"
	"prahari/services/moc/internal/domain/changerequest"
	"prahari/services/moc/internal/domain/events"
	"prahari/services/moc/internal/domain/impactassessment"
	"prahari/services/moc/internal/domain/riskreview"
	"prahari/services/moc/internal/domain/safetyreview"
	"prahari/services/moc/internal/domain/status"
	"prahari/services/moc/internal/domain/technicalreview"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetRequestByID(ctx context.Context, id string) (*changerequest.Request, error)
	SaveRequest(ctx context.Context, req *changerequest.Request) error
	SaveImpactAssessment(ctx context.Context, ia *impactassessment.Assessment) error
	SaveTechnicalReview(ctx context.Context, tr *technicalreview.Review) error
	SaveRiskReview(ctx context.Context, rr *riskreview.Review) error
	SaveSafetyReview(ctx context.Context, sr *safetyreview.Review) error
	SaveApproval(ctx context.Context, app *approval.Record) error
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type RiskClient interface {
	CreateRiskAssessment(ctx context.Context, plantID, title, description string) (string, error)
}

type Service struct {
	repo       Repository
	publisher  EventPublisher
	riskClient RiskClient
}

func NewService(repo Repository, pub EventPublisher, riskClient RiskClient) *Service {
	return &Service{
		repo:       repo,
		publisher:  pub,
		riskClient: riskClient,
	}
}

func (s *Service) SubmitImpactAssessment(ctx context.Context, ia *impactassessment.Assessment) error {
	req, err := s.repo.GetRequestByID(ctx, ia.ChangeRequestID)
	if err != nil {
		return err
	}
	if err := status.ValidateTransition(status.Code(req.Status), status.CodeImpactAssessment); err != nil {
		return err
	}
	req.Status = string(status.CodeImpactAssessment)
	req.UpdatedAt = time.Now()

	ia.ID = fmt.Sprintf("ia-%d", time.Now().UnixNano())
	ia.AssessedAt = time.Now()

	if err := s.repo.SaveImpactAssessment(ctx, ia); err != nil {
		return err
	}
	_ = s.repo.SaveRequest(ctx, req)
	_ = s.publisher.Publish(ctx, events.EventMOCReviewStarted, ia)
	return nil
}

func (s *Service) SubmitTechnicalReview(ctx context.Context, tr *technicalreview.Review) error {
	tr.ID = fmt.Sprintf("tr-%d", time.Now().UnixNano())
	tr.ReviewedAt = time.Now()
	return s.repo.SaveTechnicalReview(ctx, tr)
}

func (s *Service) SubmitRiskReview(ctx context.Context, rr *riskreview.Review) error {
	rr.ID = fmt.Sprintf("rr-%d", time.Now().UnixNano())
	rr.ReviewedAt = time.Now()

	req, _ := s.repo.GetRequestByID(ctx, rr.ChangeRequestID)
	if req != nil && s.riskClient != nil && rr.RiskAssessmentID == "" {
		raID, err := s.riskClient.CreateRiskAssessment(ctx, req.PlantID, fmt.Sprintf("Risk Assessment for MOC: %s", req.Title), req.Description)
		if err == nil {
			rr.RiskAssessmentID = raID
		}
	}

	return s.repo.SaveRiskReview(ctx, rr)
}

func (s *Service) ApproveChange(ctx context.Context, app *approval.Record) error {
	req, err := s.repo.GetRequestByID(ctx, app.ChangeRequestID)
	if err != nil {
		return err
	}
	app.ID = fmt.Sprintf("app-%d", time.Now().UnixNano())
	app.ApprovedAt = time.Now()

	if app.Decision == "APPROVED" {
		if err := status.ValidateTransition(status.Code(req.Status), status.CodeApproval); err == nil {
			req.Status = string(status.CodeApproval)
			req.UpdatedAt = time.Now()
			_ = s.repo.SaveRequest(ctx, req)
			_ = s.publisher.Publish(ctx, events.EventMOCApproved, app)
		}
	} else {
		req.Status = string(status.CodeRejected)
		req.UpdatedAt = time.Now()
		_ = s.repo.SaveRequest(ctx, req)
	}

	prahariLogger.Info(ctx, "MOC approval decision registered", prahariLogger.String("decision", app.Decision))
	return s.repo.SaveApproval(ctx, app)
}
