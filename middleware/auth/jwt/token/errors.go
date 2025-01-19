package token

import "errors"

var (
	ErrParsingClaims            = errors.New("failed to parse Claims")
	ErrUnsupportedSigningMethod = errors.New("unsupported signing method")
	ErrTokenInvalid             = errors.New("the token has an invalid")
	ErrUnsupportedType          = errors.New("unsupported claims type")

	ErrTokenInvalidUID   = errors.New("token has invalid UID")
	ErrTokenInvalidEmail = errors.New("token has invalid email")
)
