package exec

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	sdkmCmd "github.com/itbasis/tools/sdkm/internal/cmd"
	"github.com/spf13/cobra"
)

const (
	_idxArgPlugin = 0
)

func NewExecCommand() *cobra.Command {
	return &cobra.Command{
		Use:                itbasisMiddlewareCmd.BuildUse("exec", sdkmCmd.UseArgRequirePlugins, "{<program>}", "[<args...>]"),
		Short:              "Execute a command in a plugin",
		DisableFlagParsing: true,
		Args:               cobra.MatchAll(cobra.MinimumNArgs(2)),
		ArgAliases:         []string{sdkmCmd.ArgAliasPlugin, "program"},
		RunE:               _runE,
	}
}

func _runE(cmd *cobra.Command, args []string) error {
	//nolint:wrapcheck // _
	return sdkmCmd.GetPluginByName(cmd, args[_idxArgPlugin]).
		Exec(
			cmd.Context(),
			itbasisMiddlewareOs.Pwd(),
			cmd.InOrStdin(),
			cmd.OutOrStdout(),
			cmd.OutOrStderr(),
			args[1:],
		)
}
