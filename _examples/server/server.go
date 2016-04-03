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
