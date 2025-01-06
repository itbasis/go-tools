package test

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	builderCmd "github.com/itbasis/tools/builder/internal/cmd"
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareExec "github.com/itbasis/tools/middleware/exec"
	itbasisMiddlewareLog "github.com/itbasis/tools/middleware/log"
	ginkgoCommand "github.com/onsi/ginkgo/v2/ginkgo/command"
	ginkgoRun "github.com/onsi/ginkgo/v2/ginkgo/run"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	reportDir = path.Join("build", "reports")

	coverUnit          = "coverage-unit"
	coverUnitOut       = coverUnit + ".out"
	coverUnitHTML      = coverUnit + ".html"
	ginkgoCoverUnitOut = "ginkgo-" + coverUnitOut

	junitReportOut = "junit-report.xml"
)

func NewUnitTestCommand() *cobra.Command {
	return &cobra.Command{
		Use:    itbasisMiddlewareCmd.BuildUse("unit-test", builderCmd.UseArgPackages),
		Args:   cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
		PreRun: itbasisMiddlewareCmd.LogCommand,
		Run:    _run,
	}
}

func _run(cmd *cobra.Command, args []string) {
	itbasisMiddlewareCmd.RequireNoError(cmd, os.MkdirAll(reportDir, 0755))

	(&ginkgoCommand.Program{
		OutWriter:      cmd.OutOrStdout(),
		ErrWriter:      cmd.ErrOrStderr(),
		DefaultCommand: ginkgoRun.BuildRunCommand(),
		Exiter: func(code int) {
			slog.Debug(fmt.Sprintf("Ginkgo exit code: %d", code))

			if code != 0 {
				os.Exit(code)
			}
		},
	}).RunAndExit(
		[]string{
			"-race",
			"--cover", `--coverprofile=` + ginkgoCoverUnitOut,
			`--junit-report=` + junitReportOut,
			builderCmd.ArgPackages(builderCmd.DefaultPackages, args),
		},
	)

	itbasisMiddlewareCmd.RequireNoError(cmd, moveJunitReport(junitReportOut, path.Join(reportDir, junitReportOut)))
	itbasisMiddlewareCmd.RequireNoError(cmd, moveAndFilterCoverage(ginkgoCoverUnitOut, path.Join(reportDir, ginkgoCoverUnitOut)))

	var goToolCoverExec, err = itbasisBuilderExec.NewGoToolWithCobra(cmd)

	itbasisMiddlewareCmd.RequireNoError(cmd, err)
	itbasisMiddlewareCmd.RequireNoError(
		cmd,
		goToolCoverExec.Execute(
			itbasisMiddlewareExec.WithRerun(),
			itbasisMiddlewareExec.WithRestoreArgsIncludePrevious(
				itbasisMiddlewareExec.IncludePrevArgsBefore,
				"cover",
				"-func", ginkgoCoverUnitOut,
				"-o", path.Join(reportDir, coverUnitOut),
			),
		),
	)
	itbasisMiddlewareCmd.RequireNoError(
		cmd,
		goToolCoverExec.Execute(
			itbasisMiddlewareExec.WithRerun(),
			itbasisMiddlewareExec.WithRestoreArgsIncludePrevious(
				itbasisMiddlewareExec.IncludePrevArgsBefore,
				"cover",
				"-html", ginkgoCoverUnitOut,
				"-o", path.Join(reportDir, coverUnitHTML),
			),
		),
	)
}

func moveJunitReport(source, target string) error {
	slog.Debug("moving Junit report", slog.String("source", source), slog.String("target", target))

	if err := os.Rename(source, target); err != nil {
		return errors.Wrap(err, ErrMoveFile.Error())
	}

	return nil
}

func moveAndFilterCoverage(source, target string) error {
	slog.Debug("filtering and moving coverage", slog.String("source", source), slog.String("target", target))

	var sourceFile, errOpenFile = os.Open(source)
	if errOpenFile != nil {
		return errors.Wrap(errOpenFile, ErrMoveFile.Error())
	}

	defer func() {
		if err := sourceFile.Close(); err != nil {
			itbasisMiddlewareLog.Panic(fmt.Sprintf("fail close file: %s", source), itbasisMiddlewareLog.SlogAttrError(err))
		}
	}()

	var targetFile, errCreateFile = os.Create(target)
	if errCreateFile != nil {
		return errors.Wrap(errCreateFile, ErrMoveFile.Error())
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
			return errors.Wrap(errWrite, ErrMoveFile.Error())
		}
	}

	return errors.Wrap(scanner.Err(), ErrMoveFile.Error())
}
