package current

import (
	"strings"

	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	"github.com/itbasis/tools/sdkm/plugins"
	"github.com/spf13/cobra"
)

func NewCurrentCommand() *cobra.Command {
	return &cobra.Command{
		Use:        itbasisMiddlewareCmd.BuildUse("current", "{"+strings.Join(plugins.PluginNames, "|")+"}"),
		Short:      "Display current version",
		ArgAliases: []string{"plugin"},
		Args:       cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		PreRun:     itbasisMiddlewareCmd.LogCommand,
		Run:        _run,
	}
}

func _run(cmd *cobra.Command, _ []string) {
	var (
		sdkmPlugin      = plugins.GetPluginByID(cmd)
		sdkVersion, err = sdkmPlugin.Current(cmd.Context(), itbasisMiddlewareOs.Pwd())
	)

	if err != nil {
		itbasisMiddlewareCmd.Fatal(cmd, err)
	}

	cmd.Println(sdkVersion.PrintWithOptions(false, true, true))
}
