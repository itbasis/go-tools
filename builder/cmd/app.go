package cmd

import (
	"log"

	builderCmdBuild "github.com/itbasis/go-tools/builder/cmd/build"
	builderCmdDependencies "github.com/itbasis/go-tools/builder/cmd/dependencies"
	builderCmdGenerate "github.com/itbasis/go-tools/builder/cmd/generate"
	builderCmdLint "github.com/itbasis/go-tools/builder/cmd/lint"
	builderCmdTest "github.com/itbasis/go-tools/builder/cmd/test"
	builderCmdUpdate "github.com/itbasis/go-tools/builder/cmd/update"
	itbasisCoreApp "github.com/itbasis/go-tools/core/app"
	itbasisCoreCmd "github.com/itbasis/go-tools/core/cmd"
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
