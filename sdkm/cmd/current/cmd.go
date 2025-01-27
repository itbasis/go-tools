package current

import (
	itbasisCoreCmd "github.com/itbasis/go-tools/core/cmd"
	itbasisCoreOs "github.com/itbasis/go-tools/core/os"
	sdkmCmd "github.com/itbasis/go-tools/sdkm/internal/cmd"
	sdkmPlugins "github.com/itbasis/go-tools/sdkm/plugins"
	"github.com/spf13/cobra"
)

func NewCurrentCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "current",
		Short: "Display current version",
	}

	sdkmCmd.InitFlagRebuildCache(cmd.PersistentFlags())

	sdkmPlugins.AddPluginsAsSubCommands(
		cmd, func(cmdChild *cobra.Command) {
			cmdChild.Args = cobra.NoArgs
			cmdChild.Run = _run
		},
	)

	return cmd
}

func _run(cmd *cobra.Command, _ []string) {
	var (
		sdkmPlugin       = sdkmPlugins.GetPluginByID(cmd)
		flagRebuildCache = sdkmCmd.IsFlagRebuildCache(cmd)
		sdkVersion, err  = sdkmPlugin.Current(cmd.Context(), flagRebuildCache, itbasisCoreOs.Pwd())
	)

	if err != nil {
		itbasisCoreCmd.Fatal(cmd, err)
	}

	cmd.Println(sdkVersion.PrintWithOptions(false, true, true))
}
