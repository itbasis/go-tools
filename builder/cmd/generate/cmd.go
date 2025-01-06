package generate

import (
	builderCmd "github.com/itbasis/tools/builder/internal/cmd"
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	"github.com/itbasis/tools/middleware/exec"
	"github.com/spf13/cobra"
)

func NewGenerateCommand() *cobra.Command {
	return &cobra.Command{
		Use:    itbasisMiddlewareCmd.BuildUse("generate", builderCmd.UseArgPackages),
		Args:   cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
		PreRun: itbasisMiddlewareCmd.LogCommand,
		Run:    _run,
	}
}

func _run(cmd *cobra.Command, args []string) {
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
}
