package installer

import (
	"log/slog"

	itbasisCoreExec "github.com/itbasis/tools/core/exec"
	itbasisCoreLog "github.com/itbasis/tools/core/log"
)

func (r *Installer) installGo() {
	for dependencyName, dependency := range r.dependencies.GoDependencies {
		slog.Info("install dependency: " + dependencyName + " [" + dependency.Version + "]")

		if err := r.goInstaller.Execute(
			itbasisCoreExec.WithRerun(),
			itbasisCoreExec.WithRestoreArgsIncludePrevious(
				itbasisCoreExec.IncludePrevArgsBefore,
				dependency.String(),
			),
		); err != nil {
			slog.Error("fail install dependency: "+dependencyName, itbasisCoreLog.SlogAttrError(err))
		}
	}
}
