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

	lb := balancer.NewRandom()

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
