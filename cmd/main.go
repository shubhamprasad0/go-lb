package main

import (
	"flag"
	"log"
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
	var servers []string
	flag.Func("servers", "comma-separated list of application server addresses", func(s string) error {
		splits := strings.Split(s, ",")
		for _, server := range splits {
			if server != "" {
				servers = append(servers, server)
			}
		}
		return nil
	})

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
		Servers:             servers,
	}
	config.Update(configFromCLI)

	if len(config.Servers) == 0 {
		log.Panicf("No application servers provided. Exiting.")
	}

	// start load balancer
	lbServer := lb.NewLoadBalancer(config)
	lbServer.Start()
}
