package list

import "github.com/spf13/cobra"

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List installed versions",
	}

	cmd.AddCommand(newListAllCommand())

	return cmd
}
