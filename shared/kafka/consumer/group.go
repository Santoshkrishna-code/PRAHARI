package consumer

import (
	"context"
	"sync"
)

// Group aggregates and runs multiple Consumers concurrently (scaling horizontally).
type Group struct {
	consumers []*Consumer
	wg        sync.WaitGroup
}

// NewGroup constructs a consumer group executor.
func NewGroup(consumers ...*Consumer) *Group {
	return &Group{consumers: consumers}
}

// Start launches partition read loops for all aggregated consumers.
func (g *Group) Start(ctx context.Context, handler func(ctx context.Context, key, val []byte) error) {
	for _, c := range g.consumers {
		g.wg.Add(1)
		go func(cons *Consumer) {
			defer g.wg.Done()
			_ = cons.StartConsumer(ctx, handler)
		}(c)
	}
}

// Wait blocks until all group threads exit.
func (g *Group) Wait() {
	g.wg.Wait()
}
