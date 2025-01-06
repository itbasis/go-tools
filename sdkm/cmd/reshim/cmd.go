package reshim

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	sdkmScripts "github.com/itbasis/tools/sdkm/scripts"
	"github.com/spf13/cobra"
)

func NewReshimCommand() *cobra.Command {
	return &cobra.Command{
		Use:   itbasisMiddlewareCmd.BuildUse("reshim"),
		Short: "Unpacking scripts and shims",
		RunE: func(_ *cobra.Command, _ []string) error {
			return sdkmScripts.Unpack(itbasisMiddlewareOs.ExecutableDir())
		},
	}
}
