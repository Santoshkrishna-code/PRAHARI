package chat

import (
	"context"
	"fmt"
	"time"

	"prahari/services/ai/internal/domain/conversation"
	"prahari/services/ai/internal/domain/generation"
	"prahari/services/ai/internal/domain/policy"
)

type ConversationRepository interface {
	SaveThread(ctx context.Context, t *conversation.Thread) error
	SaveMessage(ctx context.Context, m *conversation.Message) error
	GetMessagesByThread(ctx context.Context, threadID string) ([]*conversation.Message, error)
}

type LLMClient interface {
	Generate(ctx context.Context, prompt string, context []string) (string, error)
}

type Service struct {
	repo ConversationRepository
	llm  LLMClient
}

func NewService(repo ConversationRepository, llm LLMClient) *Service {
	return &Service{repo: repo, llm: llm}
}

func (s *Service) CreateThread(ctx context.Context, userID, plantID string) (*conversation.Thread, error) {
	t := &conversation.Thread{
		ID:        fmt.Sprintf("th-%d", time.Now().UnixNano()),
		UserID:    userID,
		PlantID:   plantID,
		CreatedAt: time.Now(),
	}
	if err := s.repo.SaveThread(ctx, t); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) SubmitMessage(ctx context.Context, threadID, role, content string) (*generation.Response, error) {
	// Guardrail check
	if role == "USER" && !policy.EvaluateQuerySafety(content) {
		return &generation.Response{
			Answer:     "I cannot fulfill this request as it violates safety guidelines.",
			Confidence: 0.0,
		}, nil
	}

	cleaned := policy.RedactPII(content)
	msg := &conversation.Message{
		ID:        fmt.Sprintf("msg-%d", time.Now().UnixNano()),
		ThreadID:  threadID,
		Role:      role,
		Content:   cleaned,
		CreatedAt: time.Now(),
	}

	if err := s.repo.SaveMessage(ctx, msg); err != nil {
		return nil, err
	}

	// LLM mock generation
	answer, err := s.llm.Generate(ctx, cleaned, nil)
	if err != nil {
		return nil, err
	}

	assistantMsg := &conversation.Message{
		ID:        fmt.Sprintf("msg-%d", time.Now().UnixNano()),
		ThreadID:  threadID,
		Role:      "ASSISTANT",
		Content:   answer,
		CreatedAt: time.Now(),
	}

	_ = s.repo.SaveMessage(ctx, assistantMsg)

	return &generation.Response{
		Answer:     answer,
		Confidence: 0.95,
	}, nil
}
