package exec

import (
	"io"
	"os"
	"os/exec"

	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
	"github.com/spf13/cobra"
)

const _optionOutKey = "option-out"

func WithCobraOut(cmd *cobra.Command) itbasisMiddlewareOption.Option[exec.Cmd] {
	return WithCustomOut(cmd.OutOrStdout(), cmd.ErrOrStderr())
}

func WithStdOut() itbasisMiddlewareOption.Option[exec.Cmd] {
	return WithCustomOut(os.Stdout, os.Stderr)
}

func WithCustomOut(out, err io.Writer) itbasisMiddlewareOption.Option[exec.Cmd] {
	return &_optionOut{out: out, err: err}
}

type _optionOut struct {
	out io.Writer
	err io.Writer
}

func (r *_optionOut) Key() itbasisMiddlewareOption.Key { return _optionOutKey }
func (r *_optionOut) Apply(cmd *exec.Cmd) error {
	cmd.Stdout = r.out
	cmd.Stderr = r.err

	return nil
}
