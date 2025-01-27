package install

import (
	itbasisCoreCmd "github.com/itbasis/go-tools/core/cmd"
	itbasisCoreOs "github.com/itbasis/go-tools/core/os"
	sdkmCmd "github.com/itbasis/go-tools/sdkm/internal/cmd"
	sdkmPlugins "github.com/itbasis/go-tools/sdkm/plugins"
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
			cmdChild.Use = itbasisCoreCmd.BuildUse(cmdChild.Use, sdkmCmd.UseArgVersion)
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
		return sdkmPlugin.Install(cmd.Context(), flagRebuildCache, itbasisCoreOs.Pwd()) //nolint:wrapcheck // TODO
	}

	return sdkmPlugin.InstallVersion(cmd.Context(), args[_idxArgVersion]) //nolint:wrapcheck // TODO
}
