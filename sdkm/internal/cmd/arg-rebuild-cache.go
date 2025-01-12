package cmd

import (
	cmd2 "github.com/itbasis/tools/middleware/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	_flagRebuildCache = "rebuild-cache"
)

func InitFlagRebuildCache(flags *pflag.FlagSet) {
	flags.Bool(_flagRebuildCache, false, "rebuild cache SDK versions")
}

func IsFlagRebuildCache(cmd *cobra.Command) bool {
	flag, err := cmd.Flags().GetBool(_flagRebuildCache)
	cmd2.RequireNoError(cmd, err)

	return flag
}
