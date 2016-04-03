package balancer

import (
	"github.com/nbio/st"
	"testing"
)

func TestRoundRobinDistribution(t *testing.T) {
	pool := []string{"foo", "bar", "baz"}

	lb := NewRoundRobin()

	for _, want := range []string{
		"foo", "bar", "baz",
		"foo", "bar", "baz",
		"foo", "bar", "baz",
		"foo", "bar", "baz",
	} {
		server, err := lb.Balance(pool)
		st.Expect(t, err, nil)
		st.Expect(t, server, want)
	}
}

func TestRoundRobinError(t *testing.T) {
	random := NewRoundRobin()
	_, err := random.Balance([]string{})
	st.Expect(t, err, ErrNoServers)
}
