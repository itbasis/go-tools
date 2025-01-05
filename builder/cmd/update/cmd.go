package update

import (
	builderCmd "github.com/itbasis/tools/builder/internal/cmd"
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	"github.com/itbasis/tools/middleware/exec"
	"github.com/spf13/cobra"
)

var CmdUpdate = &cobra.Command{
	Use:   itbasisMiddlewareCmd.BuildUse("update", itbasisMiddlewareCmd.UseFlags, builderCmd.UseArgPackages),
	Short: "update dependencies",
	Args:  cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
	Run: itbasisMiddlewareCmd.WrapActionLogging(
		func(cmd *cobra.Command, args []string) {
			// itbasisMiddlewareCmd.ExecuteRequireNoError(dependencies.CmdDependencies)

			execGoMod, errGoMod := itbasisBuilderExec.NewGoModWithCobra(cmd)
			itbasisMiddlewareCmd.RequireNoError(cmd, errGoMod)
			itbasisMiddlewareCmd.RequireNoError(
				cmd, execGoMod.Execute(
					exec.WithRestoreArgsIncludePrevious(exec.IncludePrevArgsBefore, "tidy"),
				),
			)

			execGoGet, errGoGet := itbasisBuilderExec.NewGoGetWithCobra(cmd)
			itbasisMiddlewareCmd.RequireNoError(cmd, errGoGet)
			itbasisMiddlewareCmd.RequireNoError(
				cmd, execGoGet.Execute(
					exec.WithRestoreArgsIncludePrevious(
						exec.IncludePrevArgsBefore,
						"-t",
						"-v",
						"-u",
						builderCmd.ArgPackages(builderCmd.DefaultPackages, args),
					),
				),
			)
		},
	),
}
