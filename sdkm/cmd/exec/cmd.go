package exec

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	sdkmCmd "github.com/itbasis/tools/sdkm/internal/cmd"
	"github.com/itbasis/tools/sdkm/plugins"
	"github.com/spf13/cobra"
)

func NewExecCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:                "exec",
		Short:              "Execute a command in a plugin",
		DisableFlagParsing: true,
	}

	sdkmCmd.InitFlagRebuildCache(cmd.PersistentFlags())

	plugins.AddPluginsAsSubCommands(
		cmd, func(cmdChild *cobra.Command) {
			cmdChild.Use = itbasisMiddlewareCmd.BuildUse(cmdChild.Use, "{<program>}", "[<args...>]")
			cmdChild.Args = cobra.MinimumNArgs(1)
			cmdChild.ArgAliases = []string{"program"}
			cmdChild.RunE = _runE
		},
	)

	return cmd
}

func _runE(cmd *cobra.Command, args []string) error {
	//nolint:wrapcheck // TODO
	return plugins.GetPluginByID(cmd).
		Exec(
			cmd.Context(),
			sdkmCmd.IsFlagRebuildCache(cmd),
			itbasisMiddlewareOs.Pwd(),
			cmd.InOrStdin(),
			cmd.OutOrStdout(),
			cmd.OutOrStderr(),
			args,
		)
}
