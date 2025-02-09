package exec

import (
	"io"
	"os"
	"os/exec"

	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
)

const _optionOutKey itbasisCoreOption.Key = "option-out"

type CobraOut interface {
	OutOrStdout() io.Writer
	ErrOrStderr() io.Writer
}

func WithCobraOut(cmd CobraOut) Option {
	return WithCustomOut(cmd.OutOrStdout(), cmd.ErrOrStderr())
}

func WithStdOut() Option {
	return WithCustomOut(os.Stdout, os.Stderr)
}

func WithCustomOut(out, err io.Writer) Option {
	return &_optionOut{out: out, err: err}
}

type _optionOut struct {
	out io.Writer
	err io.Writer
}

func (r *_optionOut) Key() itbasisCoreOption.Key { return _optionOutKey }
func (r *_optionOut) Apply(cmd *exec.Cmd) error {
	cmd.Stdout = r.out
	cmd.Stderr = r.err

	return nil
}
