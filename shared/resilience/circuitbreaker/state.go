package circuitbreaker

// State represents the current operational mode of a circuit breaker.
type State string

const (
	StateClosed   State = "CLOSED"
	StateOpen     State = "OPEN"
	StateHalfOpen State = "HALF-OPEN"
)
