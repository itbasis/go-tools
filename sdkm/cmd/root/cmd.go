package root

import (
	itbasisCoreCmd "github.com/itbasis/go-tools/core/cmd"
	sdkmCmdCurrent "github.com/itbasis/go-tools/sdkm/cmd/current"
	sdkmCmdEnv "github.com/itbasis/go-tools/sdkm/cmd/env"
	sdkmCmdExec "github.com/itbasis/go-tools/sdkm/cmd/exec"
	sdkmCmdInstall "github.com/itbasis/go-tools/sdkm/cmd/install"
	sdkmCmdLatest "github.com/itbasis/go-tools/sdkm/cmd/latest"
	sdkmCmdList "github.com/itbasis/go-tools/sdkm/cmd/list"
	sdkmCmdPlugins "github.com/itbasis/go-tools/sdkm/cmd/plugins"
	sdkmCmdReshim "github.com/itbasis/go-tools/sdkm/cmd/reshim"
	sdkmCmd "github.com/itbasis/go-tools/sdkm/internal/cmd"
	"github.com/spf13/cobra"
)

func NewRootCommand() (*cobra.Command, error) {
	var cmdRoot, err = itbasisCoreCmd.InitDefaultCmdRoot("SDK Manager")
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
