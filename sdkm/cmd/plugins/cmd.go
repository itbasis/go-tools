package plugins

import (
	"strings"

	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	"github.com/itbasis/tools/sdkm/plugins"
	"github.com/spf13/cobra"
)

func NewPluginsCommand() *cobra.Command {
	return &cobra.Command{
		Use:    itbasisMiddlewareCmd.BuildUse("plugins"),
		Short:  "List plugins",
		Args:   cobra.NoArgs,
		PreRun: itbasisMiddlewareCmd.LogCommand,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Println("Available plugins: " + strings.Join(plugins.PluginNames, ", "))
		},
	}
}
