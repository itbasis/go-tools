package update

import (
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	"github.com/itbasis/tools/middleware/exec"
	"github.com/spf13/cobra"
)

var CmdUpdate = &cobra.Command{
	Use:   "update",
	Short: "update dependencies",
	Args:  cobra.NoArgs,
	Run: itbasisMiddlewareCmd.WrapActionLogging(
		func(cmd *cobra.Command, _ []string) {
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
					exec.WithRestoreArgsIncludePrevious(exec.IncludePrevArgsBefore, "-t", "-v", "-u", itbasisBuilderExec.DefaultPackages),
				),
			)
		},
	),
}
