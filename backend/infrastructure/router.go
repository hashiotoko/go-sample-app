package infrastructure

import (
	"log/slog"
	"net/http"

	api "github.com/hashiotoko/go-sample-app/backend/api/generated"
	validatorMiddleware "github.com/hashiotoko/go-sample-app/backend/middleware/validator"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/oapi-codegen/echo-middleware"
)

const healthCheckPath = "/health"

func Init(router *echo.Echo) {
	router.HideBanner = true
	// router.HidePort = true // サーバー起動時に port を表示しないようにする

	// echo にはリクエストのパラメータなどを検証する機能が無いが設定のためのI/Fだけは生えているので、
	// そのI/Fに沿ったカスタムバリデータを設定する処理
	// ref. https://echo.labstack.com/docs/request#validate-data
	router.Validator = validatorMiddleware.New()

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}
	setupOpenApiValidator(router, swagger)

	router.GET(healthCheckPath, func(c echo.Context) error {
		slog.Info("This service is healthy!")
		return c.NoContent(http.StatusOK)
	})
}

func setupOpenApiValidator(router *echo.Echo, swagger *openapi3.T) {
	swagger.Servers = nil
	router.Use(echoMiddleware.OapiRequestValidatorWithOptions(swagger, &echoMiddleware.Options{
		Skipper: func(c echo.Context) bool {
			return c.Request().RequestURI == healthCheckPath
		},
	}))
}
