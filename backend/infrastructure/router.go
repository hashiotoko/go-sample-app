package infrastructure

import (
	"log/slog"
	"net/http"

	validatorMiddleware "github.com/hashiotoko/go-sample-app/backend/middleware/validator"

	"github.com/labstack/echo/v4"
)

func Init(router *echo.Echo) {
	router.HideBanner = true
	// router.HidePort = true // サーバー起動時に port を表示しないようにする

	// echo にはリクエストのパラメータなどを検証する機能が無いが設定のためのI/Fだけは生えているので、
	// そのI/Fに沿ったカスタムバリデータを設定する処理
	// ref. https://echo.labstack.com/docs/request#validate-data
	router.Validator = validatorMiddleware.New()

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	router.GET("/health", func(c echo.Context) error {
		slog.Info("This service is healthy!")
		return c.NoContent(http.StatusOK)
	})
}
