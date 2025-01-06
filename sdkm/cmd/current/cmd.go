package current

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	sdkmCmd "github.com/itbasis/tools/sdkm/internal/cmd"
	"github.com/spf13/cobra"
)

const (
	_idxArgPlugin = 0
)

func NewCurrentCommand() *cobra.Command {
	return &cobra.Command{
		Use:        itbasisMiddlewareCmd.BuildUse("current"),
		Short:      "Display current version",
		ArgAliases: []string{sdkmCmd.ArgAliasPlugin},
		Args:       cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run:        _run,
	}
}

func _run(cmd *cobra.Command, args []string) {
	var (
		sdkmPlugin      = sdkmCmd.GetPluginByName(cmd, args[_idxArgPlugin])
		sdkVersion, err = sdkmPlugin.Current(cmd.Context(), itbasisMiddlewareOs.Pwd())
	)

	if err != nil {
		itbasisMiddlewareCmd.Fatal(cmd, err)
	}

	cmd.Println(sdkVersion.PrintWithOptions(false, true, true))
}
