package balancer

import (
	"github.com/nbio/st"
	"net/http"
	"net/url"
	"testing"
)

func TestBalanceMiddleware(t *testing.T) {
	servers := []string{
		"http://foo.com",
		"http://bar.com",
		"http://baz.com",
	}

	var calls int
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
	})
	r := &http.Request{URL: &url.URL{Host: "invalid"}}

	balancer := New(servers...)

	for i, server := range servers {
		balancer.BalanceHTTP(nil, r, handler)
		st.Expect(t, calls, i+1)

		u, err := url.Parse(server)
		st.Expect(t, err, nil)
		st.Expect(t, r.URL.Host, u.Host)
		st.Expect(t, r.URL.Scheme, u.Scheme)
	}
}

func TestBalanceMiddlewareError(t *testing.T) {
	var called bool
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})
	r := &http.Request{URL: &url.URL{Host: "original", Scheme: "http://"}}

	balancer := New(":\\foo")

	var errored bool
	balancer.OnError(func(err error, w http.ResponseWriter, r *http.Request, h http.Handler) {
		errored = true
		h.ServeHTTP(w, r)
	})

	balancer.BalanceHTTP(nil, r, handler)
	st.Expect(t, called, true)
	st.Expect(t, errored, true)
	st.Expect(t, r.URL.Host, "original")
	st.Expect(t, r.URL.Scheme, "http://")
}
