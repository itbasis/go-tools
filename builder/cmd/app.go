package cmd

import (
	"log"

	builderCmdBuild "github.com/itbasis/tools/builder/cmd/build"
	builderCmdDependencies "github.com/itbasis/tools/builder/cmd/dependencies"
	builderCmdGenerate "github.com/itbasis/tools/builder/cmd/generate"
	builderCmdLint "github.com/itbasis/tools/builder/cmd/lint"
	builderCmdTest "github.com/itbasis/tools/builder/cmd/test"
	builderCmdUpdate "github.com/itbasis/tools/builder/cmd/update"
	itbasisCoreApp "github.com/itbasis/tools/core/app"
	itbasisCoreCmd "github.com/itbasis/tools/core/cmd"
)

func InitApp() *itbasisCoreApp.App {
	var cmdRoot, err = itbasisCoreCmd.InitDefaultCmdRoot("itbasis-builder")
	if err != nil {
		log.Fatal(err)
	}

	cmdRoot.AddCommand(
		builderCmdDependencies.NewDependenciesCommand(),
		builderCmdUpdate.NewUpdateCommand(),
		builderCmdGenerate.NewGenerateCommand(),
		builderCmdLint.NewLintCommand(),
		builderCmdTest.NewUnitTestCommand(),
		builderCmdBuild.NewBuildCommand(),
	)

	return itbasisCoreApp.NewApp(cmdRoot)
}
