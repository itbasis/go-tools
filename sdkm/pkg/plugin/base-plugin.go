package plugin

import (
	"fmt"
	"io"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=base-plugin.mock.go
type BasePlugin interface {
	fmt.GoStringer

	GetSDKDir() string
	GetSDKVersionDir(pluginID ID, version string) string
	HasInstalled(pluginID ID, version string) bool

	Exec(cli string, overrideEnv map[string]string, stdIn io.Reader, stdOut, stdErr io.Writer, args []string) error
}
