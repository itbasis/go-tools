package option

import "log/slog"

func _slogAttrOptionKey(key Key) slog.Attr {
	return slog.String("option_key", string(key))
}
