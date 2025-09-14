package config

import (
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	Port        string
	Host        string
	TemplateDir string
	StaticDir   string
	DevMode     bool
}

// Load loads configuration from environment variables with defaults
func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		Host:        getEnv("HOST", "localhost"),
		TemplateDir: getEnv("TEMPLATE_DIR", "web/templates"),
		StaticDir:   getEnv("STATIC_DIR", "web/static"),
		DevMode:     getEnvAsBool("DEV_MODE", true),
	}
}

// Address returns the full server address
func (c *Config) Address() string {
	return c.Host + ":" + c.Port
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsBool gets an environment variable as boolean or returns a default value
func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	}
	return defaultValue
}