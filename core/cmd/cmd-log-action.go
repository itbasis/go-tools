package cmd

import (
	"log/slog"

	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
	itbasisCoreOs "github.com/itbasis/go-tools/core/os"
	"github.com/spf13/cobra"
)

func LogCommand(cmd *cobra.Command, args []string) {
	slog.Debug(
		"execute command",
		itbasisCoreLog.SlogAttrCommand(
			cmd.Name(),
			itbasisCoreOs.Pwd(),
			args,
		),
	)
}
