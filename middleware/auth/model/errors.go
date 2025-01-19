package model

import "github.com/pkg/errors"

var (
	MsgAuthHeaderFound    = "authentication token was found"
	MsgAuthHeaderNotFound = "authentication token was not found"

	ErrSessionWithoutAuth = errors.New("session without authentication")

	ErrSessionUserNotFoundOrInvalid = errors.New("session user not found or invalid")
	ErrSessionUserInvalid           = errors.New("session contains an invalid user object")
)
