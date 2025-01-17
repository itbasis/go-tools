package log

import (
	"log/slog"
)

func SlogAttrMap[M ~map[K]V, K comparable, V any](key string, value M) slog.Attr {
	var attrs = make([]slog.Attr, 0, len(value))

	for k, v := range value {
		attrs = append(attrs, slog.Any(slog.AnyValue(k).String(), v))
	}

	return slog.Attr{Key: key, Value: slog.GroupValue(attrs...)}
}
