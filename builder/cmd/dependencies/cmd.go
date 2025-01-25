package dependencies

import (
	_ "embed"
	"log/slog"

	builderInstaller "github.com/itbasis/tools/builder/internal/installer"
	itbasisCoreCmd "github.com/itbasis/tools/core/cmd"
	"github.com/spf13/cobra"
)

//go:embed dependencies.json
var _defaultDependencies []byte

var (
	_flagDependenciesFile string
	_flagShow             bool
)

func NewDependenciesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    itbasisCoreCmd.BuildUse("dependencies"),
		Short:  "Install dependencies",
		Args:   cobra.NoArgs,
		PreRun: itbasisCoreCmd.LogCommand,
		Run:    _run,
	}

	flags := cmd.Flags()

	flags.StringVarP(
		&_flagDependenciesFile,
		"dependencies-file",
		"f",
		"",
		"dependencies file path. If not specified, the embedded list will be used",
	)
	flags.BoolVar(&_flagShow, "show-default", false, "show default dependencies for install")

	return cmd
}

func _run(cmd *cobra.Command, _ []string) {
	if _flagShow {
		_, err := cmd.OutOrStdout().Write(_defaultDependencies)
		itbasisCoreCmd.RequireNoError(cmd, err)

		return
	}

	var optionDependencies builderInstaller.Option

	if _flagDependenciesFile != "" {
		slog.Info("using dependencies file: " + _flagDependenciesFile)

		optionDependencies = builderInstaller.WithFile(_flagDependenciesFile)
	} else {
		optionDependencies = builderInstaller.WithJsonData(_defaultDependencies)
	}

	installer, errInstaller := builderInstaller.NewInstaller(cmd, optionDependencies)
	itbasisCoreCmd.RequireNoError(cmd, errInstaller)

	installer.Install()
}
