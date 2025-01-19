package token

import "log/slog"

func SlogAttrClaims(claims *SessionUserClaims) slog.Attr {
	return slog.Any("claims", claims)
}
