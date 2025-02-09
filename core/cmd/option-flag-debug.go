package cmd

import (
	"log/slog"
	"slices"

	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	_optionDebugFlagKey itbasisCoreOption.Key = "option-debug-flag"

	DefaultFlagDebugName        = "debug"
	DefaultFlagDebugDescription = "debug mode"
)

func WithDefaultFlagDebug() Option {
	return WithFlagDebug(DefaultFlagDebugName, DefaultFlagDebugDescription, true, false)
}

func WithFlagDebug(name, description string, persistent, ignoreDisableFlagParsing bool) Option {
	return &_optionDebugFlag{
		persistent:               persistent,
		ignoreDisableFlagParsing: ignoreDisableFlagParsing,
		name:                     name,
		description:              description,
	}
}

type _optionDebugFlag struct {
	persistent               bool
	ignoreDisableFlagParsing bool

	flag bool

	name        string
	description string
}

func (r *_optionDebugFlag) Key() itbasisCoreOption.Key { return _optionDebugFlagKey }

func (r *_optionDebugFlag) Apply(cmd *cobra.Command) error {
	var flags *pflag.FlagSet

	if r.persistent {
		flags = cmd.PersistentFlags()
		cmd.PersistentPreRun = MultipleActions(r.setRootLogLevel, cmd.PersistentPreRun)
	} else {
		flags = cmd.Flags()
		cmd.PreRun = MultipleActions(r.setRootLogLevel, cmd.PersistentPreRun)
	}

	flags.BoolVar(&r.flag, r.name, r.flag, r.description)

	return nil
}

func (r *_optionDebugFlag) setRootLogLevel(cmd *cobra.Command, args []string) {
	flagName := "--" + r.name

	if r.ignoreDisableFlagParsing && cmd.DisableFlagParsing {
		r.flag = slices.Contains(args, flagName)
	}

	if r.flag {
		itbasisCoreLog.SetRootLogLevel(slog.LevelDebug)
	}
}
