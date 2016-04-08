// Package balancer provides simple well-known distribution
// algorithms for easy traffic balancing.
package balancer

import (
	"errors"
	"net/http"
	"net/url"

	"gopkg.in/vinxi/layer.v0"
)

// ErrNoServers is returned when no servers are passed to the balancer.
var ErrNoServers = errors.New("balancer: server list is empty")

// DefaultBalancer is the default balancer used by the middleware.
var DefaultBalancer = NewRoundRobin()

// Balancer represents the interface implemented by balancer providers.
type Balancer interface {
	// Balance balances the given servers and returns
	// the next server or an error.
	Balance([]string) (string, error)
}

// ErrorHandler represents the required interface implemented by balancer error handlers.
type ErrorHandler func(error, http.ResponseWriter, *http.Request, http.Handler)

// DefaultErrorHandler is used as default error handler, which is a no-op handler.
var DefaultErrorHandler = func(err error, w http.ResponseWriter, r *http.Request, h http.Handler) {
	h.ServeHTTP(w, r)
}

// ProxyBalancer balances an incoming HTTP request across a pool of server
// based on the selected balancing algorithm.
type ProxyBalancer struct {
	// Servers stores the list of server URLs to balance.
	Servers []string

	// Balancer stores the balancer to be used.
	Balancer Balancer

	// ErrorHandler optionally stores the error handler to handle the balancer error and reply accordingly.
	ErrorHandler ErrorHandler
}

// New creates a new balancer which will balance traffic across the specific URLs.
func New(servers ...string) *ProxyBalancer {
	return &ProxyBalancer{
		Servers:      servers,
		Balancer:     DefaultBalancer,
		ErrorHandler: DefaultErrorHandler,
	}
}

// OnError defines the error handler to be used to handle the balancer error.
func (b *ProxyBalancer) OnError(handler ErrorHandler) {
	b.ErrorHandler = handler
}

// Register registers the middleware.
func (b *ProxyBalancer) Register(mw layer.Middleware) {
	mw.UsePriority("request", layer.Head, b.BalanceHTTP)
}

// BalanceHTTP handles an incoming HTTP request and defines
// the target server to forward the request.
func (b *ProxyBalancer) BalanceHTTP(w http.ResponseWriter, r *http.Request, h http.Handler) {
	server, err := b.Balancer.Balance(b.Servers)
	if err != nil {
		b.ErrorHandler(err, w, r, h)
		return
	}

	uri, err := url.Parse(server)
	if err != nil {
		b.ErrorHandler(err, w, r, h)
		return
	}

	r.Host = uri.Host
	r.URL.Host = uri.Host
	r.URL.Scheme = uri.Scheme

	h.ServeHTTP(w, r)
}
