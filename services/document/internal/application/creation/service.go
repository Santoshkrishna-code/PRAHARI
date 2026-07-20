package creation

import (
	"context"
	"fmt"
	"time"

	"prahari/services/document/internal/domain/document"
	"prahari/services/document/internal/domain/events"
	"prahari/services/document/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
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

func (s *Service) CreateDocument(ctx context.Context, doc *document.Document) error {
	doc.ID = fmt.Sprintf("doc-%d", time.Now().UnixNano())
	doc.DocumentNumber = fmt.Sprintf("DOC-%s-%d", doc.PlantID, time.Now().Unix()%100000)
	doc.CurrentVersion = "1.0"
	doc.Status = string(status.CodeDraft)
	doc.CreatedAt = time.Now()
	doc.UpdatedAt = time.Now()

	if err := s.repo.SaveDocument(ctx, doc); err != nil {
		return fmt.Errorf("failed to save controlled document: %w", err)
	}

	_ = s.publisher.Publish(ctx, events.EventDocumentCreated, doc)
	prahariLogger.Info(ctx, "Controlled Document created", prahariLogger.String("doc_number", doc.DocumentNumber))
	return nil
}
