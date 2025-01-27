package token

import (
	"context"

	itbasisMiddlewareAuthModel "github.com/itbasis/go-tools/middleware/auth/model"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=token-parser.mock.go
type JwtTokenParser interface {
	Parse(ctx context.Context, tokenString string) (*itbasisMiddlewareAuthModel.SessionUser, error)
}
