package update

import (
	builderCmd "github.com/itbasis/tools/builder/internal/cmd"
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareExec "github.com/itbasis/tools/middleware/exec"
	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	return &cobra.Command{
		Use:    itbasisMiddlewareCmd.BuildUse("update", builderCmd.UseArgPackages),
		Short:  "update dependencies",
		Args:   cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
		PreRun: itbasisMiddlewareCmd.LogCommand,
		Run:    _run,
	}
}

func _run(cmd *cobra.Command, args []string) {
	execGoMod, errGoMod := itbasisBuilderExec.NewGoModWithCobra(cmd)
	itbasisMiddlewareCmd.RequireNoError(cmd, errGoMod)
	itbasisMiddlewareCmd.RequireNoError(
		cmd, execGoMod.Execute(
			itbasisMiddlewareExec.WithRestoreArgsIncludePrevious(itbasisMiddlewareExec.IncludePrevArgsBefore, "tidy"),
		),
	)

	execGoGet, errGoGet := itbasisBuilderExec.NewGoGetWithCobra(cmd)
	itbasisMiddlewareCmd.RequireNoError(cmd, errGoGet)
	itbasisMiddlewareCmd.RequireNoError(
		cmd, execGoGet.Execute(
			itbasisMiddlewareExec.WithRestoreArgsIncludePrevious(
				itbasisMiddlewareExec.IncludePrevArgsBefore,
				"-t",
				"-v",
				"-u",
				builderCmd.ArgPackages(builderCmd.DefaultPackages, args),
			),
		),
	)
}
