package log

import "log/slog"

func SlogAttrCommand(command string, args ...string) slog.Attr {
	return slog.Group(
		"command",
		slog.String("cmd", command),
		SlogAttrSliceWithSeparator("args", " ", args),
	)
}
