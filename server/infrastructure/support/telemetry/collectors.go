// Package telemetry
package telemetry

import (
	"context"
	"encoding/base64"
	"finassisty/server/config"
	"finassisty/server/infrastructure/support"
	"fmt"
	"maps"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

var (
	logger            = support.Logger
	finassistyHeaders = map[string]string{"X-Finassisty": "abc"}
	env               = config.Env()
)

const (
	grpc      = "grpc"
	queueSize = 16000
)

func getBasic(user string) string {
	authBytes := fmt.Appendf([]byte{}, "%s:%s", user, env.OTelExporter.Password)

	return base64.StdEncoding.EncodeToString(authBytes)
}

func getHeaders(user string) map[string]string {
	h := make(map[string]string)

	maps.Copy(h, finassistyHeaders)

	h["Authorization"] = "Basic " + getBasic(user)

	return h
}

func getMetricHeaders() map[string]string {
	user := config.Env().OTelExporter.Metrics.User
	return getHeaders(user)
}

func getTraceHeaders() map[string]string {
	user := config.Env().OTelExporter.Traces.User
	return getHeaders(user)
}

func getLogHeaders() map[string]string {
	user := config.Env().OTelExporter.Logs.User
	return getHeaders(user)
}

func newResource(ctx context.Context) *resource.Resource {
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceName(config.Env().AppName),
			semconv.ServiceVersion(config.Version()),
		),
	)
	if err != nil {
		otel.Handle(err)
		return resource.Environment()
	}

	return res
}

func StartCollectors(ctx context.Context) error {
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	meterProvider(ctx)
	traceProvider(ctx)
	logProvider(ctx)

	return nil
}
