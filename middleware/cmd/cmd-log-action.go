package cmd

import (
	"log/slog"

	"github.com/itbasis/tools/middleware/log"
	"github.com/spf13/cobra"
)

func LogCommand(cmd *cobra.Command, args []string) {
	slog.Debug(
		"execute command",
		slog.String("command", cmd.Name()),
		log.SlogAttrSliceWithSeparator("args", "", args),
	)
}
