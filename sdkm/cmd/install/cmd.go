package install

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	sdkmCmd "github.com/itbasis/tools/sdkm/internal/cmd"
	"github.com/spf13/cobra"
)

const (
	_idxArgPlugin  = 0
	_idxArgVersion = 1
)

func NewInstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:        itbasisMiddlewareCmd.BuildUse("install", sdkmCmd.UseArgRequirePlugins, sdkmCmd.UseArgVersion),
		Short:      "Install the SDK",
		Args:       cobra.MatchAll(cobra.RangeArgs(1, 2), cobra.OnlyValidArgs),
		ArgAliases: []string{sdkmCmd.ArgAliasPlugin, sdkmCmd.ArgAliasVersion},
		RunE:       _runE,
	}
}

func _runE(cmd *cobra.Command, args []string) error {
	var sdkmPlugin = sdkmCmd.GetPluginByName(cmd, args[_idxArgPlugin])

	if len(args) == 1 {
		return sdkmPlugin.Install(cmd.Context(), itbasisMiddlewareOs.Pwd()) //nolint:wrapcheck // _
	}

	return sdkmPlugin.InstallVersion(cmd.Context(), args[_idxArgVersion]) //nolint:wrapcheck // _
}
