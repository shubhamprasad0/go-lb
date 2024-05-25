package main

import (
	lb "github.com/shubhamprasad0/go-lb/pkg"
)

func main() {
	lbConfig := lb.DefaultLBConfig()
	addresses := []string{
		":8000", ":8001", ":8002",
	}
	lbServer := lb.NewLoadBalancer(lbConfig, addresses)
	lbServer.Start()
}
