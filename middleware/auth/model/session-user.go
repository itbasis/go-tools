package model

import (
	"github.com/google/uuid"
)

type SessionUser struct {
	UID uuid.UUID

	Username string
	Email    string

	HasGuest bool
}
