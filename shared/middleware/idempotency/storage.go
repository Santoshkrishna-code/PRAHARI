package idempotency

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	redisSDK "github.com/redis/go-redis/v9"
	prahariRedis "prahari/shared/redis"
)

// Response maps cached HTTP headers and status code details.
type Response struct {
	StatusCode int                 `json:"status_code"`
	Headers    map[string][]string `json:"headers"`
	Body       []byte              `json:"body"`
}

// Manager coordinates idempotency checks inside Redis.
type Manager struct {
	client *prahariRedis.Client
}

// NewManager constructs a new Manager instance.
func NewManager(client *prahariRedis.Client) *Manager {
	return &Manager{client: client}
}

// AcquireLock attempts to secure the idempotency key using SETNX.
// Returns true if the key was successfully claimed (this is the first execution).
func (m *Manager) AcquireLock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	if m.client == nil || m.client.UniversalClient == nil {
		return false, errors.New("redis client is uninitialized")
	}

	redisKey := fmt.Sprintf("idempotency:lock:%s", key)
	ok, err := m.client.UniversalClient.SetNX(ctx, redisKey, "processing", ttl).Result()
	if err != nil {
		return false, fmt.Errorf("failed to claim idempotency lock: %w", err)
	}

	return ok, nil
}

// GetResponse retrieves the cached response matching the key.
// Returns nil if the request is still processing or does not exist.
func (m *Manager) GetResponse(ctx context.Context, key string) (*Response, error) {
	if m.client == nil || m.client.UniversalClient == nil {
		return nil, errors.New("redis client is uninitialized")
	}

	redisKey := fmt.Sprintf("idempotency:resp:%s", key)
	dataBytes, err := m.client.UniversalClient.Get(ctx, redisKey).Bytes()
	if err != nil {
		if errors.Is(err, redisSDK.Nil) {
			return nil, nil // Not cached yet
		}
		return nil, fmt.Errorf("failed to fetch cached response: %w", err)
	}

	var resp Response
	err = json.Unmarshal(dataBytes, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize cached response: %w", err)
	}

	return &resp, nil
}

// SaveResponse caches the finished HTTP response payload.
func (m *Manager) SaveResponse(ctx context.Context, key string, statusCode int, headers http.Header, body []byte, ttl time.Duration) error {
	if m.client == nil || m.client.UniversalClient == nil {
		return errors.New("redis client is uninitialized")
	}

	resp := Response{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}

	dataBytes, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("failed to marshal response payload: %w", err)
	}

	redisKey := fmt.Sprintf("idempotency:resp:%s", key)
	err = m.client.UniversalClient.Set(ctx, redisKey, dataBytes, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to save response inside Redis: %w", err)
	}

	// Delete temporary lock key
	lockKey := fmt.Sprintf("idempotency:lock:%s", key)
	_ = m.client.UniversalClient.Del(ctx, lockKey).Err()

	return nil
}
