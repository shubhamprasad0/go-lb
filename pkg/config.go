package lb

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type LBConfig struct {
	Port                uint16   `yaml:"port"`
	BufferSize          uint16   `yaml:"bufferSize"`
	HealthCheckInterval uint16   `yaml:"healthCheckInterval"`
	HealthCheckRoute    string   `yaml:"healthCheckRoute"`
	Servers             []string `yaml:"servers"`
}

func DefaultLBConfig() *LBConfig {
	return &LBConfig{
		Port:                8080,
		BufferSize:          1024, // bytes
		HealthCheckInterval: 10,   // seconds
		HealthCheckRoute:    "/health",
		Servers:             []string{},
	}
}

func FromYaml(filepath string) *LBConfig {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Panicf("Error reading config file: %+v", err)
	}
	var config LBConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Panicf("Error unmarshaling config file: %+v", err)
	}
	return &config
}

func (c *LBConfig) Update(other *LBConfig) {
	if other.Port != 0 {
		c.Port = other.Port
	}

	if other.BufferSize != 0 {
		c.BufferSize = other.BufferSize
	}

	if other.HealthCheckInterval != 0 {
		c.HealthCheckInterval = other.HealthCheckInterval
	}

	if other.HealthCheckRoute != "" {
		c.HealthCheckRoute = other.HealthCheckRoute
	}

	if len(other.Servers) != 0 {
		c.Servers = other.Servers
	}
}
