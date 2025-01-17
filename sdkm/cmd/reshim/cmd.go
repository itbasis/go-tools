package reshim

import (
	itbasisCoreCmd "github.com/itbasis/tools/core/cmd"
	itbasisCoreOs "github.com/itbasis/tools/core/os"
	sdkmScripts "github.com/itbasis/tools/sdkm/scripts"
	"github.com/spf13/cobra"
)

func NewReshimCommand() *cobra.Command {
	return &cobra.Command{
		Use:    itbasisCoreCmd.BuildUse("reshim"),
		Short:  "Unpacking scripts and shims",
		PreRun: itbasisCoreCmd.LogCommand,
		RunE: func(_ *cobra.Command, _ []string) error {
			return sdkmScripts.Unpack(itbasisCoreOs.ExecutableDir())
		},
	}
}
