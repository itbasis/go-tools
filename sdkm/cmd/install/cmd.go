package install

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	sdkmCmd "github.com/itbasis/tools/sdkm/internal/cmd"
	sdkmPlugins "github.com/itbasis/tools/sdkm/plugins"
	"github.com/spf13/cobra"
)

const (
	_idxArgVersion = 0
)

func NewInstallCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "install",
		Short: "Install the SDK",
	}

	sdkmCmd.InitFlagRebuildCache(cmd.PersistentFlags())

	sdkmPlugins.AddPluginsAsSubCommands(
		cmd, func(cmdChild *cobra.Command) {
			cmdChild.Use = itbasisMiddlewareCmd.BuildUse(cmdChild.Use, sdkmCmd.UseArgVersion)
			cmdChild.ArgAliases = []string{sdkmCmd.ArgAliasVersion}
			cmdChild.Args = cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs)
			cmdChild.RunE = _runE
		},
	)

	return cmd
}

func _runE(cmd *cobra.Command, args []string) error {
	var (
		sdkmPlugin       = sdkmPlugins.GetPluginByID(cmd)
		flagRebuildCache = sdkmCmd.IsFlagRebuildCache(cmd)
	)

	if len(args) == 0 {
		return sdkmPlugin.Install(cmd.Context(), flagRebuildCache, itbasisMiddlewareOs.Pwd()) //nolint:wrapcheck // TODO
	}

	return sdkmPlugin.InstallVersion(cmd.Context(), args[_idxArgVersion]) //nolint:wrapcheck // TODO
}
