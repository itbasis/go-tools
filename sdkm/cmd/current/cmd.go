package current

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	sdkmCmd "github.com/itbasis/tools/sdkm/internal/cmd"
	sdkmPlugins "github.com/itbasis/tools/sdkm/plugins"
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
		sdkVersion, err  = sdkmPlugin.Current(cmd.Context(), flagRebuildCache, itbasisMiddlewareOs.Pwd())
	)

	if err != nil {
		itbasisMiddlewareCmd.Fatal(cmd, err)
	}

	cmd.Println(sdkVersion.PrintWithOptions(false, true, true))
}
