package golang

import (
	"fmt"
	"log/slog"

	itbasisBuilderExec "github.com/itbasis/go-tools/builder/internal/exec"
	"github.com/itbasis/go-tools/builder/internal/installer/model"
	itbasisCoreExec "github.com/itbasis/go-tools/core/exec"
	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
)

const ProviderGoKey model.ProviderKey = "go"

type GoInstaller interface {
	Install(list model.GoDependencyList) error
}

var GoInstallEmpty = (GoInstaller)(nil)

type installer struct {
	exec *itbasisCoreExec.Executable
}

func NewGoInstaller(out itbasisCoreExec.CobraOut) (GoInstaller, error) {
	var exec, err = itbasisBuilderExec.NewGoInstallWithCobra(out)
	if err != nil {
		return GoInstallEmpty, err //nolint:wrapcheck // TODO
	}

	return &installer{
		exec: exec,
	}, nil
}

func (r *installer) Install(list model.GoDependencyList) error {
	slog.Info("installing Go dependencies")

	for name, dependency := range list {
		slog.Info(fmt.Sprintf("install dependency: %s[%s]", name, dependency.Version))

		if err := r.exec.Execute(
			itbasisCoreExec.WithRerun(),
			itbasisCoreExec.WithRestoreArgsIncludePrevious(
				itbasisCoreExec.IncludePrevArgsBefore,
				dependency.String(),
			),
		); err != nil {
			slog.Error("fail install dependency: "+name, itbasisCoreLog.SlogAttrError(err))
		}
	}

	return nil
}
