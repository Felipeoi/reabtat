package config

import (
	"log"
	"os"
)

type Config struct {
	Env         string
	HTTP        struct{ Port string }
	DB          struct{ DSN string }
	JWT         struct{ Secret string }
	FrontendURL string
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("variável de ambiente obrigatória não encontrada: %s", key)
	}
	return v
}

func Load() *Config {
	var c Config
	c.Env = getenv("APP_ENV", "dev")
	c.HTTP.Port = getenv("PORT", "8080")

	// estas OBRIGATÓRIAS:
	c.DB.DSN = mustEnv("DB_DSN")
	c.JWT.Secret = mustEnv("JWT_SECRET")

	c.FrontendURL = getenv("FRONTEND_URL", "http://localhost:5173")

	return &c
}
