package log

import (
	"io"
	"log/slog"

	"github.com/dusted-go/logging/prettylog"
)

var (
	logLevel = &slog.LevelVar{}
)

func InitDefaultLoggerWithConsole(logOutput io.Writer) {
	slog.SetDefault(
		slog.New(
			prettylog.New(
				&slog.HandlerOptions{Level: logLevel},
				prettylog.WithDestinationWriter(logOutput),
				prettylog.WithColor(),
			),
		),
	)

	SetRootLogLevel(slog.LevelInfo)
}

func SetRootLogLevel(level slog.Level) {
	oldLevel := logLevel.Level()

	if oldLevel == level {
		return
	}

	logLevel.Set(level)

	slog.Warn(
		"change root logger level",
		slog.String("prev_level", oldLevel.String()),
		slog.String("new_level", level.String()),
	)
}
