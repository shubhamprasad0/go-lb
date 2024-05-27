package lb

import (
	"os"
	"testing"

	testUtils "github.com/shubhamprasad0/go-lb/test"
	"github.com/stretchr/testify/assert"
)

func TestDefaultLBConfig(t *testing.T) {
	config := DefaultLBConfig()
	assert.Equal(t, uint16(8080), config.Port)
	assert.Equal(t, uint16(1024), config.BufferSize)
	assert.Equal(t, uint16(10), config.HealthCheckInterval)
	assert.Equal(t, "/health", config.HealthCheckRoute)
	assert.Empty(t, config.Servers)
}

func TestFromYaml(t *testing.T) {
	yamlData := `
port: 9090
bufferSize: 2048
healthCheckInterval: 20
healthCheckRoute: "/status"
servers:
  - "server1:8080"
  - "server2:8080"
`
	file := testUtils.CreateTempFile(t, "config.yaml", yamlData)
	defer os.Remove(file.Name())

	config := FromYaml(file.Name())
	assert.Equal(t, uint16(9090), config.Port)
	assert.Equal(t, uint16(2048), config.BufferSize)
	assert.Equal(t, uint16(20), config.HealthCheckInterval)
	assert.Equal(t, "/status", config.HealthCheckRoute)
	assert.Equal(t, []string{"server1:8080", "server2:8080"}, config.Servers)
}

func TestUpdate(t *testing.T) {
	config := LBConfig{
		Port:                8080,
		BufferSize:          1024,
		HealthCheckInterval: 10,
		HealthCheckRoute:    "/health",
		Servers:             []string{"127.0.0.1:8000", "127.0.0.1:8001"},
	}
	other := LBConfig{
		HealthCheckInterval: 20,
		HealthCheckRoute:    "/ready",
		Servers:             []string{"127.0.0.1:8003", "127.0.0.1:8004"},
	}
	config.Update(&other)
	assert.Equal(t, uint16(8080), config.Port)
	assert.Equal(t, uint16(1024), config.BufferSize)
	assert.Equal(t, uint16(20), config.HealthCheckInterval)
	assert.Equal(t, "/ready", config.HealthCheckRoute)
	assert.Equal(t, []string{"127.0.0.1:8003", "127.0.0.1:8004"}, config.Servers)
}
