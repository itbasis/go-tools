package log

import "log/slog"

func Panic(msg string, args ...any) {
	slog.Error(msg, args...)

	panic(msg)
}
