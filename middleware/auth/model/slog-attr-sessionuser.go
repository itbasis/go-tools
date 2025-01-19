package model

import (
	"log/slog"

	itbasisCoreLog "github.com/itbasis/tools/core/log"
)

func SlogAttrSessionUser(sessionUser SessionUser) slog.Attr {
	return slog.Group(
		"session_user",
		itbasisCoreLog.SlogAttrUUID("uid", sessionUser.UID),
		slog.String("username", sessionUser.Username),
		slog.String("email", sessionUser.Email),
		slog.Bool("guest", sessionUser.HasGuest),
	)
}

func SlogAttrSessionUserUID(sessionUser SessionUser) slog.Attr {
	return itbasisCoreLog.SlogAttrUUID("session_user_uid", sessionUser.UID)
}
