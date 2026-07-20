package idempotency_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"prahari/shared/middleware/idempotency"
	prahariRedis "prahari/shared/redis"
)

func TestIdempotencyManager(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()

	client, _ := prahariRedis.NewClient(prahariRedis.Config{Address: mr.Addr()})
	defer client.Close()

	ctx := context.Background()
	m := idempotency.NewManager(client)
	key := "req-id-100"

	// 1. Claim lock (first time) -> expect success
	ok, err := m.AcquireLock(ctx, key, 10*time.Second)
	if err != nil {
		t.Fatalf("failed to acquire lock: %v", err)
	}
	if !ok {
		t.Error("expected lock to be acquired successfully")
	}

	// 2. Claim lock again -> expect failure (locked)
	ok, err = m.AcquireLock(ctx, key, 10*time.Second)
	if err != nil {
		t.Fatalf("failed second lock call: %v", err)
	}
	if ok {
		t.Error("expected lock acquisition to fail on duplicate request")
	}

	// 3. Save final response
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")
	err = m.SaveResponse(ctx, key, http.StatusCreated, headers, []byte(`{"status":"created"}`), 10*time.Second)
	if err != nil {
		t.Fatalf("failed to save response: %v", err)
	}

	// 4. Retrieve cached response
	resp, err := m.GetResponse(ctx, key)
	if err != nil {
		t.Fatalf("failed to get response: %v", err)
	}

	if resp == nil {
		t.Fatal("expected cached response, got nil")
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", resp.StatusCode)
	}

	if string(resp.Body) != `{"status":"created"}` {
		t.Errorf("expected JSON body, got %s", string(resp.Body))
	}
}
