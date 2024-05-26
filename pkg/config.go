package lb

type LBConfig struct {
	Address             string
	BufferSize          uint16
	HealthCheckInterval uint16
	HealthCheckRoute    string
}

func DefaultLBConfig() *LBConfig {
	return &LBConfig{
		Address:             ":8080",
		BufferSize:          1024, // bytes
		HealthCheckInterval: 10,   // seconds
		HealthCheckRoute:    "/health",
	}
}
