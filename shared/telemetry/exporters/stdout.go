package exporters

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/sdk/trace"
)

// StdoutExporter implements trace.SpanExporter to print trace logs to the console.
type StdoutExporter struct{}

// NewStdoutExporter constructs a new StdoutExporter.
func NewStdoutExporter() *StdoutExporter {
	return &StdoutExporter{}
}

// ExportSpans prints finished read-only spans details.
func (e *StdoutExporter) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) error {
	for _, s := range spans {
		fmt.Printf("[OTEL-CON] Span: %s, ID: %s, Duration: %v\n", s.Name(), s.SpanContext().SpanID(), s.EndTime().Sub(s.StartTime()))
	}
	return nil
}

// Shutdown handles cleanup actions.
func (e *StdoutExporter) Shutdown(ctx context.Context) error {
	return nil
}
