package token_test

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	itbasisMiddlewareAuthJwtToken "github.com/itbasis/go-tools/middleware/auth/jwt/token"
	itbasisMiddlewareAuthJwtTokenImpl "github.com/itbasis/go-tools/middleware/auth/jwt/token/impl"
	itbasisMiddlewareAuthModel "github.com/itbasis/go-tools/middleware/auth/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/itbasis/go-clock/v2"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"JwtToken", func() {
		var (
			mockClock            = clock.NewMock()
			accessTokenDuration  = time.Second * 11
			refreshTokenDuration = time.Second * 13
			ctx                  context.Context
			jwtTokenConfig       = itbasisMiddlewareAuthJwtToken.Config{
				JwtSecretKey:            "test-key",
				JwtSigningMethod:        "HS512",
				JwtAccessTokenDuration:  accessTokenDuration,
				JwtRefreshTokenDuration: refreshTokenDuration,
			}
			jwtToken itbasisMiddlewareAuthJwtToken.JwtToken
		)

		ginkgo.BeforeEach(
			func() {
				mockClock.Set(time.Now())
				ctx = clock.WithContext(context.Background(), mockClock)

				var token, err = itbasisMiddlewareAuthJwtTokenImpl.NewJwtTokenCustomConfig(ctx, jwtTokenConfig)
				gomega.Expect(err).To(gomega.Succeed())
				jwtToken = token
			},
		)

		ginkgo.DescribeTable(
			"Success creating access token", func(testSessionUser itbasisMiddlewareAuthModel.SessionUser) {
				accessToken, _, err := jwtToken.CreateAccessToken(ctx, testSessionUser)
				gomega.Expect(err).To(gomega.Succeed())
				gomega.Expect(accessToken).NotTo(gomega.BeEmpty())
			},
			ginkgo.Entry(nil, itbasisMiddlewareAuthModel.SessionUser{Username: "test-user"}),
			ginkgo.Entry(
				nil,
				itbasisMiddlewareAuthModel.SessionUser{
					UID:      uuid.MustParse("39910003-f693-44ae-979f-f83714a6d459"),
					Username: "test-user",
					Email:    "test@example.org",
				},
			),
		)

		ginkgo.Describe(
			"Empty fields", func() {
				ginkgo.DescribeTable(
					"Required fields", func(testSessionUser itbasisMiddlewareAuthModel.SessionUser, wantError error) {
						accessToken, expiredAt, err := jwtToken.CreateAccessToken(ctx, testSessionUser)
						gomega.Expect(accessToken).NotTo(gomega.BeEmpty())
						gomega.Expect(expiredAt).To(gomega.HaveValue(gomega.BeTemporally("==", mockClock.Now().Add(accessTokenDuration))))
						gomega.Expect(err).To(gomega.Succeed())
						slog.Info(fmt.Sprintf("accessToken: %s", accessToken))

						gomega.Expect(jwtToken.Parse(ctx, accessToken)).
							Error().To(gomega.MatchError(gomega.ContainSubstring(wantError.Error())))
					},
					ginkgo.Entry(
						"empty email",
						itbasisMiddlewareAuthModel.SessionUser{UID: uuid.MustParse("39910003-f693-44ae-979f-f83714a6d459"), Username: "test-user"},
						itbasisMiddlewareAuthJwtToken.ErrTokenInvalidEmail,
					),
					ginkgo.Entry(
						"empty issuer",
						itbasisMiddlewareAuthModel.SessionUser{UID: uuid.MustParse("39910003-f693-44ae-979f-f83714a6d459")},
						jwt.ErrTokenInvalidIssuer,
					),
				)

				ginkgo.DescribeTable(
					"Optional fields", func(testSessionUser itbasisMiddlewareAuthModel.SessionUser) {
						accessToken, expiredAt, err := jwtToken.CreateAccessToken(ctx, testSessionUser)
						gomega.Expect(accessToken).NotTo(gomega.BeEmpty())
						gomega.Expect(expiredAt).To(gomega.HaveValue(gomega.BeTemporally("==", mockClock.Now().Add(accessTokenDuration))))
						gomega.Expect(err).To(gomega.Succeed())
						slog.Info(fmt.Sprintf("accessToken: %s", accessToken))

						gomega.Expect(jwtToken.Parse(ctx, accessToken)).To(gomega.HaveValue(gomega.BeEquivalentTo(testSessionUser)))
					},
					ginkgo.Entry(
						"empty UID",
						itbasisMiddlewareAuthModel.SessionUser{Username: "test-user", Email: "test@example.org"},
					),
				)
			},
		)
	},
)
