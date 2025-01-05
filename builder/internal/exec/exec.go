package exec

import (
	"os/exec"

	itbasisMiddlewareExec "github.com/itbasis/tools/middleware/exec"
	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
	"github.com/spf13/cobra"
)

func NewGoExecutable(opts ...itbasisMiddlewareOption.Option[exec.Cmd]) (*itbasisMiddlewareExec.Executable, error) {
	return itbasisMiddlewareExec.NewExecutable("go", opts...)
}

func NewGoInstallWithCobra(cmd *cobra.Command) (*itbasisMiddlewareExec.Executable, error) {
	return NewGoExecutable(
		itbasisMiddlewareExec.WithArgs("install"),
		itbasisMiddlewareExec.WithCobraOut(cmd),
	)
}

func NewGoGetWithCobra(cmd *cobra.Command) (*itbasisMiddlewareExec.Executable, error) {
	return NewGoExecutable(
		itbasisMiddlewareExec.WithArgs("get"),
		itbasisMiddlewareExec.WithCobraOut(cmd),
	)
}

func NewGoModWithCobra(cmd *cobra.Command) (*itbasisMiddlewareExec.Executable, error) {
	return NewGoExecutable(
		itbasisMiddlewareExec.WithArgs("mod"),
		itbasisMiddlewareExec.WithCobraOut(cmd),
	)
}

func NewGoToolWithCobra(cmd *cobra.Command) (*itbasisMiddlewareExec.Executable, error) {
	return NewGoExecutable(
		itbasisMiddlewareExec.WithArgs("tool"),
		itbasisMiddlewareExec.WithCobraOut(cmd),
	)
}

func NewGoGenerateWithCobra(cmd *cobra.Command) (*itbasisMiddlewareExec.Executable, error) {
	return NewGoExecutable(
		itbasisMiddlewareExec.WithArgs("generate"),
		itbasisMiddlewareExec.WithCobraOut(cmd),
	)
}

func NewGoBuildWithCobra(cmd *cobra.Command) (*itbasisMiddlewareExec.Executable, error) {
	return NewGoExecutable(
		itbasisMiddlewareExec.WithArgs("build"),
		itbasisMiddlewareExec.WithCobraOut(cmd),
	)
}

func NewGoRunWithCobra(cmd *cobra.Command) (*itbasisMiddlewareExec.Executable, error) {
	return NewGoExecutable(
		itbasisMiddlewareExec.WithArgs("run"),
		itbasisMiddlewareExec.WithCobraOut(cmd),
	)
}
