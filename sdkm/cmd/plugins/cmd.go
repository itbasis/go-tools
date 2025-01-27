package plugins

import (
	"strings"

	itbasisCoreCmd "github.com/itbasis/go-tools/core/cmd"
	"github.com/itbasis/go-tools/sdkm/plugins"
	"github.com/spf13/cobra"
)

func NewPluginsCommand() *cobra.Command {
	return &cobra.Command{
		Use:    itbasisCoreCmd.BuildUse("plugins"),
		Short:  "List plugins",
		Args:   cobra.NoArgs,
		PreRun: itbasisCoreCmd.LogCommand,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Println("Available plugins: " + strings.Join(plugins.PluginNames, ", "))
		},
	}
}
