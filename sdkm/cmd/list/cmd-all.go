package list

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	sdkmCmd "github.com/itbasis/tools/sdkm/internal/cmd"
	sdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
	sdkmPlugins "github.com/itbasis/tools/sdkm/plugins"
	"github.com/spf13/cobra"
)

const (
	_idxArgVersion = 0
)

func newListAllCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "all",
		Short: "List all versions",
	}

	sdkmPlugins.AddPluginsAsSubCommands(
		cmd, func(cmdChild *cobra.Command) {
			cmdChild.Use = itbasisMiddlewareCmd.BuildUse(cmdChild.Use, sdkmCmd.UseArgVersion)
			cmdChild.ArgAliases = []string{sdkmCmd.ArgAliasVersion}
			cmdChild.Args = cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs)
			cmdChild.Run = _run
		},
	)

	return cmd
}

func _run(cmd *cobra.Command, args []string) {
	var (
		plugin      = sdkmPlugins.GetPluginByID(cmd)
		sdkVersions []sdkmSDKVersion.SDKVersion
		err         error
	)

	var flagRebuildCache = sdkmCmd.IsFlagRebuildCache(cmd)

	if len(args) == 0 {
		sdkVersions, err = plugin.ListAllVersions(cmd.Context(), flagRebuildCache)
	} else {
		sdkVersions, err = plugin.ListAllVersionsByPrefix(cmd.Context(), flagRebuildCache, args[_idxArgVersion])
	}

	if err != nil {
		itbasisMiddlewareCmd.Fatal(cmd, err)
	}

	for _, sdkVersion := range sdkVersions {
		// TODO code smell
		cmd.Println(sdkVersion)
	}
}
