package token

import "log/slog"

func SlogAttrToken(token string) slog.Attr {
	return slog.String("token", token)
}
