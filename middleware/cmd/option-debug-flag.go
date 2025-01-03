package cmd

import (
	"log/slog"

	itbasisMiddlewareLog "github.com/itbasis/tools/middleware/log"
	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
	"github.com/spf13/cobra"
)

const _optionDebugFlagKey = "option-debug-flag"

func WithDefaultDebugAction() itbasisMiddlewareOption.Option[cobra.Command] {
	return &_optionDebugFlag{}
}

type _optionDebugFlag struct {
	flag *bool
}

func (r *_optionDebugFlag) Key() itbasisMiddlewareOption.Key { return _optionDebugFlagKey }
func (r *_optionDebugFlag) Apply(cmd *cobra.Command) error {
	r.flag = cmd.PersistentFlags().BoolP("debug", "d", false, "debug mode")

	cmd.PersistentPreRun = MultipleActions(r.setRootLogLevel, cmd.PersistentPreRun)

	return nil
}

func (r *_optionDebugFlag) setRootLogLevel(_ *cobra.Command, _ []string) {
	if *r.flag {
		itbasisMiddlewareLog.SetRootLogLevel(slog.LevelDebug)
	}
}
