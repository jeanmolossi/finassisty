package config

import (
	supenv "finassisty/server/infrastructure/support/env"
	"sync"

	"github.com/caarlos0/env/v11"
)

type OTLPDataConfig struct {
	Endpoint string            `env:"ENDPOINT"`
	Headers  map[string]string `env:"HEADERS" envSeparator:"," envKeyValSeparator:"="`
	Protocol string            `env:"PROTOCOL"`
	User     string            `env:"USERNAME"`
}

type ExporterOTLPConfig struct {
	Endpoint string            `env:"ENDPOINT"`
	Headers  map[string]string `env:"HEADERS" envSeparator:"," envKeyValSeparator:"="`
	Protocol string            `env:"PROTOCOL" envDefault:"http/protobuf"`
	Password string            `env:"PASSWORD"`

	Traces  OTLPDataConfig `envPrefix:"TRACES_"`
	Metrics OTLPDataConfig `envPrefix:"METRICS_"`
	Logs    OTLPDataConfig `envPrefix:"LOGS_"`
}

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"Finassisty"`

	OTelExporter ExporterOTLPConfig `envPrefix:"OTEL_EXPORTER_OTLP_"`
}

var (
	config Config

	Env = sync.OnceValue(loadEnv)
)

func init() {
	Env()
}

func loadEnv() *Config {
	supenv.LoadEnv()

	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
