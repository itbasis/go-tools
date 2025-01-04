package dependencies

import (
	"fmt"
	"log/slog"

	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	"github.com/itbasis/tools/middleware/exec"
	"github.com/spf13/cobra"
)

var CmdDependencies = &cobra.Command{
	Use:   "dependencies",
	Short: "Install dependencies",
	Args:  cobra.NoArgs,
	Run: itbasisMiddlewareCmd.WrapActionLogging(
		func(cmd *cobra.Command, _ []string) {
			var execGoInstall, err = itbasisBuilderExec.NewGoInstallWithCobra(cmd)
			itbasisMiddlewareCmd.RequireNoError(cmd, err)

			for _, dependency := range _dependencies.GoInstall {
				slog.Info(fmt.Sprintf("Installing dependency: %s", dependency))

				itbasisMiddlewareCmd.RequireNoError(
					cmd, execGoInstall.Execute(
						exec.WithRerun(),
						exec.WithRestoreArgsIncludePrevious(exec.IncludePrevArgsBefore, dependency.ToString()),
					),
				)
			}
		},
	),
}
