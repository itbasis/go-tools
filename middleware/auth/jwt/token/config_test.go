package token_test

import (
	"os"
	"time"

	itbasisCoreEnv "github.com/itbasis/go-tools/core/env"
	itbasisMiddlewareAuthJwtToken "github.com/itbasis/go-tools/middleware/auth/jwt/token"

	"github.com/google/uuid"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

const (
	envJwtSecretKey            = "JWT_SECRET_KEY"
	envJwtSigningMethod        = "JWT_SIGNING_METHOD"
	envJwtAccessTokenDuration  = "JWT_ACCESS_TOKEN_DURATION"
	envJwtRefreshTokenDuration = "JWT_REFRESH_TOKEN_DURATION"
)

var _ = ginkgo.Describe(
	"JWT Config", func() {
		var config itbasisMiddlewareAuthJwtToken.Config

		ginkgo.AfterEach(
			func() {
				gomega.Expect(os.Unsetenv(envJwtSecretKey)).To(gomega.Succeed())
				gomega.Expect(os.Unsetenv(envJwtSigningMethod)).To(gomega.Succeed())
				gomega.Expect(os.Unsetenv(envJwtAccessTokenDuration)).To(gomega.Succeed())
				gomega.Expect(os.Unsetenv(envJwtRefreshTokenDuration)).To(gomega.Succeed())

				config = itbasisMiddlewareAuthJwtToken.Config{}
			},
		)

		ginkgo.It(
			"test empty config", func() {
				gomega.Expect(itbasisCoreEnv.ReadEnvConfig(&config)).To(gomega.Succeed())
				gomega.Expect(config.JwtSecretKey).To(gomega.BeEmpty())
				gomega.Expect(config.JwtSigningMethod).To(gomega.Equal("HS512"))
				gomega.Expect(config.JwtAccessTokenDuration).To(gomega.Equal(itbasisMiddlewareAuthJwtToken.DefaultAccessTokenDuration))
				gomega.Expect(config.JwtRefreshTokenDuration).To(gomega.Equal(itbasisMiddlewareAuthJwtToken.DefaultRefreshTokenDuration))
			},
		)

		ginkgo.It(
			"test custom config", func() {
				testKey := uuid.New()

				gomega.Expect(os.Setenv(envJwtSecretKey, testKey.String())).To(gomega.Succeed())
				gomega.Expect(os.Setenv(envJwtSigningMethod, "HS256")).To(gomega.Succeed())
				gomega.Expect(os.Setenv(envJwtAccessTokenDuration, "20s")).To(gomega.Succeed())
				gomega.Expect(os.Setenv(envJwtRefreshTokenDuration, "30s")).To(gomega.Succeed())

				gomega.Expect(itbasisCoreEnv.ReadEnvConfig(&config)).To(gomega.Succeed())
				gomega.Expect(config.JwtSecretKey).To(gomega.Equal(testKey.String()))
				gomega.Expect(config.JwtSigningMethod).To(gomega.Equal("HS256"))
				gomega.Expect(config.JwtAccessTokenDuration).To(gomega.Equal(time.Second * 20))
				gomega.Expect(config.JwtRefreshTokenDuration).To(gomega.Equal(time.Second * 30))
			},
		)
	},
)
