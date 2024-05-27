package lb

import (
	"fmt"
	"net/http"
	"time"
)

// HealthChecker represents a health checker that monitors the health of application servers.
type HealthChecker struct {
	AllServers          []string
	HealthyServers      []string
	HealthCheckInterval uint16
	HealthCheckRoute    string
}

// NewHealthChecker creates a new HealthChecker instance with the provided servers, health check route, and health check interval.
func NewHealthChecker(servers []string, healthCheckRoute string, healthCheckInterval uint16) *HealthChecker {
	return &HealthChecker{
		AllServers:          servers,
		HealthCheckRoute:    healthCheckRoute,
		HealthCheckInterval: healthCheckInterval,
	}
}

// performHealthChecks checks the health of all servers and updates the list of healthy servers.
func (h *HealthChecker) performHealthChecks() {
	healthyServers := make([]string, 0)
	for _, server := range h.AllServers {
		if h.isHealthy(server) {
			healthyServers = append(healthyServers, server)
		}
	}
	h.HealthyServers = healthyServers[:]
}

// start initiates the health checks, performing them initially and then periodically based on the health check interval.
func (h *HealthChecker) start() {
	// Perform health checks once at the beginning
	h.performHealthChecks()

	// Then run periodically every `h.HealthCheckInterval` seconds
	for range time.Tick(time.Second * time.Duration(h.HealthCheckInterval)) {
		h.performHealthChecks()
	}
}

// Run starts the health checker in a separate goroutine.
func (h *HealthChecker) Run() {
	go func() {
		h.start()
	}()
}

// GetHealthyServers returns the list of currently healthy servers.
func (h *HealthChecker) GetHealthyServers() []string {
	return h.HealthyServers
}

// isHealthy checks if a specific server is healthy by making an HTTP GET request to the health check route.
func (h *HealthChecker) isHealthy(serverAddress string) bool {
	healthCheckAddress := fmt.Sprintf("http://%s%s", serverAddress, h.HealthCheckRoute)
	resp, err := http.Get(healthCheckAddress)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
