package cmd

import (
	"io"
	"os"

	itbasisCoreOption "github.com/itbasis/tools/core/option"
	"github.com/spf13/cobra"
)

const _optionOutKey = "option-out"

func WithDefaultOut() itbasisCoreOption.Option[cobra.Command] {
	return &_optionOut{out: os.Stdout}
}

type _optionOut struct {
	out io.Writer
}

func (r *_optionOut) Key() itbasisCoreOption.Key { return _optionOutKey }
func (r *_optionOut) Apply(cmd *cobra.Command) error {
	cmd.SetOut(r.out)

	return nil
}
