package installer

import (
	itbasisCoreExec "github.com/itbasis/go-tools/core/exec"
	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
)

type Installer struct {
	dependencies Dependencies

	cobraOut itbasisCoreExec.CobraOut
}

func NewInstaller(cobraOut itbasisCoreExec.CobraOut, opts ...Option) (*Installer, error) {
	installer := &Installer{
		cobraOut: cobraOut,
	}

	if err := itbasisCoreOption.ApplyOptions(installer, opts, nil); err != nil {
		return nil, err
	}

	return installer, nil
}

func (r *Installer) Install() error {
	if err := r.installGo(); err != nil {
		return err
	}
	// if err := r.installGitHub(); err != nil {
	// 	return err
	// }

	return nil
}
