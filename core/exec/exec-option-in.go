package exec

import (
	"io"
	"os"
	"os/exec"

	itbasisCoreOption "github.com/itbasis/tools/core/option"
	"github.com/spf13/cobra"
)

const _optionInKey = "option-in"

func WithCobraIn(cmd *cobra.Command) itbasisCoreOption.Option[exec.Cmd] {
	return WithCustomIn(cmd.InOrStdin())
}
func WithStdIn() itbasisCoreOption.Option[exec.Cmd] {
	return WithCustomIn(os.Stdin)
}
func WithCustomIn(in io.Reader) itbasisCoreOption.Option[exec.Cmd] {
	return &_optionIn{in: in}
}

type _optionIn struct {
	in io.Reader
}

func (r *_optionIn) Key() itbasisCoreOption.Key { return _optionInKey }
func (r *_optionIn) Apply(cmd *exec.Cmd) error {
	cmd.Stdin = r.in

	return nil
}
