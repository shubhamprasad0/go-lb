package lb

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type LBConfig struct {
	Address             string   `yaml:"address"`
	BufferSize          uint16   `yaml:"bufferSize"`
	HealthCheckInterval uint16   `yaml:"healthCheckInterval"`
	HealthCheckRoute    string   `yaml:"healthCheckRoute"`
	Servers             []string `yaml:"servers"`
}

func DefaultLBConfig() *LBConfig {
	return &LBConfig{
		Address:             ":8080",
		BufferSize:          1024, // bytes
		HealthCheckInterval: 10,   // seconds
		HealthCheckRoute:    "/health",
		Servers: []string{
			"127.0.0.1:8000", "127.0.0.1:8001", "127.0.0.1:8002",
		},
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
