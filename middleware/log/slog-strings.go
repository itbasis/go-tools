package log

import (
	"log/slog"
	"strings"
)

func SlogAttrStrings(key string, value []string) slog.Attr {
	return slog.String(key, strings.Join(value, ""))
}

func SlogAttrStringsWithSeparator(key, separator string, value []string) slog.Attr {
	return slog.String(key, strings.Join(value, separator))
}
