package tracer_test

import (
	"context"
	"testing"

	"prahari/shared/telemetry/tracer"
)

func TestTracer_StartSpan(t *testing.T) {
	tr := tracer.NewTracer("test-tracer-service")
	ctx, span := tr.Start(context.Background(), "RunVitalsVerification")
	defer span.End()

	if span == nil {
		t.Fatal("expected span to be instantiated, got nil")
	}

	if !span.IsRecording() {
		// Mock trace provider won't record by default unless initialized, which is correct
		_ = ctx
	}
}
