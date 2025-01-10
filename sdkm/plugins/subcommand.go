package plugins

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	sdkmPlugin "github.com/itbasis/tools/sdkm/pkg/plugin"
	"github.com/spf13/cobra"
)

const (
	AnnotationPluginID = "pluginID"
)

func AddPluginsAsSubCommands(cmdParent *cobra.Command, funcEnrichCommand sdkmPlugin.ExtCommandInit) {
	for _, pluginMeta := range _plugins {
		cmdChild := &cobra.Command{
			Use:    string(pluginMeta.ID),
			PreRun: itbasisMiddlewareCmd.LogCommand,
		}
		cmdChild.Annotations = map[string]string{AnnotationPluginID: string(pluginMeta.ID)}

		funcEnrichCommand(cmdChild)

		if pluginMeta.ExtCommandInit != nil {
			pluginMeta.ExtCommandInit(cmdChild)
		}

		cmdParent.AddCommand(cmdChild)
	}
}
