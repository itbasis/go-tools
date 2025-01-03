package generate

import (
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	"github.com/itbasis/tools/middleware/exec"
	"github.com/spf13/cobra"
)

var CmdGenerate = &cobra.Command{
	Use:  "generate",
	Args: cobra.NoArgs,
	Run: itbasisMiddlewareCmd.WrapActionLogging(
		func(cmd *cobra.Command, _ []string) {
			// itbasisMiddlewareCmd.ExecuteRequireNoError(dependencies.CmdDependencies)

			execGoGenerate, err := itbasisBuilderExec.NewGoGenerateWithCobra(cmd)
			itbasisMiddlewareCmd.RequireNoError(cmd, err)
			itbasisMiddlewareCmd.RequireNoError(
				cmd, execGoGenerate.Execute(
					exec.WithRestoreArgsIncludePrevious(exec.IncludePrevArgsBefore, itbasisBuilderExec.DefaultPackages),
				),
			)

			// itbasisMiddlewareCmd.ExecuteRequireNoError(update.CmdUpdate)
		},
	),
}
