package env

import "errors"

var (
	ErrFailedReadConfiguration = errors.New("failed to read configuration from environment")
)
