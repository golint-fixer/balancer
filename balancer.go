// Package balancer provides simple well-known distribution
// algorithms for easy traffic balancing.
package balancer

import "errors"

var (
	// ErrNoServers is returned when no servers are passed to the balancer.
	ErrNoServers = errors.New("balancer: server list is empty")
)

// Balancer represents the interface implemented by balancer providers.
type Balancer interface {
	// Balance balances the given servers and returns
	// the next server or an error.
	Balance([]string) (string, error)
}
