package generate

import (
	"fmt"

	builderCmd "github.com/itbasis/tools/builder/internal/cmd"
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	"github.com/itbasis/tools/middleware/exec"
	"github.com/spf13/cobra"
)

var CmdGenerate = &cobra.Command{
	Use:  itbasisMiddlewareCmd.BuildUse("generate", itbasisMiddlewareCmd.UseFlags, builderCmd.UseArgPackages),
	Args: cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
	Run: itbasisMiddlewareCmd.WrapActionLogging(
		func(cmd *cobra.Command, args []string) {
			// itbasisMiddlewareCmd.ExecuteRequireNoError(dependencies.CmdDependencies)

			execGoGenerate, err := itbasisBuilderExec.NewGoGenerateWithCobra(cmd)
			itbasisMiddlewareCmd.RequireNoError(cmd, err)
			itbasisMiddlewareCmd.RequireNoError(
				cmd, execGoGenerate.Execute(
					exec.WithRestoreArgsIncludePrevious(
						exec.IncludePrevArgsBefore,
						builderCmd.ArgPackages(builderCmd.DefaultPackages, args),
					),
				),
			)

			// itbasisMiddlewareCmd.ExecuteRequireNoError(update.CmdUpdate)
		},
	),
}
