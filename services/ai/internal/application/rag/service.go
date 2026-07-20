package rag

import (
	"context"
	"fmt"
	"time"

	"prahari/services/ai/internal/domain/document"
	"prahari/services/ai/internal/domain/events"
	"prahari/services/ai/internal/domain/retrieval"
)

type DocumentRepository interface {
	SaveDocument(ctx context.Context, d *document.Doc) error
	SaveChunk(ctx context.Context, c *document.Chunk) error
}

type VectorStore interface {
	IndexChunk(ctx context.Context, chunkID string, embedding []float32) error
	QuerySimilar(ctx context.Context, embedding []float32, limit int) ([]*retrieval.Result, error)
}

type EmbeddingClient interface {
	GenerateEmbedding(ctx context.Context, text string) ([]float32, error)
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	docRepo    DocumentRepository
	vectorSvc  VectorStore
	embedSvc   EmbeddingClient
	publisher  EventPublisher
}

func NewService(dr DocumentRepository, vs VectorStore, ec EmbeddingClient, pub EventPublisher) *Service {
	return &Service{
		docRepo:   dr,
		vectorSvc: vs,
		embedSvc:  ec,
		publisher: pub,
	}
}

func (s *Service) IndexDocument(ctx context.Context, sourceID, title, content string) error {
	doc := &document.Doc{
		ID:        fmt.Sprintf("doc-%d", time.Now().UnixNano()),
		SourceID:  sourceID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := s.docRepo.SaveDocument(ctx, doc); err != nil {
		return err
	}

	// Parsing chunk
	chunk := &document.Chunk{
		ID:         fmt.Sprintf("chk-%d", time.Now().UnixNano()),
		DocID:      doc.ID,
		Content:    content,
		PageNumber: 1,
	}

	if err := s.docRepo.SaveChunk(ctx, chunk); err != nil {
		return err
	}

	vector, err := s.embedSvc.GenerateEmbedding(ctx, chunk.Content)
	if err != nil {
		return err
	}

	if err := s.vectorSvc.IndexChunk(ctx, chunk.ID, vector); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventDocumentIndexed, doc)
	return nil
}

func (s *Service) SearchContext(ctx context.Context, query string, limit int) ([]*retrieval.Result, error) {
	vector, err := s.embedSvc.GenerateEmbedding(ctx, query)
	if err != nil {
		return nil, err
	}
	return s.vectorSvc.QuerySimilar(ctx, vector, limit)
}
