package loadbalancer

import (
	"errors"
	"sync/atomic"
)

// Balancer manages sequential server routing via Round-Robin calculations.
type Balancer struct {
	hosts []string
	index uint64
}

// NewBalancer constructs a Balancer.
func NewBalancer(hosts []string) *Balancer {
	return &Balancer{
		hosts: hosts,
		index: 0,
	}
}

// Next selects the target host sequentially.
func (b *Balancer) Next() (string, error) {
	if len(b.hosts) == 0 {
		return "", errors.New("loadbalancer: no active hosts registered")
	}

	// Increment sequentially, wrapping around using modulus
	idx := atomic.AddUint64(&b.index, 1) - 1
	target := b.hosts[idx%uint64(len(b.hosts))]

	return target, nil
}
