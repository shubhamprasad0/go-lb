package lb

type LBConfig struct {
	Address    string
	BufferSize uint16
}

func DefaultLBConfig() *LBConfig {
	return &LBConfig{
		Address:    ":8080",
		BufferSize: 1024,
	}
}
