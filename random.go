package balancer

import "math/rand"

// Random is a simple load balance using a random seed for balancing.
type Random struct {
	rand *rand.Rand
}

// NewRandom creates a new random algorithm to balance servers.
func NewRandom() *Random {
	return &Random{rand: rand.New(rand.NewSource(1 << 10))}
}

// Balance returns the next server to be used.
// Balance implements the Balance interface.
func (b *Random) Balance(pool []string) (string, error) {
	if len(pool) <= 0 {
		return "", ErrNoServers
	}
	return pool[b.rand.Intn(len(pool))], nil
}
