package impl

import (
	"context"
	"time"

	itbasisMiddlewareAuthModel "github.com/itbasis/go-tools/middleware/auth/model"
)

func (receiver *_jwtToken) CreateAccessToken(
	ctx context.Context,
	sessionUser itbasisMiddlewareAuthModel.SessionUser,
) (string, *time.Time, error) {
	return receiver.CreateTokenCustomDuration(ctx, sessionUser, receiver.accessTokenDuration)
}
