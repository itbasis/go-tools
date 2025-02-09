package cmd

import (
	"io"
	"os"

	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
	"github.com/spf13/cobra"
)

const _optionOutKey itbasisCoreOption.Key = "option-out"

func WithDefaultOut() Option {
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
