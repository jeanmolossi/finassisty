package telemetry

import (
	"context"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func grpcMetricClient(ctx context.Context) (sdkmetric.Exporter, error) {
	options := []otlpmetricgrpc.Option{
		otlpmetricgrpc.WithHeaders(getMetricHeaders()),
	}

	return otlpmetricgrpc.New(ctx, options...)
}

func httpMetricClient(ctx context.Context) (sdkmetric.Exporter, error) {
	options := []otlpmetrichttp.Option{
		otlpmetrichttp.WithHeaders(getMetricHeaders()),
	}

	return otlpmetrichttp.New(ctx, options...)
}

func metricExporter(ctx context.Context) (sdkmetric.Exporter, error) {
	isMetricProtoGRPC := strings.EqualFold(env.OTelExporter.Metrics.Protocol, grpc)
	isGlobalProtoGRPC := strings.EqualFold(env.OTelExporter.Protocol, grpc)
	isMetricProtoEmpty := env.OTelExporter.Metrics.Protocol == ""

	if isMetricProtoGRPC || (isMetricProtoEmpty && isGlobalProtoGRPC) {
		logger.Debug("Metric exporter using GRPC", "proto", env.OTelExporter.Metrics.Protocol)
		return grpcMetricClient(ctx)
	}

	logger.Debug("Metric exporter using HTTP", "proto", env.OTelExporter.Metrics.Protocol)
	return httpMetricClient(ctx)
}

func meterProvider(ctx context.Context) *sdkmetric.MeterProvider {
	exp, err := metricExporter(ctx)
	if err != nil {
		panic(err)
	}

	var reader sdkmetric.Reader = sdkmetric.NewPeriodicReader(exp,
		sdkmetric.WithInterval(15*time.Second),
	)

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(reader),
		sdkmetric.WithResource(newResource(ctx)),
	)
	otel.SetMeterProvider(provider)

	return provider
}
