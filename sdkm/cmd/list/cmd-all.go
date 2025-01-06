package list

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	sdkmCmd "github.com/itbasis/tools/sdkm/internal/cmd"
	sdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
	"github.com/spf13/cobra"
)

const (
	_idxArgPlugin  = 0
	_idxArgVersion = 1
)

func newListAllCommand() *cobra.Command {
	return &cobra.Command{
		Use:        itbasisMiddlewareCmd.BuildUse("all", sdkmCmd.UseArgRequirePlugins, sdkmCmd.UseArgVersion),
		Short:      "List all versions",
		ArgAliases: []string{sdkmCmd.ArgAliasPlugin, sdkmCmd.ArgAliasVersion},
		Args:       cobra.MatchAll(cobra.RangeArgs(1, 2), cobra.OnlyValidArgs),
		Run:        _run,
	}
}

func _run(cmd *cobra.Command, args []string) {
	var (
		sdkmPlugin  = sdkmCmd.GetPluginByName(cmd, args[_idxArgPlugin])
		sdkVersions []sdkmSDKVersion.SDKVersion
		err         error
	)

	if len(args) == 1 {
		sdkVersions, err = sdkmPlugin.ListAllVersions(cmd.Context())
	} else {
		sdkVersions, err = sdkmPlugin.ListAllVersionsByPrefix(cmd.Context(), args[_idxArgVersion])
	}

	if err != nil {
		itbasisMiddlewareCmd.Fatal(cmd, err)
	}

	for _, sdkVersion := range sdkVersions {
		// TODO code smell
		cmd.Println(sdkVersion)
	}
}
