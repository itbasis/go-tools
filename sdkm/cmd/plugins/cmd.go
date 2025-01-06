package plugins

import (
	"strings"

	"github.com/itbasis/tools/middleware/cmd"
	itbasisSdkmPlugins "github.com/itbasis/tools/sdkm/plugins"
	"github.com/spf13/cobra"
)

func NewPluginsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   cmd.BuildUse("plugins"),
		Short: "List plugins",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Println("Available plugins: " + strings.Join(itbasisSdkmPlugins.PluginNames, ", "))
		},
	}
}
