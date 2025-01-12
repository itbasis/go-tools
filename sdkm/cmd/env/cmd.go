package env

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

func NewEnvCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "env",
		Short: "Displays environment variables inside the environment used for the plugin",
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
		sdkmPlugin = sdkmPlugins.GetPluginByID(cmd)
		envMap     map[string]string
		err        error
	)

	if len(args) == 0 {
		envMap, err = sdkmPlugin.Env(cmd.Context(), sdkmCmd.IsFlagRebuildCache(cmd), itbasisMiddlewareOs.Pwd())
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
