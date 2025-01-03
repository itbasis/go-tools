package cmd

import (
	"io"
	"os"

	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
	"github.com/spf13/cobra"
)

const _optionOutKey = "option-out"

func WithDefaultOut() itbasisMiddlewareOption.Option[cobra.Command] {
	return &_optionOut{out: os.Stdout}
}

type _optionOut struct {
	out io.Writer
}

func (r *_optionOut) Key() itbasisMiddlewareOption.Key { return _optionOutKey }
func (r *_optionOut) Apply(cmd *cobra.Command) error {
	cmd.SetOut(r.out)

	return nil
}
