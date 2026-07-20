package locks_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	prahariRedis "prahari/shared/redis"
	"prahari/shared/redis/locks"
)

func TestMutex_LockUnlock(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("failed to start miniredis: %v", err)
	}
	defer mr.Close()

	client, err := prahariRedis.NewClient(prahariRedis.Config{
		Address: mr.Addr(),
	})
	if err != nil {
		t.Fatalf("failed to connect to miniredis: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	lockKey := "lock:test-resource"
	holderID := "process-abc"

	m1 := locks.NewMutex(client, lockKey, holderID, 5*time.Second)
	m2 := locks.NewMutex(client, lockKey, "process-xyz", 5*time.Second)

	// Acquire lock
	err = m1.Lock(ctx)
	if err != nil {
		t.Fatalf("expected to acquire lock, got: %v", err)
	}

	// Attempt double acquire -> expect ErrLockAcquireFailed
	err = m2.Lock(ctx)
	if !errors.Is(err, locks.ErrLockAcquireFailed) {
		t.Errorf("expected ErrLockAcquireFailed, got %v", err)
	}

	// Release lock
	err = m1.Unlock(ctx)
	if err != nil {
		t.Fatalf("expected to release lock, got: %v", err)
	}

	// Re-acquire by m2 after release -> expect success
	err = m2.Lock(ctx)
	if err != nil {
		t.Errorf("expected m2 to acquire released lock, got: %v", err)
	}
}

func TestLease_Extension(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()

	client, _ := prahariRedis.NewClient(prahariRedis.Config{Address: mr.Addr()})
	defer client.Close()

	ctx := context.Background()
	lockKey := "lock:lease-resource"
	holderID := "process-lease"

	m := locks.NewMutex(client, lockKey, holderID, 1*time.Second)

	err := m.Lock(ctx)
	if err != nil {
		t.Fatalf("expected lock, got: %v", err)
	}

	// Start lease extension: checks every 100ms and extends to 1s
	cancel := locks.StartLeaseExtension(ctx, m, 100*time.Millisecond, 1*time.Second)

	// Wait 1.5 seconds. If watchdog was not running, lock would have expired.
	time.Sleep(1500 * time.Millisecond)

	// Since watchdog is running, the lock should still be active and owned by us
	m2 := locks.NewMutex(client, lockKey, "process-other", 1*time.Second)
	err = m2.Lock(ctx)
	if !errors.Is(err, locks.ErrLockAcquireFailed) {
		t.Error("expected lock to still be active and blocked due to watchdog extension")
	}

	cancel() // Stop watchdog

	// Fast forward miniredis clock by 2s -> lock should expire
	mr.FastForward(2 * time.Second)

	err = m2.Lock(ctx)
	if err != nil {
		t.Errorf("expected lock to expire and be acquirable after watchdog cancels, got: %v", err)
	}
}
