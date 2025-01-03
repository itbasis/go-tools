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

			for _, dependency := range []string{
				// curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) latest
				"github.com/golangci/golangci-lint/cmd/golangci-lint@latest",
				// "github.com/onsi/ginkgo/v2/ginkgo@latest",
				"go.uber.org/mock/mockgen@latest",
				"github.com/editorconfig-checker/editorconfig-checker/v3/cmd/editorconfig-checker@latest",
			} {
				slog.Info(fmt.Sprintf("Installing dependency: %s", dependency))

				itbasisMiddlewareCmd.RequireNoError(
					cmd, execGoInstall.Execute(
						exec.WithRestoreArgsIncludePrevious(exec.IncludePrevArgsBefore, dependency),
					),
				)
			}
		},
	),
}
