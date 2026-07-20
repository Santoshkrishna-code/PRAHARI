package versioning

import (
	"context"
	"fmt"
	"time"

	"prahari/services/document/internal/domain/document"
	"prahari/services/document/internal/domain/documentversion"
	"prahari/services/document/internal/domain/events"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetDocumentByID(ctx context.Context, id string) (*document.Document, error)
	SaveDocument(ctx context.Context, doc *document.Document) error
	SaveVersion(ctx context.Context, ver *documentversion.Version) error
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

func (s *Service) Checkout(ctx context.Context, id, userID string) error {
	doc, err := s.repo.GetDocumentByID(ctx, id)
	if err != nil {
		return err
	}

	if doc.CheckedOutBy != "" {
		return fmt.Errorf("document already checked out by %s", doc.CheckedOutBy)
	}

	now := time.Now()
	doc.CheckedOutBy = userID
	doc.CheckedOutAt = &now
	doc.UpdatedAt = now

	if err := s.repo.SaveDocument(ctx, doc); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Document checked out for editing", prahariLogger.String("doc_id", id), prahariLogger.String("user_id", userID))
	return nil
}

func (s *Service) Checkin(ctx context.Context, id, userID, fileURL, fileHash, changeSummary string) (*documentversion.Version, error) {
	doc, err := s.repo.GetDocumentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if doc.CheckedOutBy != userID {
		return nil, fmt.Errorf("user %s does not hold checkout lock on document %s", userID, id)
	}

	ver := &documentversion.Version{
		ID:            fmt.Sprintf("ver-%d", time.Now().UnixNano()),
		DocumentID:    id,
		VersionNumber: fmt.Sprintf("%.1f", 1.0+float64(time.Now().Unix()%100)*0.1),
		FileURL:       fileURL,
		FileHash:      fileHash,
		ChangeSummary: changeSummary,
		CreatedBy:     userID,
		CreatedAt:     time.Now(),
	}

	doc.CheckedOutBy = ""
	doc.CheckedOutAt = nil
	doc.CurrentVersion = ver.VersionNumber
	doc.UpdatedAt = time.Now()

	if err := s.repo.SaveVersion(ctx, ver); err != nil {
		return nil, err
	}
	if err := s.repo.SaveDocument(ctx, doc); err != nil {
		return nil, err
	}

	_ = s.publisher.Publish(ctx, events.EventDocumentRevised, ver)
	prahariLogger.Info(ctx, "Document checked in with new revision version", prahariLogger.String("version", ver.VersionNumber))
	return ver, nil
}
