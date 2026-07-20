package redis

// Manager wraps client configurations and utility sub-modules.
type Manager struct {
	Client *Client
}

// NewManager constructs a unified Redis manager.
func NewManager(client *Client) *Manager {
	return &Manager{
		Client: client,
	}
}
