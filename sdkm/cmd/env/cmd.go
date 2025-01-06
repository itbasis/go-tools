package env

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

func NewEnvCommand() *cobra.Command {
	return &cobra.Command{
		Use:        itbasisMiddlewareCmd.BuildUse("env", sdkmCmd.UseArgRequirePlugins, sdkmCmd.UseArgVersion),
		Short:      "Displays environment variables inside the environment used for the plugin",
		ArgAliases: []string{sdkmCmd.ArgAliasPlugin, sdkmCmd.ArgAliasVersion},
		Args:       cobra.MatchAll(cobra.RangeArgs(1, 2), cobra.OnlyValidArgs),
		Run:        _run,
	}
}

func _run(cmd *cobra.Command, args []string) {
	var (
		sdkmPlugin = sdkmCmd.GetPluginByName(cmd, args[_idxArgPlugin])
		envMap     map[string]string
		err        error
	)

	if len(args) == 1 {
		envMap, err = sdkmPlugin.Env(cmd.Context(), itbasisMiddlewareOs.Pwd())
	} else {
		envMap, err = sdkmPlugin.EnvByVersion(cmd.Context(), args[_idxArgVersion])
	}

	if err != nil {
		itbasisMiddlewareCmd.Fatal(cmd, err)
	}

	for _, env := range itbasisMiddlewareOs.EnvMapToSlices(envMap) {
		cmd.Println(env)
	}
}
