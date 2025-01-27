package impl

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"

	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
	itbasisMiddlewareAuthJwtToken "github.com/itbasis/go-tools/middleware/auth/jwt/token"
	itbasisMiddlewareAuthModel "github.com/itbasis/go-tools/middleware/auth/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Parse TODO Adding Firebase parsing
func (receiver *_jwtToken) Parse(ctx context.Context, rawToken string) (*itbasisMiddlewareAuthModel.SessionUser, error) {
	return receiver.parseWithSecretKey(ctx, rawToken)
}

func (receiver *_jwtToken) parseWithSecretKey(ctx context.Context, rawToken string) (*itbasisMiddlewareAuthModel.SessionUser, error) {
	logger := itbasisCoreLog.Logger(ctx)

	logger.Debug("receive token", itbasisMiddlewareAuthJwtToken.SlogAttrToken(rawToken))

	token, err := jwt.ParseWithClaims(
		rawToken, &itbasisMiddlewareAuthJwtToken.SessionUserClaims{}, func(token *jwt.Token) (interface{}, error) {
			method, ok := token.Method.(*jwt.SigningMethodHMAC)

			if ok {
				logger.Debug(
					"found method",
					slog.Any("method", method),
					itbasisMiddlewareAuthJwtToken.SlogAttrSecretKey(receiver.signSecretKey),
				)

				return receiver.signSecretKey, nil
			}

			alg := token.Header["alg"]

			msg := fmt.Sprintf("unsupported token signing algorithm: %v", alg)
			errWrap := errors.Wrap(itbasisMiddlewareAuthJwtToken.ErrUnsupportedSigningMethod, msg)

			logger.Error(msg, itbasisCoreLog.SlogAttrError(itbasisMiddlewareAuthJwtToken.ErrUnsupportedSigningMethod))

			return nil, errWrap
		},
	)
	if err != nil {
		errWrap := errors.Wrap(err, itbasisMiddlewareAuthJwtToken.ErrParsingClaims.Error())
		logger.Error(errWrap.Error(), itbasisCoreLog.SlogAttrError(errWrap))

		return nil, errWrap
	}

	// TODO check - this seems like a redundant check
	if !token.Valid {
		logger.Error(
			itbasisMiddlewareAuthJwtToken.ErrTokenInvalid.Error(),
			itbasisCoreLog.SlogAttrError(itbasisMiddlewareAuthJwtToken.ErrTokenInvalid),
		)

		return nil, itbasisMiddlewareAuthJwtToken.ErrTokenInvalid
	}

	claims, ok := token.Claims.(*itbasisMiddlewareAuthJwtToken.SessionUserClaims)
	if !ok {
		errWrap := fmt.Errorf("%w: found type: %T", itbasisMiddlewareAuthJwtToken.ErrUnsupportedType, token.Claims)
		logger.Error(errWrap.Error(), itbasisCoreLog.SlogAttrError(errWrap))

		return nil, errWrap
	}

	return receiver.enrichSessionUser(logger, claims)
}

func (receiver *_jwtToken) enrichSessionUser(logger *slog.Logger, claims *itbasisMiddlewareAuthJwtToken.SessionUserClaims) (
	*itbasisMiddlewareAuthModel.SessionUser, error,
) {
	logger = logger.With(itbasisMiddlewareAuthJwtToken.SlogAttrClaims(claims))
	logger.Debug("receive claims for enrich")

	sessionUser := &itbasisMiddlewareAuthModel.SessionUser{}

	if claims.UID != uuid.Nil {
		slog.Debug("UID is valid")

		sessionUser.UID = claims.UID
	} else {
		logger.Warn("UID is empty", itbasisCoreLog.SlogAttrError(itbasisMiddlewareAuthJwtToken.ErrTokenInvalidUID))
	}

	if len(claims.Issuer) > 0 {
		slog.Debug("Issuer is valid")

		sessionUser.Username = claims.Issuer
	} else {
		logger.Error(ErrMsgIsEmpty, itbasisCoreLog.SlogAttrError(jwt.ErrTokenInvalidIssuer))

		return nil, errors.Wrap(jwt.ErrTokenInvalidIssuer, ErrMsgIsEmpty)
	}

	if len(claims.Email) > 0 {
		slog.Debug("Email is valid")

		sessionUser.Email = claims.Email
	} else {
		logger.Error(ErrMsgIsEmpty, itbasisCoreLog.SlogAttrError(itbasisMiddlewareAuthJwtToken.ErrTokenInvalidEmail))

		return nil, errors.Wrap(itbasisMiddlewareAuthJwtToken.ErrTokenInvalidEmail, ErrMsgIsEmpty)
	}

	// TODO sessionUser.hasGuest
	// TODO sessionUser.UID

	return sessionUser, nil
}
