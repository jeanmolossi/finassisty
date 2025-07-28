package middlewares

import (
	"finassisty/server/config"
	"finassisty/server/infrastructure/support"
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	logger = support.Logger
	env    = config.Env()
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			ctx := req.Context()

			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
				return err
			}

			dur := time.Since(start)

			logger.InfoContext(ctx, "http request",
				slog.String("started_at", start.Format(time.RFC3339)),
				slog.String("finished_at", time.Now().Format(time.RFC3339)),
				slog.Group("http",
					slog.Group("request",
						slog.String("method", req.Method),
						slog.String("path", req.URL.Path),
						slog.String("query", req.URL.RawQuery),
						slog.String("user_agent", req.UserAgent()),
						slog.Any("headers", req.Header),
					),
					slog.Group("response",
						slog.Int("status_code", res.Status),
						slog.Any("headers", res.Header()),
						slog.Int64("content_length", res.Size),
						slog.String("human_duration", dur.String()),
						slog.Duration("duration", dur),
					),
				),
			)

			return nil
		}
	}
}
