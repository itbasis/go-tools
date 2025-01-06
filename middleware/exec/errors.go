package exec

import (
	"errors"
	"strconv"
)

var (
	ErrFailedExecuteCommand = errors.New("execute command failed")
)

type UnsupportedIncludePrevArgsError struct {
	value IncludePrevArgs
}

func NewUnsupportedIncludePrevArgsError(value IncludePrevArgs) error {
	return &UnsupportedIncludePrevArgsError{value: value}
}

func (err *UnsupportedIncludePrevArgsError) Error() string {
	return "unsupported IncludePrevArgs: " + strconv.Itoa(int(err.value))
}
