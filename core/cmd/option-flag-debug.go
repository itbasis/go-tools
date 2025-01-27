package cmd

import (
	"log/slog"

	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const _optionDebugFlagKey = "option-debug-flag"

var _flagDebug = "debug"

func WithDefaultFlagDebug() itbasisCoreOption.Option[cobra.Command] {
	return WithFlagDebug(_flagDebug, true)
}

func WithFlagDebug(name string, persistent bool) itbasisCoreOption.Option[cobra.Command] {
	_flagDebug = name

	return &_optionDebugFlag{persistent: persistent}
}

type _optionDebugFlag struct {
	persistent bool

	flag bool
}

func (r *_optionDebugFlag) Key() itbasisCoreOption.Key { return _optionDebugFlagKey }

func (r *_optionDebugFlag) Apply(cmd *cobra.Command) error {
	var flags *pflag.FlagSet

	if r.persistent {
		flags = cmd.PersistentFlags()
	} else {
		flags = cmd.Flags()
	}

	flags.BoolVar(&r.flag, _flagDebug, false, "debug mode")

	cmd.PersistentPreRun = MultipleActions(r.setRootLogLevel, cmd.PersistentPreRun)

	return nil
}

func (r *_optionDebugFlag) setRootLogLevel(_ *cobra.Command, _ []string) {
	if r.flag {
		itbasisCoreLog.SetRootLogLevel(slog.LevelDebug)
	}
}
