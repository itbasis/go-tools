package installer

import (
	"log/slog"

	itbasisBuilderExec "github.com/itbasis/go-tools/builder/internal/exec"
	itbasisCoreExec "github.com/itbasis/go-tools/core/exec"
	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
)

func (r *Installer) installGo() error {
	goInstaller, errGoInstaller := itbasisBuilderExec.NewGoInstallWithCobra(r.cobraOut)
	if errGoInstaller != nil {
		return errGoInstaller
	}

	for dependencyName, dependency := range r.dependencies.GoDependencies {
		slog.Info("install Go dependency: " + dependencyName + " [" + dependency.Version + "]")

		if err := goInstaller.Execute(
			itbasisCoreExec.WithRerun(),
			itbasisCoreExec.WithRestoreArgsIncludePrevious(
				itbasisCoreExec.IncludePrevArgsBefore,
				dependency.String(),
			),
		); err != nil {
			slog.Error("fail install dependency: "+dependencyName, itbasisCoreLog.SlogAttrError(err))
		}
	}

	return nil
}
