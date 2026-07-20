package xray_test

import (
	"context"
	"testing"

	prahariXRay "prahari/shared/aws/xray"
)

func TestClient_SegmentLifecycle(t *testing.T) {
	client := prahariXRay.NewClient(true)
	
	err := client.Ping(context.Background())
	if err != nil {
		t.Fatalf("expected no ping error from X-Ray daemon checker, got: %v", err)
	}

	ctx, segment := client.StartSegment(context.Background(), "test-db-query")
	if segment == nil {
		t.Fatal("expected segment to be instantiated, got nil")
	}
	
	_ = ctx
	segment.Close()
}
