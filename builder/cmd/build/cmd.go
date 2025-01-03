package build

import (
	"log/slog"
	"os"
	"path/filepath"
	"reflect"

	"github.com/itbasis/tools/builder/internal/exec"
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
	_flagVersion string = "0.0.0"
)

var _cmdBuild = &cobra.Command{
	Use:        "build [flags] {<path>}",
	Short:      "Building an application for the current platform",
	ArgAliases: []string{"path"},
	Args:       cobra.MatchAll(cobra.OnlyValidArgs),
	Run:        _run,
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

	execGoBuild, errGoBuild := exec.NewGoBuildWithCobra(cmd)
	itbasisMiddlewareCmd.RequireNoError(cmd, errGoBuild)
	itbasisMiddlewareCmd.RequireNoError(
		cmd,
		execGoBuild.Execute(
			itbasisMiddlewareExec.WithRestoreArgsIncludePrevious(itbasisMiddlewareExec.IncludePrevArgsBefore, buildArgs...),
		),
	)
}
