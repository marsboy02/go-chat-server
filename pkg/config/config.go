package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port        string
	Host        string
	TemplateDir string
	StaticDir   string
	DevMode     bool
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		Host:        getEnv("HOST", "localhost"),
		TemplateDir: getEnv("TEMPLATE_DIR", "web/templates"),
		StaticDir:   getEnv("STATIC_DIR", "web/static"),
		DevMode:     getEnvAsBool("DEV_MODE", true),
	}
}

func (c *Config) Address() string {
	return c.Host + ":" + c.Port
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	}
	return defaultValue
}