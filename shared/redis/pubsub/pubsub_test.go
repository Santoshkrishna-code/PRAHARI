package pubsub_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	prahariRedis "prahari/shared/redis"
	"prahari/shared/redis/pubsub"
)

func TestPubSub_PublishSubscribe(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("failed to start miniredis: %v", err)
	}
	defer mr.Close()

	client, err := prahariRedis.NewClient(prahariRedis.Config{
		Address: mr.Addr(),
	})
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer client.Close()

	pub := pubsub.NewPublisher(client)
	sub := pubsub.NewSubscriber(client)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	channel := "test-events-channel"
	messagePayload := map[string]string{"event": "incident_created", "id": "inc-456"}

	var wg sync.WaitGroup
	wg.Add(1)

	var receivedPayload string

	// Start subscriber loop in a goroutine
	go func() {
		_ = sub.Subscribe(ctx, channel, func(c context.Context, payload []byte) error {
			receivedPayload = string(payload)
			wg.Done()
			return nil
		})
	}()

	// Small pause to let the subscription bind
	time.Sleep(100 * time.Millisecond)

	// Publish message
	err = pub.Publish(context.Background(), channel, messagePayload)
	if err != nil {
		t.Fatalf("failed to publish: %v", err)
	}

	// Wait for message delivery
	wg.Wait()

	if !stringsContains(receivedPayload, "incident_created") {
		t.Errorf("expected payload to contain 'incident_created', got '%s'", receivedPayload)
	}
}

func stringsContains(s, substr string) bool {
	// Simple inline check
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
