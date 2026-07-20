package health

import (
	"context"

	prahariRes "prahari/shared/resilience"
)

// Registry aggregates downstream checkers to expose a single diagnostics registry.
type Registry struct {
	checkers []prahariRes.Checker
}

// NewRegistry constructs a Registry.
func NewRegistry(checkers ...prahariRes.Checker) *Registry {
	return &Registry{
		checkers: checkers,
	}
}

// CheckHealth queries all registered downstreams and compiles a connection errors map.
func (r *Registry) CheckHealth(ctx context.Context) map[string]error {
	results := make(map[string]error)

	for _, checker := range r.checkers {
		if checker == nil {
			continue
		}
		if err := checker.Ping(ctx); err != nil {
			results[checker.Name()] = err
		} else {
			results[checker.Name()] = nil
		}
	}

	return results
}
