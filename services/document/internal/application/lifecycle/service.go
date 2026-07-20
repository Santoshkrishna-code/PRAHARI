package lifecycle

import (
	"context"
	"time"

	"prahari/services/document/internal/domain/document"
	"prahari/services/document/internal/domain/events"
	"prahari/services/document/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)


type Repository interface {
	GetDocumentByID(ctx context.Context, id string) (*document.Document, error)
	SaveDocument(ctx context.Context, doc *document.Document) error
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

func (s *Service) PublishDocument(ctx context.Context, id string) error {
	doc, err := s.repo.GetDocumentByID(ctx, id)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(doc.Status), status.CodePublished); err != nil {
		return err
	}

	doc.Status = string(status.CodePublished)
	doc.UpdatedAt = time.Now()

	if err := s.repo.SaveDocument(ctx, doc); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventDocumentPublished, doc)
	prahariLogger.Info(ctx, "Controlled Document officially published", prahariLogger.String("doc_number", doc.DocumentNumber))
	return nil
}

func (s *Service) ArchiveDocument(ctx context.Context, id string) error {
	doc, err := s.repo.GetDocumentByID(ctx, id)
	if err != nil {
		return err
	}

	doc.Status = string(status.CodeArchived)
	doc.UpdatedAt = time.Now()

	if err := s.repo.SaveDocument(ctx, doc); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventDocumentArchived, doc)
	prahariLogger.Info(ctx, "Controlled Document archived", prahariLogger.String("id", id))
	return nil
}
