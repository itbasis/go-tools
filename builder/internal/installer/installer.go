package installer

import (
	itbasisBuilderExec "github.com/itbasis/tools/builder/internal/exec"
	itbasisCoreExec "github.com/itbasis/tools/core/exec"
	itbasisCoreOption "github.com/itbasis/tools/core/option"
)

type Installer struct {
	dependencies Dependencies

	goInstaller *itbasisCoreExec.Executable
}

func NewInstaller(cobraOut itbasisCoreExec.CobraOut, opts ...Option) (*Installer, error) {
	goInstaller, errGoInstaller := itbasisBuilderExec.NewGoInstallWithCobra(cobraOut)
	if errGoInstaller != nil {
		return nil, errGoInstaller
	}

	installer := &Installer{
		goInstaller: goInstaller,
	}

	if err := itbasisCoreOption.ApplyOptions(installer, opts, nil); err != nil {
		return nil, err
	}

	return installer, nil
}

func (r *Installer) Install() {
	r.installGo()
}
