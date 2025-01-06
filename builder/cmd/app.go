package cmd

import (
	"log"

	builderCmdBuild "github.com/itbasis/tools/builder/cmd/build"
	builderCmdDependencies "github.com/itbasis/tools/builder/cmd/dependencies"
	builderCmdGenerate "github.com/itbasis/tools/builder/cmd/generate"
	builderCmdLint "github.com/itbasis/tools/builder/cmd/lint"
	builderCmdTest "github.com/itbasis/tools/builder/cmd/test"
	builderCmdUpdate "github.com/itbasis/tools/builder/cmd/update"
	itbasisMiddlewareApp "github.com/itbasis/tools/middleware/app"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
)

func InitApp() *itbasisMiddlewareApp.App {
	var cmdRoot, err = itbasisMiddlewareCmd.InitDefaultCmdRoot("itbasis-builder")
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

	return itbasisMiddlewareApp.NewApp(cmdRoot)
}
