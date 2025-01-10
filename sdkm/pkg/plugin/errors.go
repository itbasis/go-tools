package plugin

import "errors"

var (
	ErrSDKVersionNotFound = errors.New("SDK version not found")

	ErrSDKInstall     = errors.New("SDK install error")
	ErrDownloadFailed = errors.New("download failed")

	ErrExecuteFailed = errors.New("execute failed")
)

type Error struct {
	pluginID ID
	msg      string
}

func NewPluginNotFoundError(pluginID ID) error {
	return &Error{pluginID: pluginID, msg: "plugin not found"}
}

func ErrorInitializePlugin(pluginID ID) error {
	return &Error{pluginID: pluginID, msg: "failed to initialize plugin"}
}

func (err *Error) Error() string {
	return err.msg + ": " + string(err.pluginID)
}
