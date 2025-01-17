package exec

import (
	"os/exec"

	itbasisCoreExec "github.com/itbasis/tools/core/exec"
	itbasisCoreOption "github.com/itbasis/tools/core/option"
	"github.com/spf13/cobra"
)

func NewGoExecutable(opts ...itbasisCoreOption.Option[exec.Cmd]) (*itbasisCoreExec.Executable, error) {
	return itbasisCoreExec.NewExecutable("go", opts...) //nolint:wrapcheck // TODO
}

func NewGoInstallWithCobra(cmd *cobra.Command) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("install"),
		itbasisCoreExec.WithCobraOut(cmd),
	)
}

func NewGoGetWithCobra(cmd *cobra.Command) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("get"),
		itbasisCoreExec.WithCobraOut(cmd),
	)
}

func NewGoModWithCobra(cmd *cobra.Command) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("mod"),
		itbasisCoreExec.WithCobraOut(cmd),
	)
}

func NewGoToolWithCobra(cmd *cobra.Command) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("tool"),
		itbasisCoreExec.WithCobraOut(cmd),
	)
}

func NewGoGenerateWithCobra(cmd *cobra.Command) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("generate"),
		itbasisCoreExec.WithCobraOut(cmd),
	)
}

func NewGoBuildWithCobra(cmd *cobra.Command) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("build"),
		itbasisCoreExec.WithCobraOut(cmd),
	)
}

func NewGoRunWithCobra(cmd *cobra.Command) (*itbasisCoreExec.Executable, error) {
	return NewGoExecutable(
		itbasisCoreExec.WithArgs("run"),
		itbasisCoreExec.WithCobraOut(cmd),
	)
}
