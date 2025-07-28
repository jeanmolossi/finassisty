package telemetry

import (
	"context"
	"finassisty/server/config"
	"strings"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

func grpcLogClient(ctx context.Context) (sdklog.Exporter, error) {
	options := []otlploggrpc.Option{
		otlploggrpc.WithHeaders(getLogHeaders()),
	}

	return otlploggrpc.New(ctx, options...)
}

func httpLogClient(ctx context.Context) (sdklog.Exporter, error) {
	options := []otlploghttp.Option{
		otlploghttp.WithHeaders(getLogHeaders()),
	}

	return otlploghttp.New(ctx, options...)
}

func logExporter(ctx context.Context) (sdklog.Exporter, error) {
	env := config.Env()

	isLogProtoGRPC := strings.EqualFold(env.OTelExporter.Logs.Protocol, grpc)
	isGlobalProtoGRPC := strings.EqualFold(env.OTelExporter.Protocol, grpc)
	isLogProtoEmpty := env.OTelExporter.Logs.Protocol == ""

	if isLogProtoGRPC || (isLogProtoEmpty && isGlobalProtoGRPC) {
		logger.Debug("Log exporter using GRPC", "proto", env.OTelExporter.Logs.Protocol)
		return grpcLogClient(ctx)
	}

	logger.Debug("Log exporter using HTTP", "proto", env.OTelExporter.Logs.Protocol)
	return httpLogClient(ctx)
}

func logProvider(ctx context.Context) *sdklog.LoggerProvider {
	exp, err := logExporter(ctx)
	if err != nil {
		panic(err)
	}

	options := []sdklog.BatchProcessorOption{
		sdklog.WithMaxQueueSize(queueSize),
		sdklog.WithExportMaxBatchSize(queueSize),
		sdklog.WithExportInterval(10 * time.Second),
		sdklog.WithExportTimeout(10 * time.Second),
	}

	bsp := sdklog.NewBatchProcessor(exp, options...)

	provider := sdklog.NewLoggerProvider(
		sdklog.WithResource(newResource(ctx)),
		sdklog.WithProcessor(bsp),
	)
	global.SetLoggerProvider(provider)

	return provider
}
