package chat_test

import (
	"context"
	"testing"

	"prahari/services/ai/internal/application/chat"
	"prahari/services/ai/internal/domain/conversation"
)

type mockRepo struct {
	thread  *conversation.Thread
	message *conversation.Message
}

func (m *mockRepo) SaveThread(ctx context.Context, t *conversation.Thread) error {
	m.thread = t
	return nil
}

func (m *mockRepo) SaveMessage(ctx context.Context, msg *conversation.Message) error {
	m.message = msg
	return nil
}

func (m *mockRepo) GetMessagesByThread(ctx context.Context, threadID string) ([]*conversation.Message, error) {
	return []*conversation.Message{m.message}, nil
}

type mockLLM struct{}

func (m *mockLLM) Generate(ctx context.Context, prompt string, context []string) (string, error) {
	return "Mock LLM output response", nil
}

func TestChatSession(t *testing.T) {
	repo := &mockRepo{}
	llm := &mockLLM{}
	svc := chat.NewService(repo, llm)

	thread, err := svc.CreateThread(context.Background(), "u-1", "P01")
	if err != nil {
		t.Fatalf("unexpected thread creation error: %v", err)
	}

	if thread.ID == "" {
		t.Error("expected generated thread ID")
	}

	resp, err := svc.SubmitMessage(context.Background(), thread.ID, "USER", "Tell me about gas limits")
	if err != nil {
		t.Fatalf("unexpected message submission error: %v", err)
	}

	if resp.Answer != "Mock LLM output response" {
		t.Errorf("expected mock answer, got %s", resp.Answer)
	}
}
