package cmd

import (
	"log/slog"

	itbasisMiddlewareLog "github.com/itbasis/tools/middleware/log"
	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	FlagDebug = "debug"

	_optionDebugFlagKey = "option-debug-flag"
)

func WithDefaultFlagDebug() itbasisMiddlewareOption.Option[cobra.Command] {
	return WithFlagDebug(true)
}

func WithFlagDebug(persistent bool) itbasisMiddlewareOption.Option[cobra.Command] {
	return &_optionDebugFlag{persistent: persistent}
}

type _optionDebugFlag struct {
	persistent bool

	flag bool
}

func (_ *_optionDebugFlag) Key() itbasisMiddlewareOption.Key { return _optionDebugFlagKey }

func (r *_optionDebugFlag) Apply(cmd *cobra.Command) error {
	var flags *pflag.FlagSet

	if r.persistent {
		flags = cmd.PersistentFlags()
	} else {
		flags = cmd.Flags()
	}

	flags.BoolVar(&r.flag, FlagDebug, false, "debug mode")

	cmd.PersistentPreRun = MultipleActions(r.setRootLogLevel, cmd.PersistentPreRun)

	return nil
}

func (r *_optionDebugFlag) setRootLogLevel(_ *cobra.Command, _ []string) {
	if r.flag {
		itbasisMiddlewareLog.SetRootLogLevel(slog.LevelDebug)
	}
}
