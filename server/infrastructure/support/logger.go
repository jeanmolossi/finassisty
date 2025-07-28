// Package support
package support

import (
	"context"
	"finassisty/server/config"
	"log/slog"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"
)

var Logger = slog.New(
	NewTeeHandler(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
		NewOTelHandler(),
	),
)

func NewOTelHandler() slog.Handler {
	appName := config.Env().AppName
	return otelslog.NewHandler(appName)
}

type TeeHandler struct {
	handlers []slog.Handler
}

var _ (slog.Handler) = (*TeeHandler)(nil)

func NewTeeHandler(handlers ...slog.Handler) *TeeHandler {
	if len(handlers) == 0 {
		panic("provide at least one not nil slog.Handler")
	}

	if handlers[0] == nil {
		panic("handlers should not be nil")
	}

	return &TeeHandler{handlers}
}

// Enabled implements slog.Handler.
func (t *TeeHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	for _, h := range t.handlers {
		if h.Enabled(ctx, lvl) {
			return true
		}
	}

	return false
}

// Handle implements slog.Handler.
func (t *TeeHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, h := range t.handlers {
		if err := h.Handle(ctx, record); err != nil {
			return err
		}
	}

	return nil
}

// WithAttrs implements slog.Handler.
func (t *TeeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	for i, h := range t.handlers {
		t.handlers[i] = h.WithAttrs(attrs)
	}

	return t
}

// WithGroup implements slog.Handler.
func (t *TeeHandler) WithGroup(name string) slog.Handler {
	for i, h := range t.handlers {
		t.handlers[i] = h.WithGroup(name)
	}

	return t
}
