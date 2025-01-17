package generate

import (
	builderCmd "github.com/itbasis/tools/builder/internal/cmd"
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisCoreCmd "github.com/itbasis/tools/core/cmd"
	itbasisCoreExec "github.com/itbasis/tools/core/exec"
	"github.com/spf13/cobra"
)

func NewGenerateCommand() *cobra.Command {
	return &cobra.Command{
		Use:    itbasisCoreCmd.BuildUse("generate", builderCmd.UseArgPackages),
		Args:   cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
		PreRun: itbasisCoreCmd.LogCommand,
		Run:    _run,
	}
}

func _run(cmd *cobra.Command, args []string) {
	execGoGenerate, err := itbasisBuilderExec.NewGoGenerateWithCobra(cmd)
	itbasisCoreCmd.RequireNoError(cmd, err)
	itbasisCoreCmd.RequireNoError(
		cmd, execGoGenerate.Execute(
			itbasisCoreExec.WithRestoreArgsIncludePrevious(
				itbasisCoreExec.IncludePrevArgsBefore,
				builderCmd.ArgPackages(builderCmd.DefaultPackages, args),
			),
		),
	)
}
