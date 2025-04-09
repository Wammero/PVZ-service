package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Redis    RedisConfig
	JWT      JWTConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Port string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type JWTConfig struct {
	SecretKey string
}

func NewConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     getEnvOrFatal("DATABASE_HOST"),
			Port:     getEnvOrFatal("DATABASE_PORT"),
			User:     getEnvOrFatal("DATABASE_USER"),
			Password: getEnvOrFatal("DATABASE_PASSWORD"),
			Name:     getEnvOrFatal("DATABASE_NAME"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Redis: RedisConfig{
			Host:     getEnvOrFatal("REDIS_HOST"),
			Port:     getEnvOrFatal("REDIS_PORT"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       parseRedisDB(getEnv("REDIS_DB", "0")),
		},
		JWT: JWTConfig{
			SecretKey: getEnvOrFatal("JWT_SECRET_KEY"),
		},
	}
}

func (db DatabaseConfig) GetConnStr() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		db.User, db.Password, db.Host, db.Port, db.Name)
}

func parseRedisDB(db string) int {
	parsed, err := strconv.Atoi(db)
	if err != nil {
		log.Fatalf("Invalid REDIS_DB value: %s", db)
	}
	return parsed
}

func getEnvOrFatal(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return val
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
