package impl

import (
	"context"
	"time"

	itbasisMiddlewareAuthModel "github.com/itbasis/tools/middleware/auth/model"
)

func (receiver *_jwtToken) CreateRefreshToken(
	ctx context.Context,
	sessionUser itbasisMiddlewareAuthModel.SessionUser,
) (string, *time.Time, error) {
	return receiver.CreateTokenCustomDuration(ctx, sessionUser, receiver.refreshTokenDuration)
}
