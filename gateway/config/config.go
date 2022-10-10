package config

import (
	"os"
	"strconv"
	"time"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); !exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, strconv.Itoa(defaultValue))
	if v, err := strconv.Atoi(valueStr); err == nil {
		return v
	}
	return defaultValue
}

type Config struct {
	ServerConfig     ServerConfig
	PostClientConfig PostClientConfig
}
type PostClientConfig struct {
	Addr string
}
type ServerConfig struct {
	Addr           string
	MaxHeaderbytes int
	ReadTimeout,
	WriteTimeout time.Duration
}

func NewConfig() *Config {
	return &Config{
		ServerConfig: ServerConfig{
			Addr:           getEnv("HTTP_SERVER_ADDRESS", "localhost:8081"),
			MaxHeaderbytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
		PostClientConfig: PostClientConfig{
			Addr: getEnv("GRPC_CLIENT_ADDRESS", "localhost:50051"),
		},
	}

}
