package log

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

var (
	GrpcLogInterceptorOpts = []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}
)

func GrpcInterceptorLogger() logging.Logger {
	return logging.LoggerFunc(
		func(ctx context.Context, level logging.Level, msg string, fields ...any) {
			slog.Log(ctx, slog.Level(level), msg, fields...)
		},
	)
}
