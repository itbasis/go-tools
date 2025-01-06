package dependencies

import (
	"fmt"
	"log/slog"

	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareExec "github.com/itbasis/tools/middleware/exec"
	"github.com/spf13/cobra"
)

func NewDependenciesCommand() *cobra.Command {
	return &cobra.Command{
		Use:    itbasisMiddlewareCmd.BuildUse("dependencies"),
		Short:  "Install dependencies",
		Args:   cobra.NoArgs,
		PreRun: itbasisMiddlewareCmd.LogCommand,
		Run:    _run,
	}
}

func _run(cmd *cobra.Command, _ []string) {
	var execGoInstall, err = itbasisBuilderExec.NewGoInstallWithCobra(cmd)

	itbasisMiddlewareCmd.RequireNoError(cmd, err)

	for _, dependency := range _dependencies.GoInstall {
		slog.Info(fmt.Sprintf("Installing dependency: %s", dependency))

		itbasisMiddlewareCmd.RequireNoError(
			cmd, execGoInstall.Execute(
				itbasisMiddlewareExec.WithRerun(),
				itbasisMiddlewareExec.WithRestoreArgsIncludePrevious(itbasisMiddlewareExec.IncludePrevArgsBefore, dependency.ToString()),
			),
		)
	}
}
