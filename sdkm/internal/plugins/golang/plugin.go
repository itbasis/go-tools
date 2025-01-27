package golang

import (
	"log/slog"
	"path"
	"runtime"

	sdkmCache "github.com/itbasis/go-tools/sdkm/internal/cache"
	pluginGoConsts "github.com/itbasis/go-tools/sdkm/internal/plugins/golang/consts"
	pluginsGoDownloader "github.com/itbasis/go-tools/sdkm/internal/plugins/golang/downloader"
	pluginGoVersions "github.com/itbasis/go-tools/sdkm/internal/plugins/golang/versions"
	sdkmPlugin "github.com/itbasis/go-tools/sdkm/pkg/plugin"
	sdkmSDKVersion "github.com/itbasis/go-tools/sdkm/pkg/sdk-version"
)

type goPlugin struct {
	sdkmPlugin.SDKMPlugin

	basePlugin  sdkmPlugin.BasePlugin
	sdkVersions sdkmSDKVersion.SDKVersions
	downloader  *pluginsGoDownloader.Downloader

	goCacheRootDir string
}

func GetPlugin(basePlugin sdkmPlugin.BasePlugin) (sdkmPlugin.SDKMPlugin, error) {
	downloader := pluginsGoDownloader.NewDownloader(runtime.GOOS, runtime.GOARCH, pluginGoConsts.URLReleases, basePlugin)
	slog.Debug("init downloader: " + downloader.GoString())

	sdkVersions := pluginGoVersions.NewVersions(pluginGoConsts.URLReleases).
		WithCache(sdkmCache.NewCache())
	slog.Debug("init sdkVersions: " + sdkVersions.GoString())

	return &goPlugin{
		basePlugin:  basePlugin,
		downloader:  downloader,
		sdkVersions: sdkVersions,
	}, nil
}

func (receiver *goPlugin) WithVersions(versions sdkmSDKVersion.SDKVersions) sdkmPlugin.SDKMPlugin {
	receiver.sdkVersions = versions

	return receiver
}

func (receiver *goPlugin) enrichSDKVersion(sdkVersion *sdkmSDKVersion.SDKVersion) {
	if sdkVersion == nil {
		return
	}

	sdkVersion.Installed = sdkVersion.Installed ||
		receiver.basePlugin.HasInstalled(pluginGoConsts.PluginID, sdkVersion.ID)
}

func (receiver *goPlugin) getGoCacheDir(version string) string {
	return path.Join(receiver.goCacheRootDir, version)
}
