package cache_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	prahariRedis "prahari/shared/redis"
	"prahari/shared/redis/cache"
	"prahari/shared/redis/serializer"
)

type Worker struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestCache_SetGetDelete(t *testing.T) {
	// 1. Boot miniredis mock server
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("failed to start miniredis: %v", err)
	}
	defer mr.Close()

	// 2. Initialize client
	client, err := prahariRedis.NewClient(prahariRedis.Config{
		Address: mr.Addr(),
	})
	if err != nil {
		t.Fatalf("failed to connect to miniredis: %v", err)
	}
	defer client.Close()

	ser := serializer.NewJSONSerializer()
	c := cache.NewCache(client, ser)
	ctx := context.Background()

	worker := Worker{ID: "w-1", Name: "John"}

	// Test Set
	err = c.Set(ctx, "worker:w-1", worker, 5*time.Minute)
	if err != nil {
		t.Fatalf("failed to set cache item: %v", err)
	}

	// Test Get
	var parsed Worker
	err = c.Get(ctx, "worker:w-1", &parsed)
	if err != nil {
		t.Fatalf("failed to get cache item: %v", err)
	}

	if parsed.ID != worker.ID || parsed.Name != worker.Name {
		t.Errorf("expected %+v, got %+v", worker, parsed)
	}

	// Test GetTTL
	ttl, err := c.GetTTL(ctx, "worker:w-1")
	if err != nil {
		t.Fatalf("failed to fetch TTL: %v", err)
	}
	if ttl <= 0 {
		t.Errorf("expected positive TTL duration, got %v", ttl)
	}

	// Test Delete
	err = c.Delete(ctx, "worker:w-1")
	if err != nil {
		t.Fatalf("failed to delete cache item: %v", err)
	}

	// Test Get on Evicted key (Cache Miss)
	err = c.Get(ctx, "worker:w-1", &parsed)
	if !errors.Is(err, cache.ErrCacheMiss) {
		t.Errorf("expected ErrCacheMiss, got %v", err)
	}
}

func TestCache_BatchOperations(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()

	client, _ := prahariRedis.NewClient(prahariRedis.Config{Address: mr.Addr()})
	defer client.Close()

	c := cache.NewCache(client, serializer.NewJSONSerializer())
	ctx := context.Background()

	pairs := map[string]interface{}{
		"k1": Worker{ID: "1", Name: "A"},
		"k2": Worker{ID: "2", Name: "B"},
	}

	err := c.MSet(ctx, pairs, 1*time.Hour)
	if err != nil {
		t.Fatalf("failed to execute MSet: %v", err)
	}

	keys := []string{"k1", "k2"}
	targets := []interface{}{&Worker{}, &Worker{}}

	err = c.MGet(ctx, keys, targets)
	if err != nil {
		t.Fatalf("failed to execute MGet: %v", err)
	}

	w1 := targets[0].(*Worker)
	w2 := targets[1].(*Worker)

	if w1.Name != "A" || w2.Name != "B" {
		t.Errorf("batch unmarshal failed, got: w1=%v, w2=%v", w1, w2)
	}
}
