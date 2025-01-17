package base

import (
	"os"
	"path"

	itbasisCoreOption "github.com/itbasis/tools/core/option"
	sdkmPlugin "github.com/itbasis/tools/sdkm/pkg/plugin"
)

type basePlugin struct {
	sdkDir string
}

func NewBasePlugin(opts ...itbasisCoreOption.Option[basePlugin]) (sdkmPlugin.BasePlugin, error) {
	cmp := &basePlugin{}

	if err := itbasisCoreOption.ApplyOptions(
		cmp, opts, map[itbasisCoreOption.Key]itbasisCoreOption.LazyOptionFunc[basePlugin]{
			_optionSdkDirKey: WithDefaultSdkDir,
		},
	); err != nil {
		return nil, err //nolint:wrapcheck // TODO
	}

	return cmp, nil
}

func (receiver *basePlugin) GoString() string {
	if receiver == nil {
		return "<nil>"
	}

	return "basePlugin{sdkDir: " + receiver.sdkDir + "}"
}

func (receiver *basePlugin) GetSDKDir() string {
	return receiver.sdkDir
}

func (receiver *basePlugin) GetSDKVersionDir(pluginID sdkmPlugin.ID, version string) string {
	return path.Join(receiver.GetSDKDir(), string(pluginID), version)
}

func (receiver *basePlugin) HasInstalled(pluginID sdkmPlugin.ID, version string) bool {
	sdkFullPath := receiver.GetSDKVersionDir(pluginID, version)

	fi, err := os.Stat(sdkFullPath)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	if !fi.IsDir() {
		panic("sdk path is not a folder: " + sdkFullPath)
	}

	return true
}
