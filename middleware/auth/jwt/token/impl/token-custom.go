package impl

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	itbasisCoreLog "github.com/itbasis/tools/core/log"
	itbasisMiddlewareAuthJwtToken "github.com/itbasis/tools/middleware/auth/jwt/token"
	itbasisMiddlewareAuthModel "github.com/itbasis/tools/middleware/auth/model"

	"github.com/itbasis/go-clock/v2"
	"github.com/pkg/errors"
)

func (receiver *_jwtToken) CreateTokenCustomDuration(
	ctx context.Context,
	sessionUser itbasisMiddlewareAuthModel.SessionUser,
	expiredAtDuration time.Duration,
) (string, *time.Time, error) {
	logger := itbasisCoreLog.Logger(ctx)

	slog.Debug("creating token...", itbasisMiddlewareAuthModel.SlogAttrSessionUser(sessionUser))

	now := clock.FromContext(ctx).Now()
	expiredAt := now.Add(expiredAtDuration)
	logger.Debug(fmt.Sprintf("expiredAt: %s", expiredAt))

	claims := itbasisMiddlewareAuthJwtToken.SessionUserClaims{
		UID: sessionUser.UID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    sessionUser.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	}
	if len(sessionUser.Email) > 0 {
		claims.Email = sessionUser.Email
	}

	logger.Debug(fmt.Sprintf("claims: %++v", claims))

	token := jwt.NewWithClaims(receiver.signMethod, claims)

	signedString, err := token.SignedString(receiver.signSecretKey)
	if err != nil {
		errWrap := errors.Wrap(ErrCreatingToken, err.Error())
		logger.Error("token signing error", itbasisCoreLog.SlogAttrError(errWrap))

		return "", nil, errWrap
	}

	logger.Debug(fmt.Sprintf("signed token: %s", signedString))

	return signedString, &expiredAt, nil
}
