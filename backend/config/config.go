package config

import (
	"sync"

	"github.com/caarlos0/env/v11"
)

var Config *AppConfig
var mu sync.Mutex

type AppConfig struct {
	Environment string `env:"APP_ENV,required"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"debug"`
}

func LoadAppConfig() {
	mu.Lock()
	defer mu.Unlock()

	Config = &AppConfig{}
	err := env.Parse(Config)
	if err != nil {
		panic(err)
	}
}
