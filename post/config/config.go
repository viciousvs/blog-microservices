package config

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	Addres string
}
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type PostgresConfig struct {
	User,
	Password,
	Host,
	Port,
	DatabaseName string
}

type Config struct {
	Server   ServerConfig
	Redis    RedisConfig
	Postgres PostgresConfig
}

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

func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Addres: getEnv("SERVER_ADDR", "localhost:50051"),
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		Postgres: PostgresConfig{
			User:         getEnv("POSTGRES_USER", "postgres"),
			Password:     getEnv("POSTGRES_PASSWORD", "postgres"),
			Host:         getEnv("POSTGRES_HOST", "localhost"),
			Port:         getEnv("POSTGRES_PORT", "5432"),
			DatabaseName: getEnv("POSTGRES_DBNAME", "postgres"),
		},
	}
}
