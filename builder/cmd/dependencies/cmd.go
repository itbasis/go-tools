package dependencies

import (
	"fmt"
	"log/slog"

	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisCoreCmd "github.com/itbasis/tools/core/cmd"
	itbasisCoreExec "github.com/itbasis/tools/core/exec"
	"github.com/spf13/cobra"
)

func NewDependenciesCommand() *cobra.Command {
	return &cobra.Command{
		Use:    itbasisCoreCmd.BuildUse("dependencies"),
		Short:  "Install dependencies",
		Args:   cobra.NoArgs,
		PreRun: itbasisCoreCmd.LogCommand,
		Run:    _run,
	}
}

func _run(cmd *cobra.Command, _ []string) {
	var execGoInstall, err = itbasisBuilderExec.NewGoInstallWithCobra(cmd)

	itbasisCoreCmd.RequireNoError(cmd, err)

	for _, dependency := range _dependencies.GoInstall {
		slog.Info(fmt.Sprintf("Installing dependency: %s", dependency))

		itbasisCoreCmd.RequireNoError(
			cmd, execGoInstall.Execute(
				itbasisCoreExec.WithRerun(),
				itbasisCoreExec.WithRestoreArgsIncludePrevious(itbasisCoreExec.IncludePrevArgsBefore, dependency.ToString()),
			),
		)
	}
}
