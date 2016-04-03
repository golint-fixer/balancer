# balancer [![Build Status](https://travis-ci.org/vinxi/balancer.png)](https://travis-ci.org/vinxi/balancer) [![GoDoc](https://godoc.org/github.com/vinxi/balancer?status.svg)](https://godoc.org/github.com/vinxi/balancer) [![Coverage Status](https://coveralls.io/repos/github/vinxi/balancer/badge.svg?branch=master)](https://coveralls.io/github/vinxi/balancer?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/vinxi/balancer)](https://goreportcard.com/report/github.com/vinxi/balancer) [![API](https://img.shields.io/badge/vinxi-core-green.svg?style=flat)](https://godoc.org/github.com/vinxi/balancer) 

Simple, generic, domain-agnostic balancing algorithms for vinxi.

## Installation

```bash
go get -u gopkg.in/vinxi/balancer.v0
```

## API

See [godoc](https://godoc.org/github.com/vinxi/balancer) reference.

## Example

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
