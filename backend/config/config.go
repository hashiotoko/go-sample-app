package config

import (
	"log/slog"
	"strings"
	"sync"

	"github.com/caarlos0/env/v11"
)

var Config AppConfig
var mu sync.Mutex

func LoadConfig() {
	mu.Lock()
	defer mu.Unlock()

	err := env.Parse(&Config)
	if err != nil {
		panic(err)
	}
}

type AppConfig struct {
	// Environment    string `env:"APP_ENV,required"`
	LogLevel       string `env:"LOG_LEVEL" envDefault:"debug"`
	ServiceName    string `env:"SERVICE_NAME" envDefault:"backend"`
	ServiceVersion string `env:"SERVICE_VERSION" envDefault:"0.0.0"`
	DB             DBConfig
	AWS            AWSConfig
}

type DBConfig struct {
	Host         string `env:"DB_HOST,required"`
	Port         string `env:"DB_PORT,required"`
	User         string `env:"DB_USER,required"`
	Password     string `env:"DB_PASSWORD,required"`
	Name         string `env:"DB_NAME,required"`
	MaxOpenConns int    `env:"DB_MAX_OPEN_CONNECTIONS,required"`
}

type AWSConfig struct {
	Region string `env:"AWS_REGION" envDefault:"ap-northeast-1"`
	URL    string `env:"AWS_URL"`
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
