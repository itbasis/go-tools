package cmd

import (
	"log"

	itbasisMiddlewareApp "github.com/itbasis/tools/middleware/app"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	sdkmCmdCurrent "github.com/itbasis/tools/sdkm/cmd/current"
	sdkmCmdEnv "github.com/itbasis/tools/sdkm/cmd/env"
	sdkmCmdExec "github.com/itbasis/tools/sdkm/cmd/exec"
	sdkmCmdInstall "github.com/itbasis/tools/sdkm/cmd/install"
	sdkmCmdLatest "github.com/itbasis/tools/sdkm/cmd/latest"
	sdkmCmdList "github.com/itbasis/tools/sdkm/cmd/list"
	sdkmCmdPlugins "github.com/itbasis/tools/sdkm/cmd/plugins"
	sdkmCmdReshim "github.com/itbasis/tools/sdkm/cmd/reshim"
)

func InitApp() *itbasisMiddlewareApp.App {
	var cmdRoot, err = itbasisMiddlewareCmd.InitDefaultCmdRoot("SDK Manager")
	if err != nil {
		log.Fatal(err)
	}

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

	return itbasisMiddlewareApp.NewApp(cmdRoot)
}
