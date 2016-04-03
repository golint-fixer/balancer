# balancer [![Build Status](https://travis-ci.org/vinxi/balancer.png)](https://travis-ci.org/vinxi/balancer) [![GoDoc](https://godoc.org/github.com/vinxi/balancer?status.svg)](https://godoc.org/github.com/vinxi/balancer) [![Coverage Status](https://coveralls.io/repos/github/vinxi/balancer/badge.svg?branch=master)](https://coveralls.io/github/vinxi/balancer?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/vinxi/balancer)](https://goreportcard.com/report/github.com/vinxi/balancer) [![API](https://img.shields.io/badge/vinxi-core-green.svg?style=flat)](https://godoc.org/github.com/vinxi/balancer) 

Simple, generic, domain-agnostic balancing algorithms for vinxi.

Currently provides `random` and `roundrobin` distribution algorithms. 
New algorithms may be added in the future.

## Installation

```bash
go get -u gopkg.in/vinxi/balancer.v0
```

## API

See [godoc](https://godoc.org/github.com/vinxi/balancer) reference.

## Example

#### Use as vinxi middleware

```go
package main

import (
  "fmt"
  "gopkg.in/vinxi/balancer.v0"
  "gopkg.in/vinxi/forward.v0"
  "gopkg.in/vinxi/vinxi.v0"
  "net/http"
)

func main() {
  servers := []string{
    "http://www.nytimes.com",
    "http://www.repubblica.it",
    "http://elpais.com",
  }

  // Create the balancer
  lb := balancer.New(servers...)

  // Use custom handler in case of error
  lb.OnError(func(err error, w http.ResponseWriter, r *http.Request, h http.Handler) {
    fmt.Printf("Error: %s\n", err)
    w.WriteHeader(http.StatusBadGateway)
    w.Write([]byte("Error: " + err.Error()))
  })

  // Create a new vinxi proxy
  vs := vinxi.NewServer(vinxi.ServerOptions{Port: 3100})

  vs.Use(lb)

  fw, _ := forward.New(forward.PassHostHeader(true))
  vs.UseFinalHandler(fw)

  fmt.Printf("Server listening on port: %d\n", 3100)
  err := vs.Listen()
  if err != nil {
    fmt.Errorf("Error: %s\n", err)
  }
}
```

#### Round-robin balancing

```go
package main

import (
  "fmt"
  "gopkg.in/vinxi/balancer.v0"
)

func main() {
  servers := []string{
    "http://1.server.com",
    "http://2.server.com",
    "http://3.server.com",
  }

  lb := balancer.NewRoundRobin()

  for i := 0; i < 9; i++ {
    fmt.Println("Next balance round...")
    server, err := lb.Balance(servers)
    if err != nil {
      fmt.Printf("Error: %s\n", err)
      return
    }
    fmt.Printf("Next target server: %s\n", server)
  }
}
```

## License

MIT
