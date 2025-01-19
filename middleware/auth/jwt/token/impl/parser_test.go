package impl_test

import (
	"context"

	itbasisMiddlewareAuthJwtToken "github.com/itbasis/tools/middleware/auth/jwt/token"
	itbasisMiddlewareAuthJwtTokenImpl "github.com/itbasis/tools/middleware/auth/jwt/token/impl"

	"github.com/dchest/uniuri"
	"github.com/itbasis/go-clock/v2"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"Parsing token", func() {
		var (
			jwtToken  itbasisMiddlewareAuthJwtToken.JwtToken
			mockClock = clock.NewMock()
			secretKey = "test-key"
		)

		ginkgo.BeforeEach(
			func() {
				var err error

				jwtToken, err = itbasisMiddlewareAuthJwtTokenImpl.NewJwtTokenCustomConfig(
					clock.WithContext(context.Background(), mockClock),
					itbasisMiddlewareAuthJwtToken.Config{
						JwtSecretKey:     secretKey,
						JwtSigningMethod: itbasisMiddlewareAuthJwtToken.DefaultSigningMethod,
					},
				)
				gomega.Expect(err).To(gomega.Succeed())
			},
		)

		ginkgo.DescribeTable(
			"Invalid token", func(testToken string, expectErr string) {
				gomega.
					Expect(jwtToken.Parse(context.Background(), testToken)).
					Error().To(gomega.MatchError(gomega.ContainSubstring(expectErr)))
			},
			ginkgo.Entry("empty token", "", itbasisMiddlewareAuthJwtToken.ErrParsingClaims.Error()),
			ginkgo.Entry("random string", uniuri.New(), itbasisMiddlewareAuthJwtToken.ErrParsingClaims.Error()),
			ginkgo.Entry("invalid token", "1.2.3", itbasisMiddlewareAuthJwtToken.ErrParsingClaims.Error()),
		)

	},
)
