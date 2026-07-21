package circuitbreaker

import (
	"errors"
	"time"

	"github.com/sony/gobreaker"
)

var (
	// ErrCircuitOpen is returned when the circuit breaker rejects executions.
	ErrCircuitOpen = errors.New("circuitbreaker: request blocked because circuit is open")
)

// CircuitBreaker wraps gobreaker.CircuitBreaker to handle request failures.
type CircuitBreaker struct {
	cb *gobreaker.CircuitBreaker
}

// NewCircuitBreaker constructs a CircuitBreaker.
func NewCircuitBreaker(name string, failureRatio float64, cooldown time.Duration) *CircuitBreaker {
	settings := gobreaker.Settings{
		Name:    name,
		Timeout: cooldown,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Require at least 5 requests before tripping to avoid premature triggers
			if counts.Requests < 5 {
				return false
			}
			ratio := float64(counts.TotalFailures) / float64(counts.Requests)
			return ratio >= failureRatio
		},
	}

	return &CircuitBreaker{
		cb: gobreaker.NewCircuitBreaker(settings),
	}
}

// Execute runs the function wrapped by the circuit breaker logic.
func (c *CircuitBreaker) Execute(fn func() (interface{}, error)) (interface{}, error) {
	if c.cb == nil {
		return nil, errors.New("circuit breaker is uninitialized")
	}

	val, err := c.cb.Execute(fn)
	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			return nil, ErrCircuitOpen
		}
		return nil, err
	}

	return val, nil
}

// GetState returns the current status (CLOSED, OPEN, HALF-OPEN).
func (c *CircuitBreaker) GetState() State {
	if c.cb == nil {
		return StateClosed
	}

	switch c.cb.State() {
	case gobreaker.StateClosed:
		return StateClosed
	case gobreaker.StateOpen:
		return StateOpen
	case gobreaker.StateHalfOpen:
		return StateHalfOpen
	default:
		return StateClosed
	}
}
