package lint

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareExec "github.com/itbasis/tools/middleware/exec"
	"github.com/spf13/cobra"
)

var CmdLint = &cobra.Command{
	Use:  "lint",
	Args: cobra.NoArgs,
	Run: itbasisMiddlewareCmd.WrapActionLogging(
		func(cmd *cobra.Command, _ []string) {
			execEditorConfigChecker, err := itbasisMiddlewareExec.NewExecutable(
				"editorconfig-checker",
				itbasisMiddlewareExec.WithCobraOut(cmd),
			)
			itbasisMiddlewareCmd.RequireNoError(cmd, err)
			itbasisMiddlewareCmd.RequireNoError(cmd, execEditorConfigChecker.Execute())

			execGolangCiLint, err := itbasisMiddlewareExec.NewExecutable(
				"golangci-lint",
				itbasisMiddlewareExec.WithArgs("run"),
				itbasisMiddlewareExec.WithCobraOut(cmd),
			)
			itbasisMiddlewareCmd.RequireNoError(cmd, err)
			itbasisMiddlewareCmd.RequireNoError(cmd, execGolangCiLint.Execute())
		},
	),
}
