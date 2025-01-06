package cmd

import (
	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"

	"github.com/spf13/cobra"
)

func InitDefaultCmdRoot(shortName string, opts ...itbasisMiddlewareOption.Option[cobra.Command]) (*cobra.Command, error) {
	var cmd = &cobra.Command{Short: shortName}

	if err := itbasisMiddlewareOption.ApplyOptions(
		cmd,
		opts, map[itbasisMiddlewareOption.Key]itbasisMiddlewareOption.LazyOptionFunc[cobra.Command]{
			_optionVersionKey:   WithDefaultVersion,
			_optionOutKey:       WithDefaultOut,
			_optionErrKey:       WithDefaultErr,
			_optionDebugFlagKey: WithDefaultFlagDebug,
		},
	); err != nil {
		return nil, err //nolint:wrapcheck // _
	}

	return cmd, nil
}
