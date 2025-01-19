package token

import (
	"time"
)

const (
	DefaultAccessTokenDuration  = time.Minute * 15
	DefaultRefreshTokenDuration = time.Hour * 24 * 14

	DefaultSigningMethod = "HS512"
)

type Config struct {
	JwtSecretKey     string `env:"JWT_SECRET_KEY"`
	JwtSigningMethod string `env:"JWT_SIGNING_METHOD" envDefault:"HS512"`

	JwtAccessTokenDuration time.Duration `env:"JWT_ACCESS_TOKEN_DURATION" envDefault:"15m"`
	// Default: 14 days
	JwtRefreshTokenDuration time.Duration `env:"JWT_REFRESH_TOKEN_DURATION" envDefault:"336h"`
}
