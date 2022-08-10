package config

import "os"

type GRPCConfig struct {
	URL string
}

type Config struct {
	GRPCConfig
}

func (c *Config) readConfig() {
	grpcURL := os.Getenv("GRPC_URL")
	c.GRPCConfig = GRPCConfig{URL: grpcURL}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
