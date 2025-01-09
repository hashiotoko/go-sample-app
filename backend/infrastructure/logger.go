package infrastructure

import (
	"log/slog"
	"os"
	"sync"

	"github.com/hashiotoko/go-sample-app/backend/config"
)

var once sync.Once

func InitLogger() {
	once.Do(func() {
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     config.Config.GetLogLevel(),
		})

		logger := slog.New(handler)
		logger = logger.With(slog.Group("service",
			slog.String("name", config.Config.ServiceName),
			slog.String("version", config.Config.ServiceVersion),
		))

		slog.SetDefault(logger)
	})
}
