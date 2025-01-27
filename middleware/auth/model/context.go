package model

import (
	"context"

	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
)

type _sessionUserKey struct{}

func WithSessionUser(ctx context.Context, sessionUser SessionUser) context.Context {
	return context.WithValue(ctx, _sessionUserKey{}, sessionUser)
}

func GetSessionUser(ctx context.Context) (SessionUser, error) {
	if ctx == nil {
		panic("nil context")
	}

	logger := itbasisCoreLog.Logger(ctx)

	sessionUser, ok := ctx.Value(_sessionUserKey{}).(SessionUser)
	if !ok {
		logger.Debug("session user not found in context")

		return SessionUser{}, ErrSessionUserNotFoundOrInvalid
	}

	logger.Debug("session user found in context", SlogAttrSessionUser(sessionUser))

	return sessionUser, nil
}
