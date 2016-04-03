package balancer

import "sync/atomic"

// RoundRobin is a simple load balancer that returns each of the published
// endpoints in sequence.
type RoundRobin struct {
	counter uint64
}

// NewRoundRobin returns a new RoundRobin load balancer.
func NewRoundRobin() *RoundRobin {
	return &RoundRobin{counter: 0}
}

// Balance returns the next server to be used based on the given server list.
// Balance implements the Balance interface.
func (b *RoundRobin) Balance(pool []string) (string, error) {
	if len(pool) <= 0 {
		return "", ErrNoServers
	}

	var old uint64
	for {
		old = atomic.LoadUint64(&b.counter)
		if atomic.CompareAndSwapUint64(&b.counter, old, old+1) {
			break
		}
	}
	return pool[old%uint64(len(pool))], nil
}
