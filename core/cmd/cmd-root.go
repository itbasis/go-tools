package cmd

import (
	itbasisCoreOption "github.com/itbasis/go-tools/core/option"

	"github.com/spf13/cobra"
)

func InitDefaultCmdRoot(shortName string, opts ...Option) (*cobra.Command, error) {
	var cmd = &cobra.Command{Short: shortName}

	if err := itbasisCoreOption.ApplyOptions(
		cmd,
		opts, map[itbasisCoreOption.Key]itbasisCoreOption.LazyOptionFunc[cobra.Command]{
			_optionVersionKey:   WithDefaultVersion,
			_optionOutKey:       WithDefaultOut,
			_optionErrKey:       WithDefaultErr,
			_optionDebugFlagKey: WithDefaultFlagDebug,
		},
	); err != nil {
		return nil, err //nolint:wrapcheck // TODO
	}

	return cmd, nil
}
