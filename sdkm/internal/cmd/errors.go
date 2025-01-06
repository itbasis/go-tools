package cmd

import (
	"errors"
	"fmt"
)

var (
	ErrPluginNotFound = errors.New("plugin not found")
)

func NewErrPluginNotFound(pluginName string) error {
	return fmt.Errorf("%w: %s", ErrPluginNotFound, pluginName)
}
