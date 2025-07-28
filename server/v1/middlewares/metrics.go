package middlewares

import (
	"finassisty/server/config"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

func MetrifyRequest() echo.MiddlewareFunc {
	meterProvider := otel.GetMeterProvider()
	meter := meterProvider.Meter(env.AppName)

	requestCounter, err := meter.Int64Counter("request")
	if err != nil {
		panic(err)
	}

	attrs := []attribute.KeyValue{
		semconv.ServiceVersion(config.Version()),
		semconv.ServiceName(env.AppName),
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			ctx := req.Context()

			reqAttrs := append(attrs,
				attribute.String("http.request.method", req.Method),
				semconv.UserAgentName(req.UserAgent()),
				semconv.URLQuery(req.URL.RawQuery),
				semconv.URLPath(req.URL.Path),
			)

			err := next(c)
			if err != nil {
				requestCounter.Add(ctx, 1, metric.WithAttributes(append(reqAttrs,
					semconv.ErrorMessage(err.Error()),
					semconv.HTTPResponseStatusCode(res.Status),
				)...))

				c.Error(err)
				return err
			}

			requestCounter.Add(ctx, 1, metric.WithAttributes(append(reqAttrs,
				semconv.HTTPResponseStatusCode(res.Status),
				semconv.HTTPResponseSize(int(res.Size)),
			)...))

			return nil
		}
	}
}
