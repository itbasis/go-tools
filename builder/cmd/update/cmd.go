package update

import (
	builderCmd "github.com/itbasis/tools/builder/internal/cmd"
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisCoreCmd "github.com/itbasis/tools/core/cmd"
	itbasisCoreExec "github.com/itbasis/tools/core/exec"
	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	return &cobra.Command{
		Use:    itbasisCoreCmd.BuildUse("update", builderCmd.UseArgPackages),
		Short:  "update dependencies",
		Args:   cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
		PreRun: itbasisCoreCmd.LogCommand,
		Run:    _run,
	}
}

func _run(cmd *cobra.Command, args []string) {
	execGoMod, errGoMod := itbasisBuilderExec.NewGoModWithCobra(cmd)
	itbasisCoreCmd.RequireNoError(cmd, errGoMod)
	itbasisCoreCmd.RequireNoError(
		cmd, execGoMod.Execute(
			itbasisCoreExec.WithRestoreArgsIncludePrevious(itbasisCoreExec.IncludePrevArgsBefore, "tidy"),
		),
	)

	execGoGet, errGoGet := itbasisBuilderExec.NewGoGetWithCobra(cmd)
	itbasisCoreCmd.RequireNoError(cmd, errGoGet)
	itbasisCoreCmd.RequireNoError(
		cmd, execGoGet.Execute(
			itbasisCoreExec.WithRestoreArgsIncludePrevious(
				itbasisCoreExec.IncludePrevArgsBefore,
				"-t",
				"-v",
				"-u",
				builderCmd.ArgPackages(builderCmd.DefaultPackages, args),
			),
		),
	)
}
