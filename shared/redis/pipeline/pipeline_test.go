package pipeline_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	prahariRedis "prahari/shared/redis"
	"prahari/shared/redis/pipeline"
)

func TestPipeline_Exec(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()

	client, _ := prahariRedis.NewClient(prahariRedis.Config{Address: mr.Addr()})
	defer client.Close()

	ctx := context.Background()
	pipe, err := pipeline.NewPipeline(client)
	if err != nil {
		t.Fatalf("failed to create pipeline: %v", err)
	}

	setCmd := pipe.Set(ctx, "pipe-key-1", "value-1", 1*time.Minute)
	getCmd := pipe.Get(ctx, "pipe-key-1")

	_, err = pipe.Exec(ctx)
	if err != nil {
		t.Fatalf("pipeline execution failed: %v", err)
	}

	if setCmd.Val() != "OK" {
		t.Errorf("expected OK, got %s", setCmd.Val())
	}

	if getCmd.Val() != "value-1" {
		t.Errorf("expected value-1, got %s", getCmd.Val())
	}
}
