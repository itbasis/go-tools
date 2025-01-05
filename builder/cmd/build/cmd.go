package build

import (
	"log/slog"
	"os"
	"path/filepath"
	"reflect"

	builderCmd "github.com/itbasis/tools/builder/internal/cmd"
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareExec "github.com/itbasis/tools/middleware/exec"
	itbasisMiddlewareLog "github.com/itbasis/tools/middleware/log"
	itbasisMiddlewareVersion "github.com/itbasis/tools/middleware/version"
	"github.com/spf13/cobra"
)

var (
	_flagOs      string
	_flagArch    string
	_flagOutput  string
	_flagVersion = itbasisMiddlewareVersion.Unversioned
)

var _cmdBuild = &cobra.Command{
	Use:   itbasisMiddlewareCmd.BuildUse("build", itbasisMiddlewareCmd.UseFlags, builderCmd.UseArgPath),
	Short: "Building an application for the current platform",
	Args:  cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
	Run:   _run,
}

func NewBuildCommand() *cobra.Command {
	_cmdBuild.Flags().StringVarP(&_flagOutput, "output", "", "", "")
	_cmdBuild.Flags().StringVarP(&_flagOs, "build-os", "", "", "")
	_cmdBuild.Flags().StringVarP(&_flagArch, "build-arch", "", "", "")
	_cmdBuild.Flags().StringVarP(&_flagVersion, "build-version", "", _flagVersion, "")

	return _cmdBuild
}

func _run(cmd *cobra.Command, args []string) {
	var (
		versionPkgPath = reflect.TypeFor[itbasisMiddlewareVersion.Version]().PkgPath() + ".version"
		buildArgs      = []string{
			`-trimpath`,
			`-pgo`, `auto`,
			`-ldflags`, `-w -extldflags '-static' -X '` + versionPkgPath + `=` + _flagVersion + `'`,
			`-tags`, `musl`,
		}
	)

	if _flagOutput != "" {
		buildArgs = append(buildArgs, "-o", _flagOutput)

		itbasisMiddlewareCmd.RequireNoError(cmd, os.MkdirAll(filepath.Dir(_flagOutput), os.ModePerm))
	}

	buildArgs = append(buildArgs, args[0])

	slog.Debug("build with arguments", itbasisMiddlewareLog.SlogAttrStringsWithSeparator("buildArgs", " ", buildArgs))

	execGoBuild, errGoBuild := itbasisBuilderExec.NewGoBuildWithCobra(cmd)
	itbasisMiddlewareCmd.RequireNoError(cmd, errGoBuild)
	itbasisMiddlewareCmd.RequireNoError(
		cmd,
		execGoBuild.Execute(
			itbasisMiddlewareExec.WithRestoreArgsIncludePrevious(itbasisMiddlewareExec.IncludePrevArgsBefore, buildArgs...),
		),
	)
}
