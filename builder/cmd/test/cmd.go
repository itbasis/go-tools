package test

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	"github.com/itbasis/tools/middleware/exec"
	itbasisMiddlewareLog "github.com/itbasis/tools/middleware/log"
	ginkgoCommand "github.com/onsi/ginkgo/v2/ginkgo/command"
	ginkgoRun "github.com/onsi/ginkgo/v2/ginkgo/run"
	"github.com/spf13/cobra"
)

var (
	reportDir = path.Join("build", "reports")

	coverUnit          = "coverage-unit"
	coverUnitOut       = coverUnit + ".out"
	coverUnitHTML      = coverUnit + ".html"
	ginkgoCoverUnitOut = "ginkgo-" + coverUnitOut

	reportGinkgoCoverUnitOut = path.Join(reportDir, ginkgoCoverUnitOut)

	junitReportOut = "junit-report.xml"
)

var CmdUnitTest = &cobra.Command{
	Use:  "unit-test",
	Args: cobra.NoArgs,
	Run: itbasisMiddlewareCmd.WrapActionLogging(
		func(cmd *cobra.Command, _ []string) {
			itbasisMiddlewareCmd.RequireNoError(cmd, os.MkdirAll(reportDir, 0755))

			(&ginkgoCommand.Program{
				OutWriter:      cmd.OutOrStdout(),
				ErrWriter:      cmd.ErrOrStderr(),
				DefaultCommand: ginkgoRun.BuildRunCommand(),
			}).RunAndExit(
				[]string{
					"-race",
					"--cover", `--coverprofile=` + ginkgoCoverUnitOut,
					`--junit-report=` + junitReportOut,
					itbasisBuilderExec.DefaultPackages,
				},
			)

			itbasisMiddlewareCmd.RequireNoError(cmd, moveAndFilterCoverage(ginkgoCoverUnitOut, reportGinkgoCoverUnitOut))

			var goToolCoverExec, err = itbasisBuilderExec.NewGoToolWithCobra(cmd)
			itbasisMiddlewareCmd.RequireNoError(cmd, err)

			itbasisMiddlewareCmd.RequireNoError(
				cmd,
				goToolCoverExec.Execute(
					exec.WithRestoreArgsIncludePrevious(
						exec.IncludePrevArgsBefore,
						"cover",
						// "-func", reportGinkgoCoverUnitOut,
						"-func", ginkgoCoverUnitOut,
						"-o", path.Join(reportDir, coverUnitOut),
					),
				),
			)
			itbasisMiddlewareCmd.RequireNoError(
				cmd,
				goToolCoverExec.Execute(
					exec.WithRestoreArgsIncludePrevious(
						exec.IncludePrevArgsBefore,
						"cover",
						"-html", ginkgoCoverUnitOut,
						// "-html", reportGinkgoCoverUnitOut,
						"-o", path.Join(reportDir, coverUnitHTML),
					),
				),
			)
		},
	),
}

func moveAndFilterCoverage(source, target string) error {
	slog.Debug("filtering anf moving coverage", slog.String("source", source), slog.String("target", target))

	var sourceFile, errOpenFile = os.Open(source)
	if errOpenFile != nil {
		return errOpenFile
	}

	defer func() {
		if err := sourceFile.Close(); err != nil {
			itbasisMiddlewareLog.Panic(fmt.Sprintf("fail close file: %s", source), itbasisMiddlewareLog.SlogAttrError(err))
		}
	}()

	var targetFile, errCreateFile = os.Create(target)
	if errCreateFile != nil {
		return errCreateFile
	}

	defer func() {
		if err := targetFile.Close(); err != nil {
			itbasisMiddlewareLog.Panic(fmt.Sprintf("fail close file: %s", target), itbasisMiddlewareLog.SlogAttrError(err))
		}
	}()

	var scanner = bufio.NewScanner(sourceFile)

	for scanner.Scan() {
		var line = scanner.Text()

		if strings.Contains(line, ".mock.go") {
			continue
		}

		if _, errWrite := targetFile.WriteString(line + "\n"); errWrite != nil {
			return errWrite
		}
	}

	return scanner.Err()
}
