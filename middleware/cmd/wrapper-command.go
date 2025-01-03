package cmd

import (
	"log/slog"

	"github.com/itbasis/tools/middleware/log"
	"github.com/spf13/cobra"
)

type ActionFunc func(cmd *cobra.Command, args []string)

func WrapActionLogging(action ActionFunc) ActionFunc {
	return func(cmd *cobra.Command, args []string) {
		slog.Debug(
			"execute command",
			slog.String("command", cmd.Name()),
			log.SlogAttrStringsWithSeparator("args", "", args),
		)

		action(cmd, args)
	}
}

func MultipleActions(actions ...ActionFunc) ActionFunc {
	return func(cmd *cobra.Command, args []string) {
		for _, action := range actions {
			if action != nil {
				action(cmd, args)
			}
		}
	}
}
