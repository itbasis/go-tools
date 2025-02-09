package exec

import (
	itbasisCoreExec "github.com/itbasis/go-tools/core/exec"
)

func NewGoExecutable(opts ...itbasisCoreExec.Option) (*itbasisCoreExec.Executable, error) {
	return itbasisCoreExec.NewExecutable("go", opts...) //nolint:wrapcheck // TODO
}

func NewGoInstallWithCobra(cobraOut itbasisCoreExec.CobraOut) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("install"),
		itbasisCoreExec.WithCobraOut(cobraOut),
	)
}

func NewGoGetWithCobra(cobraOut itbasisCoreExec.CobraOut) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("get"),
		itbasisCoreExec.WithCobraOut(cobraOut),
	)
}

func NewGoModWithCobra(cobraOut itbasisCoreExec.CobraOut) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("mod"),
		itbasisCoreExec.WithCobraOut(cobraOut),
	)
}

func NewGoToolWithCobra(cobraOut itbasisCoreExec.CobraOut) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("tool"),
		itbasisCoreExec.WithCobraOut(cobraOut),
	)
}

func NewGoGenerateWithCobra(cobraOut itbasisCoreExec.CobraOut) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("generate"),
		itbasisCoreExec.WithCobraOut(cobraOut),
	)
}

func NewGoBuildWithCobra(cobraOut itbasisCoreExec.CobraOut) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("build"),
		itbasisCoreExec.WithCobraOut(cobraOut),
	)
}

func NewGoRunWithCobra(cobraOut itbasisCoreExec.CobraOut) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("run"),
		itbasisCoreExec.WithCobraOut(cobraOut),
	)
}
