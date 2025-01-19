package token

import (
	"fmt"
	"log/slog"
)

func SlogAttrSecretKey[T ~string | ~[]byte](secretKey T) slog.Attr {
	var value string

	switch any(secretKey).(type) {
	case string, []byte:
		value = string(secretKey)

	default:
		slog.Warn(fmt.Sprintf("unsupported type secretKey: %T", secretKey))

		value = slog.AnyValue(secretKey).String()
	}

	return slog.String("secret_key", value)
}
