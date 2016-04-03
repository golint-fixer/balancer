package balancer

import (
	"github.com/nbio/st"
	"math"
	"testing"
)

func TestRandomDistribution(t *testing.T) {
	lb := NewRandom()
	pool := []string{"foo", "bar", "baz"}
	iterations := 10000
	buf := make(map[string]int)

	for i := 0; i < iterations; i++ {
		server, err := lb.Balance(pool)
		st.Expect(t, err, nil)
		buf[server]++
	}

	for server, count := range buf {
		avg := math.Abs((float64(count) * 100) / float64(iterations))
		if avg < 32 || avg > 34 {
			t.Errorf("%s: invalid random distribution: %d", server, avg)
		}
	}
}

func TestRandomError(t *testing.T) {
	random := NewRandom()
	_, err := random.Balance([]string{})
	st.Expect(t, err, ErrNoServers)
}
