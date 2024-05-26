package lb

import (
	"fmt"
	"net/http"
	"time"
)

type HealthChecker struct {
	AllServers          []string
	HealthyServers      []string
	HealthCheckInterval uint16
	HealthCheckRoute    string
}

func NewHealthChecker(addresses []string, healthCheckRoute string, healthCheckInterval uint16) *HealthChecker {
	return &HealthChecker{
		AllServers:          addresses,
		HealthCheckRoute:    healthCheckRoute,
		HealthCheckInterval: healthCheckInterval,
	}
}

func (h *HealthChecker) performHealthChecks() {
	healthyServers := make([]string, 0)
	for _, server := range h.AllServers {
		if h.isHealthy(server) {
			healthyServers = append(healthyServers, server)
		}
	}
	h.HealthyServers = healthyServers[:]
}

func (h *HealthChecker) start() {
	// once in the beginning
	h.performHealthChecks()

	// then run periodically every `h.HealthCheckInterval` seconds
	for range time.Tick(time.Second * time.Duration(h.HealthCheckInterval)) {
		h.performHealthChecks()
	}
}

func (h *HealthChecker) Run() {
	go func() {
		h.start()
	}()
}

func (h *HealthChecker) GetHealthyServers() []string {
	return h.HealthyServers
}

func (h *HealthChecker) isHealthy(serverAddress string) bool {
	healthCheckAddress := fmt.Sprintf("http://%s%s", serverAddress, h.HealthCheckRoute)
	resp, err := http.Get(healthCheckAddress)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
