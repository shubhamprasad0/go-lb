package main

import (
	"flag"
	"strings"

	lb "github.com/shubhamprasad0/go-lb/pkg"
)

func main() {
	// command line arguments
	configFile := flag.String("config", "", "Path to the YAML configuration file")
	port := flag.Uint("port", 0, "Port on which the load balancer runs")
	bufferSize := flag.Uint("buffer-size", 0, "Size of buffer (in bytes) used while reading request data")
	healthCheckRoute := flag.String("health-check-route", "", "Health check endpoint on the application servers")
	healthCheckInterval := flag.Uint("health-check-interval", 0, "Number of seconds after which health check is performed periodically")
	servers := flag.String("servers", "", "comma-separated list of application server addresses")

	// parse command line arguments
	flag.Parse()

	config := lb.DefaultLBConfig()
	if configFile != nil && *configFile != "" {
		configFromYaml := lb.FromYaml(*configFile)
		config.Update(configFromYaml)
	}

	configFromCLI := &lb.LBConfig{
		Port:                uint16(*port),
		BufferSize:          uint16(*bufferSize),
		HealthCheckInterval: uint16(*healthCheckInterval),
		HealthCheckRoute:    *healthCheckRoute,
		Servers:             strings.Split(*servers, ","),
	}
	config.Update(configFromCLI)

	// start load balancer
	lbServer := lb.NewLoadBalancer(config)
	lbServer.Start()
}
