package golang

import (
	"path"

	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	sdkmCache "github.com/itbasis/tools/sdkm/internal/cache"
	cacheFileStorage "github.com/itbasis/tools/sdkm/internal/cache/storage/file-storage"
	sdkmCmd "github.com/itbasis/tools/sdkm/internal/cmd"
	pluginGoConsts "github.com/itbasis/tools/sdkm/internal/plugins/golang/consts"
	"github.com/spf13/cobra"
)

var (
	CmdFlagRebuildCache bool
)

func CmdExtPlugin(cmd *cobra.Command) {
	sdkmCmd.InitFlagCacheRootDir(cmd.Flags())

	cmd.Flags().Bool("rebuild-cache", false, "rebuild cache ")

}

func (receiver *goPlugin) InitProcessCommandFlags(cmd *cobra.Command) {
	if receiver.goCacheRootDir = sdkmCmd.GetCacheRootDir(cmd); receiver.goCacheRootDir == "" {
		receiver.goCacheRootDir = path.Join(itbasisMiddlewareOs.UserHomeDir(), ".cache", string(pluginGoConsts.PluginID))
	}

	receiver.sdkVersions = receiver.sdkVersions.WithCache(
		sdkmCache.NewCache().
			WithExternalStore(cacheFileStorage.NewFileCacheStorage(pluginGoConsts.PluginID)),
	)
}
