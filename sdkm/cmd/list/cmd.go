package list

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "list",
		Short:  "List installed versions",
		PreRun: itbasisMiddlewareCmd.LogCommand,
	}

	cmd.AddCommand(newListAllCommand())

	return cmd
}
