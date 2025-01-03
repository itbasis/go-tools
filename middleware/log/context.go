package log

import (
	"context"
	"log/slog"
)

type loggerKey struct{}

func WithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, l)
}

func Logger(ctx context.Context) *slog.Logger {
	if ctx == nil {
		panic("nil context")
	}

	if l, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
		return l
	}

	return slog.Default()
}

func Debug(ctx context.Context, msg string, args ...any) {
	Logger(ctx).Debug(msg, args...)
}
