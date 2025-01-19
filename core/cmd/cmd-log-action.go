package cmd

import (
	"log/slog"

	itbasisCoreLog "github.com/itbasis/tools/core/log"
	"github.com/spf13/cobra"
)

func LogCommand(cmd *cobra.Command, args []string) {
	slog.Debug("execute command", itbasisCoreLog.SlogAttrCommand(cmd.Name(), args...))
}
