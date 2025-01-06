package base

import (
	"os"
	"path"
	"strings"

	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	sdkmPlugin "github.com/itbasis/tools/sdkm/pkg/plugin"
)

type basePlugin struct {
	sdkDir string
}

func NewBasePlugin() sdkmPlugin.BasePlugin {
	return &basePlugin{sdkDir: path.Join(itbasisMiddlewareOs.UserHomeDir(), "sdk")}
}

func (receiver *basePlugin) WithSDKDir(sdkDir string) sdkmPlugin.BasePlugin {
	receiver.sdkDir = sdkDir

	return receiver
}

func (receiver *basePlugin) GetSDKDir() string {
	return receiver.sdkDir
}

func (receiver *basePlugin) GetSDKVersionDir(pluginName, version string) string {
	sdkFullName := strings.Join([]string{pluginName, version}, "-")

	return path.Join(receiver.GetSDKDir(), sdkFullName)
}

func (receiver *basePlugin) HasInstalled(pluginName, version string) bool {
	sdkFullPath := receiver.GetSDKVersionDir(pluginName, version)

	fi, err := os.Stat(sdkFullPath)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	if !fi.IsDir() {
		panic("sdk path is not a folder: " + sdkFullPath)
	}

	return true
}
