package golang

import (
	"path"

	itbasisCoreOs "github.com/itbasis/tools/core/os"
	sdkmCache "github.com/itbasis/tools/sdkm/internal/cache"
	cacheFileStorage "github.com/itbasis/tools/sdkm/internal/cache/storage/file-storage"
	sdkmCmd "github.com/itbasis/tools/sdkm/internal/cmd"
	pluginGoConsts "github.com/itbasis/tools/sdkm/internal/plugins/golang/consts"
	"github.com/spf13/cobra"
)

func CmdExtPlugin(cmd *cobra.Command) {
	flags := cmd.Flags()

	sdkmCmd.InitFlagCacheRootDir(flags)
	sdkmCmd.InitFlagRebuildCache(flags)
}

func (receiver *goPlugin) InitProcessCommandFlags(cmd *cobra.Command) {
	if receiver.goCacheRootDir = sdkmCmd.GetCacheRootDir(cmd); receiver.goCacheRootDir == "" {
		receiver.goCacheRootDir = path.Join(itbasisCoreOs.UserHomeDir(), ".cache", string(pluginGoConsts.PluginID))
	}

	receiver.sdkVersions = receiver.sdkVersions.WithCache(
		sdkmCache.NewCache().
			WithExternalStore(cacheFileStorage.NewFileCacheStorage(pluginGoConsts.PluginID)),
	)
}
