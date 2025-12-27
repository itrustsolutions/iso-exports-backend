package customcontext

import (
	"context"

	"github.com/rs/zerolog"
)

type LoggerKey struct{}

func ExtractLogger(ctx context.Context) *zerolog.Logger {
	if logger, ok := ctx.Value(LoggerKey{}).(*zerolog.Logger); ok {
		return logger
	}
	return nil
}

func WithLogger(ctx context.Context, logger *zerolog.Logger) context.Context {
	return context.WithValue(ctx, LoggerKey{}, logger)
}
