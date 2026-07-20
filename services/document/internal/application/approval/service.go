package approval

import (
	"context"
	"fmt"
	"time"

	"prahari/services/document/internal/domain/document"
	"prahari/services/document/internal/domain/documentapproval"
	"prahari/services/document/internal/domain/events"
	"prahari/services/document/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetDocumentByID(ctx context.Context, id string) (*document.Document, error)
	SaveDocument(ctx context.Context, doc *document.Document) error
	SaveApproval(ctx context.Context, app *documentapproval.Approval) error
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

func (s *Service) ApproveDocument(ctx context.Context, id, approverID, comments string) error {
	doc, err := s.repo.GetDocumentByID(ctx, id)
	if err != nil {
		return err
	}

	app := &documentapproval.Approval{
		ID:         fmt.Sprintf("app-%d", time.Now().UnixNano()),
		DocumentID: id,
		VersionID:  doc.CurrentVersion,
		ApproverID: approverID,
		Approved:   true,
		ApprovedAt: time.Now(),
		Comments:   comments,
	}

	if err := s.repo.SaveApproval(ctx, app); err != nil {
		return err
	}

	doc.Status = string(status.CodeApproval)
	doc.UpdatedAt = time.Now()

	if err := s.repo.SaveDocument(ctx, doc); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventDocumentApproved, app)
	prahariLogger.Info(ctx, "Controlled Document sign-off approved", prahariLogger.String("doc_id", id))
	return nil
}
