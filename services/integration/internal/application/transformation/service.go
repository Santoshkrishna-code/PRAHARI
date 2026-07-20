package transformation

import (
	"context"
	"encoding/json"
	"fmt"

	"prahari/services/integration/internal/domain/events"
	"prahari/services/integration/internal/domain/mapping"
)

type Repository interface {
	GetMappingRules(ctx context.Context, connectorID string) ([]*mapping.FieldMap, error)
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

func (s *Service) TransformPayload(ctx context.Context, connectorID string, raw []byte) ([]byte, error) {
	rules, err := s.repo.GetMappingRules(ctx, connectorID)
	if err != nil {
		return nil, err
	}

	var data map[string]any
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, fmt.Errorf("failed to parse raw payload: %w", err)
	}

	result := make(map[string]any)
	for _, rule := range rules {
		if val, ok := data[rule.ExternalKey]; ok {
			result[rule.InternalKey] = val
		}
	}

	// fallback if no rules mapped
	if len(result) == 0 {
		result = data
	}

	out, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	_ = s.publisher.Publish(ctx, events.EventMessageTransformed, result)
	return out, nil
}
