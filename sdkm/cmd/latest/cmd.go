package latest

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

func NewLatestCommand() *cobra.Command {
	return &cobra.Command{
		Use:        itbasisMiddlewareCmd.BuildUse("latest", sdkmCmd.UseArgRequirePlugins, sdkmCmd.UseArgVersion),
		Short:      "Show latest stable version",
		ArgAliases: []string{sdkmCmd.ArgAliasPlugin, sdkmCmd.ArgAliasVersion},
		Args:       cobra.MatchAll(cobra.RangeArgs(1, 2), cobra.OnlyValidArgs),
		Run:        _run,
	}
}

func _run(cmd *cobra.Command, args []string) {
	var (
		sdkmPlugin = sdkmCmd.GetPluginByName(cmd, args[_idxArgPlugin])
		sdkVersion sdkmSDKVersion.SDKVersion
		err        error
	)

	if len(args) == 1 {
		sdkVersion, err = sdkmPlugin.LatestVersion(cmd.Context())
	} else {
		sdkVersion, err = sdkmPlugin.LatestVersionByPrefix(cmd.Context(), args[_idxArgVersion])
	}

	if err != nil {
		itbasisMiddlewareCmd.Fatal(cmd, err)
	}

	cmd.Println(sdkVersion.PrintWithOptions(false, true, true))
}
