package build

import (
	"log/slog"
	"os"
	"path/filepath"
	"reflect"

	builderCmd "github.com/itbasis/tools/builder/internal/cmd"
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisCoreCmd "github.com/itbasis/tools/core/cmd"
	itbasisCoreExec "github.com/itbasis/tools/core/exec"
	itbasisCoreLog "github.com/itbasis/tools/core/log"
	itbasisCoreVersion "github.com/itbasis/tools/core/version"
	"github.com/spf13/cobra"
)

var (
	_flagOs      string
	_flagArch    string
	_flagOutput  string
	_flagVersion = itbasisCoreVersion.Unversioned
)

func NewBuildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    itbasisCoreCmd.BuildUse("build", builderCmd.UseArgPath),
		Short:  "Building an application for the current platform",
		Args:   cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
		PreRun: itbasisCoreCmd.LogCommand,
		Run:    _run,
	}

	cmd.Flags().StringVarP(&_flagOutput, "output", "", "", "")
	cmd.Flags().StringVarP(&_flagOs, "build-os", "", "", "")
	cmd.Flags().StringVarP(&_flagArch, "build-arch", "", "", "")
	cmd.Flags().StringVarP(&_flagVersion, "build-version", "", _flagVersion, "")

	return cmd
}

func _run(cmd *cobra.Command, args []string) {
	var (
		versionPkgPath = reflect.TypeFor[itbasisCoreVersion.Version]().PkgPath() + ".version"
		buildArgs      = []string{
			`-trimpath`,
			`-pgo`, `auto`,
			`-ldflags`, `-w -extldflags '-static' -X '` + versionPkgPath + `=` + _flagVersion + `'`,
			`-tags`, `musl`,
		}
	)

	if _flagOutput != "" {
		buildArgs = append(buildArgs, "-o", _flagOutput)

		itbasisCoreCmd.RequireNoError(cmd, os.MkdirAll(filepath.Dir(_flagOutput), os.ModePerm))
	}

	buildArgs = append(buildArgs, args[0])

	slog.Debug("build with arguments", itbasisCoreLog.SlogAttrSliceWithSeparator("buildArgs", " ", buildArgs))

	execGoBuild, errGoBuild := itbasisBuilderExec.NewGoBuildWithCobra(cmd)
	itbasisCoreCmd.RequireNoError(cmd, errGoBuild)
	itbasisCoreCmd.RequireNoError(
		cmd,
		execGoBuild.Execute(
			itbasisCoreExec.WithRestoreArgsIncludePrevious(itbasisCoreExec.IncludePrevArgsBefore, buildArgs...),
		),
	)
}
