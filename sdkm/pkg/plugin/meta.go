package plugin

import "github.com/spf13/cobra"

type ID string

type GetPluginFunc func(basePlugin BasePlugin) (SDKMPlugin, error)

type ExtCommandInit func(cmd *cobra.Command)

type MetaInfo struct {
	ID             ID
	GetPluginFunc  GetPluginFunc
	ExtCommandInit ExtCommandInit
}
