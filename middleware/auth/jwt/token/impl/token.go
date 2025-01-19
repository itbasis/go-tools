package impl

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"

	itbasisCoreEnv "github.com/itbasis/tools/core/env"
	itbasisCoreLog "github.com/itbasis/tools/core/log"
	itbasisMiddlewareAuthJwtToken "github.com/itbasis/tools/middleware/auth/jwt/token"
)

type _jwtToken struct {
	itbasisMiddlewareAuthJwtToken.JwtToken

	signSecretKey []byte
	signMethod    jwt.SigningMethod

	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewJwtToken(ctx context.Context) (itbasisMiddlewareAuthJwtToken.JwtToken, error) {
	config := itbasisMiddlewareAuthJwtToken.Config{}

	if err := itbasisCoreEnv.ReadEnvConfig(&config); err != nil {
		return nil, err //nolint:wrapcheck // TODO
	}

	itbasisCoreLog.Logger(ctx).Debug("jwtToken", slog.Any("config", config))

	return NewJwtTokenCustomConfig(ctx, config)
}

func NewJwtTokenCustomConfig(ctx context.Context, config itbasisMiddlewareAuthJwtToken.Config) (itbasisMiddlewareAuthJwtToken.JwtToken, error) {
	obj := &_jwtToken{
		accessTokenDuration:  config.JwtAccessTokenDuration,
		refreshTokenDuration: config.JwtRefreshTokenDuration,
	}

	if len(config.JwtSecretKey) > 0 {
		signingMethod := jwt.GetSigningMethod(config.JwtSigningMethod)
		itbasisCoreLog.Logger(ctx).Info(fmt.Sprintf("Using signing method: %++v", signingMethod))

		if signingMethod == jwt.SigningMethodNone {
			return nil, fmt.Errorf("%w: %s", jwt.ErrInvalidKeyType, config.JwtSigningMethod)
		}

		obj.signMethod = signingMethod
		obj.signSecretKey = []byte(config.JwtSecretKey)
	}

	return obj, nil
}
