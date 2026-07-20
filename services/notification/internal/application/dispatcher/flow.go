package dispatcher

import (
	"context"
	"fmt"

	"prahari/services/notification/internal/domain/channel"
)

// Flow orchestrates dynamically routing payloads to registered adapters.
type Flow struct {
	adapters map[string]channel.Channel
}

// NewFlow constructs a Flow manager.
func NewFlow() *Flow {
	return &Flow{
		adapters: make(map[string]channel.Channel),
	}
}

// Register registers pluggable channel adapters.
func (f *Flow) Register(c channel.Channel) {
	f.adapters[c.Name()] = c
}

// Dispatch executes messaging across matched adapters.
func (f *Flow) Dispatch(ctx context.Context, channelName, recipient, body string) error {
	adapter, exists := f.adapters[channelName]
	if !exists {
		return fmt.Errorf("dispatcher flow: target channel '%s' is not registered", channelName)
	}

	return adapter.Send(ctx, recipient, body)
}
