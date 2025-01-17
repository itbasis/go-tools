package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Fatal(cmd *cobra.Command, err error) {
	cmd.PrintErrln(err)

	os.Exit(1)
}

func RequireNoError(cmd *cobra.Command, err error) {
	if err == nil {
		return
	}

	Fatal(cmd, err)
}

func ExecuteRequireNoError(cmd *cobra.Command) {
	RequireNoError(cmd, cmd.Execute())
}
