package telemetry

import (
	"context"
	"finassisty/server/config"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func grpcTraceClient() otlptrace.Client {
	options := []otlptracegrpc.Option{
		otlptracegrpc.WithHeaders(getTraceHeaders()),
	}

	return otlptracegrpc.NewClient(options...)
}

func httpTraceClient() otlptrace.Client {
	options := []otlptracehttp.Option{
		otlptracehttp.WithHeaders(getTraceHeaders()),
	}

	return otlptracehttp.NewClient(options...)
}

func traceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	env := config.Env()

	isTraceProtocolGRPC := strings.EqualFold(env.OTelExporter.Traces.Protocol, grpc)
	isGlobalProtocolGRPC := strings.EqualFold(env.OTelExporter.Protocol, grpc)
	isTraceProtocolEmpty := env.OTelExporter.Traces.Protocol == ""

	var client otlptrace.Client

	if isTraceProtocolGRPC || (isTraceProtocolEmpty && isGlobalProtocolGRPC) {
		logger.Debug("Trace exporter using GRPC client", "protocol", env.OTelExporter.Traces.Protocol)
		client = grpcTraceClient()
	} else {
		logger.Debug("Trace exporter using HTTP client", "protocol", env.OTelExporter.Traces.Protocol)
		client = httpTraceClient()
	}

	return otlptrace.New(ctx, client)
}

func traceProvider(ctx context.Context) *sdktrace.TracerProvider {
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(newResource(ctx)),
	)
	otel.SetTracerProvider(provider)

	exp, err := traceExporter(ctx)
	if err != nil {
		logger.Error("failed to start trace client", "err", err)
		return provider
	}

	options := []sdktrace.BatchSpanProcessorOption{
		sdktrace.WithMaxQueueSize(queueSize),
		sdktrace.WithMaxExportBatchSize(queueSize),
		sdktrace.WithBatchTimeout(10 * time.Second),
		sdktrace.WithExportTimeout(10 * time.Second),
	}

	bsp := sdktrace.NewBatchSpanProcessor(exp, options...)
	provider.RegisterSpanProcessor(bsp)

	return provider
}
