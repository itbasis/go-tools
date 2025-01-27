package cmd

import (
	itbasisCoreCmd "github.com/itbasis/go-tools/core/cmd"
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
	itbasisCoreCmd.RequireNoError(cmd, err)

	return flag
}
