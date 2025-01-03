package log

import (
	"log/slog"
)

func SlogAttrFilePath(filePath string) slog.Attr {
	return slog.String("filePath", filePath)
}
