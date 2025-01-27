package plugins

import (
	"log/slog"

	itbasisCoreCmd "github.com/itbasis/go-tools/core/cmd"
	sdkmCmd "github.com/itbasis/go-tools/sdkm/internal/cmd"
	pluginBase "github.com/itbasis/go-tools/sdkm/internal/plugins/base"
	sdkmPlugin "github.com/itbasis/go-tools/sdkm/pkg/plugin"
	"github.com/spf13/cobra"
)

type ExtCobraPluginInit interface {
	InitProcessCommandFlags(cmd *cobra.Command)
}

func GetPluginByID(cmd *cobra.Command) sdkmPlugin.SDKMPlugin {
	var id, exist = cmd.Annotations[AnnotationPluginID]
	if !exist {
		slog.Error("Plugin ID annotation not found")

		return nil
	}

	slog.Debug("getting plugin: " + id)

	var (
		pluginID       = sdkmPlugin.ID(id)
		pluginMeta, ok = _plugins[pluginID]
	)

	if !ok {
		itbasisCoreCmd.Fatal(cmd, sdkmPlugin.NewPluginNotFoundError(pluginID))
	}

	slog.Debug("found plugin: " + string(pluginID))

	sdkPlugin, err := pluginMeta.GetPluginFunc(_initBasePlugin(cmd))

	if pluginExtInit, ok := sdkPlugin.(ExtCobraPluginInit); ok {
		pluginExtInit.InitProcessCommandFlags(cmd)
	}

	itbasisCoreCmd.RequireNoError(cmd, err)

	return sdkPlugin
}

func _initBasePlugin(cmd *cobra.Command) sdkmPlugin.BasePlugin {
	basePlugin, errBasePlugin := pluginBase.NewBasePlugin(
		pluginBase.WithCustomSdkDir(sdkmCmd.GetSdkRootDir(cmd)),
	)
	itbasisCoreCmd.RequireNoError(cmd, errBasePlugin)

	slog.Debug("init base plugin: " + basePlugin.GoString())

	return basePlugin
}
