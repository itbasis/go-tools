package log

import (
	"log/slog"
)

func SlogAttrCommand(cmd, workDir string, args []string, customAttrs ...slog.Attr) slog.Attr {
	var attrs = make([]slog.Attr, 0, len(customAttrs)+3) //nolint:mnd // number of attributes added
	attrs = append(
		attrs,
		slog.String("cmd", cmd),
		slog.String("workDir", workDir),
		SlogAttrSliceWithSeparator("args", " ", args),
	)
	attrs = append(attrs, customAttrs...)

	return slog.Attr{Key: "exec command", Value: slog.GroupValue(attrs...)}
}
