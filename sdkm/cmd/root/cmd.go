package root

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	sdkmCmdCurrent "github.com/itbasis/tools/sdkm/cmd/current"
	sdkmCmdEnv "github.com/itbasis/tools/sdkm/cmd/env"
	sdkmCmdExec "github.com/itbasis/tools/sdkm/cmd/exec"
	sdkmCmdInstall "github.com/itbasis/tools/sdkm/cmd/install"
	sdkmCmdLatest "github.com/itbasis/tools/sdkm/cmd/latest"
	sdkmCmdList "github.com/itbasis/tools/sdkm/cmd/list"
	sdkmCmdPlugins "github.com/itbasis/tools/sdkm/cmd/plugins"
	sdkmCmdReshim "github.com/itbasis/tools/sdkm/cmd/reshim"
	sdkmCmd "github.com/itbasis/tools/sdkm/internal/cmd"
	"github.com/spf13/cobra"
)

func NewRootCommand() (*cobra.Command, error) {
	var cmdRoot, err = itbasisMiddlewareCmd.InitDefaultCmdRoot("SDK Manager")
	if err != nil {
		return nil, err //nolint:wrapcheck // TODO
	}

	sdkmCmd.InitFlagSdkRootDir(cmdRoot.PersistentFlags())

	cmdRoot.AddCommand(
		sdkmCmdPlugins.NewPluginsCommand(),
		sdkmCmdList.NewListCommand(),
		sdkmCmdLatest.NewLatestCommand(),
		sdkmCmdCurrent.NewCurrentCommand(),
		sdkmCmdInstall.NewInstallCommand(),
		sdkmCmdEnv.NewEnvCommand(),
		sdkmCmdExec.NewExecCommand(),
		sdkmCmdReshim.NewReshimCommand(),
	)

	return cmdRoot, nil
}
