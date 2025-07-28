// Package config centralizes the application's environment configuration.
package config

import (
	supenv "finassisty/server/infrastructure/support/env"
	"sync"

	"github.com/caarlos0/env/v11"
)

// OTLPDataConfig holds exporter endpoint information.
type OTLPDataConfig struct {
	Endpoint string            `env:"ENDPOINT"`
	Headers  map[string]string `env:"HEADERS" envSeparator:"," envKeyValSeparator:"="`
	Protocol string            `env:"PROTOCOL"`
	User     string            `env:"USERNAME"`
}

// ExporterOTLPConfig aggregates tracing, metrics and logs exporter configs.
type ExporterOTLPConfig struct {
	Endpoint string            `env:"ENDPOINT"`
	Headers  map[string]string `env:"HEADERS" envSeparator:"," envKeyValSeparator:"="`
	Protocol string            `env:"PROTOCOL" envDefault:"http/protobuf"`
	Password string            `env:"PASSWORD"`

	Traces  OTLPDataConfig `envPrefix:"TRACES_"`
	Metrics OTLPDataConfig `envPrefix:"METRICS_"`
	Logs    OTLPDataConfig `envPrefix:"LOGS_"`
}

// GoogleOAuth holds credential configuration for OAuth login.
type GoogleOAuth struct {
	ClientID     string `env:"GOOGLE_CLIENT_ID"`
	ClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
	RedirectURL  string `env:"GOOGLE_REDIRECT_URL"`
}

// Config represents the full environment configuration parsed from .env.
type Config struct {
	AppName string `env:"APP_NAME" envDefault:"Finassisty"`

	OTelExporter ExporterOTLPConfig `envPrefix:"OTEL_EXPORTER_OTLP_"`

	GoogleOAuth GoogleOAuth
}

var (
	config Config

	// Env loads configuration once and caches the result.
	Env = sync.OnceValue(loadEnv)
)

func init() {
	Env()
}

// loadEnv parses environment variables into Config.
func loadEnv() *Config {
	supenv.LoadEnv()

	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
