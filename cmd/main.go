package main

import (
	lb "github.com/shubhamprasad0/go-lb/pkg"
)

func main() {
	lbConfig := lb.FromYaml("config/conf.yaml")
	lbServer := lb.NewLoadBalancer(lbConfig)
	lbServer.Start()
}
