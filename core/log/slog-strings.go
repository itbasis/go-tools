package log

import (
	"log/slog"
)

func SlogAttrSlice[S ~[]T, T any](key string, value S) slog.Attr {
	return SlogAttrSliceWithSeparator(key, "", value)
}

func SlogAttrSliceWithSeparator[S ~[]T, T any](key, separator string, value S) slog.Attr {
	if len(value) == 0 {
		return slog.String(key, "")
	}

	var valueAsString = slog.AnyValue(value[0]).String()

	for _, v := range value[1:] {
		valueAsString = valueAsString + separator + slog.AnyValue(v).String()
	}

	return slog.String(key, valueAsString)
}
