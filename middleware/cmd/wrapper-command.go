package cmd

import (
	"github.com/spf13/cobra"
)

type ActionFunc func(cmd *cobra.Command, args []string)

func MultipleActions(actions ...ActionFunc) ActionFunc {
	return func(cmd *cobra.Command, args []string) {
		for _, action := range actions {
			if action != nil {
				action(cmd, args)
			}
		}
	}
}
