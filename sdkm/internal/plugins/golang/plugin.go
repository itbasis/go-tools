package golang

import (
	"path"
	"runtime"

	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	sdkmCache "github.com/itbasis/tools/sdkm/internal/cache"
	cacheFileStorage "github.com/itbasis/tools/sdkm/internal/cache/storage/file-storage"
	pluginBase "github.com/itbasis/tools/sdkm/internal/plugins/base"
	pluginGoConsts "github.com/itbasis/tools/sdkm/internal/plugins/golang/consts"
	pluginsGoDownloader "github.com/itbasis/tools/sdkm/internal/plugins/golang/downloader"
	pluginGoVersions "github.com/itbasis/tools/sdkm/internal/plugins/golang/versions"
	sdkmPlugin "github.com/itbasis/tools/sdkm/pkg/plugin"
	sdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
)

type goPlugin struct {
	sdkmPlugin.SDKMPlugin

	basePlugin  sdkmPlugin.BasePlugin
	sdkVersions sdkmSDKVersion.SDKVersions
	downloader  *pluginsGoDownloader.Downloader
}

func GetPlugin() sdkmPlugin.SDKMPlugin {
	basePlugin := pluginBase.NewBasePlugin()
	downloader := pluginsGoDownloader.NewDownloader(
		runtime.GOOS, runtime.GOARCH, pluginGoConsts.URLReleases, basePlugin.GetSDKDir(),
	)
	cache := sdkmCache.NewCache().
		WithExternalStore(cacheFileStorage.NewFileCacheStorage(pluginGoConsts.PluginName))

	sdkVersions := pluginGoVersions.NewVersions(pluginGoConsts.URLReleases).
		WithCache(cache)

	return &goPlugin{
		basePlugin:  basePlugin,
		downloader:  downloader,
		sdkVersions: sdkVersions,
	}
}

func (receiver *goPlugin) WithBasePlugin(basePlugin sdkmPlugin.BasePlugin) sdkmPlugin.SDKMPlugin {
	receiver.basePlugin = basePlugin
	receiver.downloader = pluginsGoDownloader.NewDownloader(
		runtime.GOOS, runtime.GOARCH, pluginGoConsts.URLReleases, basePlugin.GetSDKDir(),
	)

	return receiver
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
		receiver.basePlugin.HasInstalled(pluginGoConsts.PluginName, sdkVersion.ID)
}

func (receiver *goPlugin) getGoCacheDir(version string) string {
	return path.Join(itbasisMiddlewareOs.UserHomeDir(), pluginGoConsts.PluginName, version)
}
