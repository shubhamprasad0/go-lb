package main

import (
	lb "github.com/shubhamprasad0/go-lb/pkg"
)

func main() {
	lbConfig := lb.DefaultLBConfig()
	addresses := []string{
		"127.0.0.1:8000", "127.0.0.1:8001", "127.0.0.1:8002",
	}
	lbServer := lb.NewLoadBalancer(lbConfig, addresses)
	lbServer.Start()
}
