package resilience

// Manager consolidates configurations for execution wrappers.
type Manager struct {
	cfg Config
}

// NewManager constructs a new Manager instance.
func NewManager(cfg Config) *Manager {
	return &Manager{cfg: cfg}
}
