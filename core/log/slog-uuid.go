package log

import (
	"log/slog"

	"github.com/google/uuid"
)

func SlogAttrUUID(key string, value uuid.UUID) slog.Attr {
	text, err := value.MarshalText()
	if err != nil {
		Panic("failed to marshal UUID value", SlogAttrError(err))
	}

	return slog.String(key, string(text))
}
