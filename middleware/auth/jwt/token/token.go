package token

import (
	"context"
	"time"

	itbasisMiddlewareAuthModel "github.com/itbasis/tools/middleware/auth/model"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=token.mock.go
type JwtToken interface {
	JwtTokenParser

	CreateAccessToken(context.Context, itbasisMiddlewareAuthModel.SessionUser) (token string, expiredAt *time.Time, err error)
	CreateRefreshToken(context.Context, itbasisMiddlewareAuthModel.SessionUser) (token string, expiredAt *time.Time, err error)
	CreateTokenCustomDuration(context.Context, itbasisMiddlewareAuthModel.SessionUser, time.Duration) (token string, expiredAt *time.Time, err error)
}
