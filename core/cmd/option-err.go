package cmd

import (
	"io"
	"os"

	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
	"github.com/spf13/cobra"
)

const _optionErrKey = "option-err"

func WithDefaultErr() itbasisCoreOption.Option[cobra.Command] {
	return &_optionErr{out: os.Stdout}
}

type _optionErr struct {
	out io.Writer
}

func (r *_optionErr) Key() itbasisCoreOption.Key { return _optionErrKey }
func (r *_optionErr) Apply(cmd *cobra.Command) error {
	cmd.SetErr(r.out)

	return nil
}
