package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type SessionUserClaims struct {
	UID   uuid.UUID `json:"uid,omitempty"`
	Email string    `json:"email,omitempty"`

	jwt.RegisteredClaims
}
