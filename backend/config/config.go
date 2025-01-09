package config

import (
	"log/slog"
	"strings"
	"sync"

	"github.com/caarlos0/env/v11"
)

var Config *AppConfig
var mu sync.Mutex

func LoadAppConfig() {
	mu.Lock()
	defer mu.Unlock()

	Config = &AppConfig{}
	err := env.Parse(Config)
	if err != nil {
		panic(err)
	}
}

type AppConfig struct {
	Environment    string `env:"APP_ENV,required"`
	LogLevel       string `env:"LOG_LEVEL" envDefault:"debug"`
	ServiceName    string `env:"SERVICE_NAME" envDefault:"backend"`
	ServiceVersion string `env:"SERVICE_VERSION" envDefault:"0.0.0"`
}

func (c AppConfig) GetLogLevel() slog.Level {
	switch strings.ToLower(c.LogLevel) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		slog.Warn("Invalid log level, defaulting to debug", "logLevel", c.LogLevel)
		return slog.LevelDebug
	}
}
