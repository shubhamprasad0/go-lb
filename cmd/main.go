package main

import (
	lb "github.com/shubhamprasad0/go-lb/pkg"
)

func main() {
	lbConfig := lb.DefaultLBConfig()
	lbServer := lb.NewLoadBalancer(lbConfig)
	lbServer.Start()
}
