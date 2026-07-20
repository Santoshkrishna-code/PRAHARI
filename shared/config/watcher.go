package config

import (
	"sync"
)

// Watcher manages thread-safe runtime configuration updates.
type Watcher struct {
	mu        sync.RWMutex
	current   *Config
	listeners []func(*Config)
}

// NewWatcher instantiates a configuration wrapper.
func NewWatcher(initial *Config) *Watcher {
	return &Watcher{
		current:   initial,
		listeners: make([]func(*Config), 0),
	}
}

// Get returns the active configuration in a thread-safe manner.
func (w *Watcher) Get() *Config {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.current
}

// Update replaces the configuration in memory and fires update listeners.
func (w *Watcher) Update(newConfig *Config) {
	w.mu.Lock()
	w.current = newConfig
	listeners := make([]func(*Config), len(w.listeners))
	copy(listeners, w.listeners)
	w.mu.Unlock()

	// Notify listeners asynchronously to prevent deadlocking callers
	for _, fn := range listeners {
		go fn(newConfig)
	}
}

// RegisterListener appends callback routines to be notified of updates.
func (w *Watcher) RegisterListener(fn func(*Config)) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.listeners = append(w.listeners, fn)
}
