package cmd

import (
	"io"
	"os"

	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
	"github.com/spf13/cobra"
)

const _optionErrKey = "option-err"

func WithDefaultErr() itbasisMiddlewareOption.Option[cobra.Command] {
	return &_optionErr{out: os.Stdout}
}

type _optionErr struct {
	out io.Writer
}

func (r *_optionErr) Key() itbasisMiddlewareOption.Key { return _optionErrKey }
func (r *_optionErr) Apply(cmd *cobra.Command) error {
	cmd.SetErr(r.out)

	return nil
}
