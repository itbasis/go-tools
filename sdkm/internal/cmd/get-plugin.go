package cmd

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	sdkmPlugin "github.com/itbasis/tools/sdkm/pkg/plugin"
	itbasisSdkmPlugins "github.com/itbasis/tools/sdkm/plugins"
	"github.com/spf13/cobra"
)

func GetPluginByName(cmd *cobra.Command, pluginName string) sdkmPlugin.SDKMPlugin {
	var pluginFunc, ok = itbasisSdkmPlugins.Plugins[pluginName]

	if !ok {
		itbasisMiddlewareCmd.Fatal(cmd, NewErrPluginNotFound(pluginName))
	}

	return pluginFunc()
}
