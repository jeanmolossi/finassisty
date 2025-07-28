package middlewares

import (
	"finassisty/server/config"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func OTel() echo.MiddlewareFunc {
	return otelecho.Middleware(config.Env().AppName)
}
