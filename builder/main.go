package main

import (
	"log"

	"github.com/itbasis/tools/builder/cmd/build"
	"github.com/itbasis/tools/builder/cmd/dependencies"
	"github.com/itbasis/tools/builder/cmd/generate"
	"github.com/itbasis/tools/builder/cmd/lint"
	"github.com/itbasis/tools/builder/cmd/test"
	"github.com/itbasis/tools/builder/cmd/update"
	itbasisMiddlewareApp "github.com/itbasis/tools/middleware/app"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
)

func main() {
	var cmdRoot, err = itbasisMiddlewareCmd.InitDefaultCmdRoot("itbasis-builder")
	if err != nil {
		log.Fatal(err)
	}

	cmdRoot.AddCommand(
		dependencies.CmdDependencies,
		update.CmdUpdate,
		generate.CmdGenerate,
		lint.CmdLint,
		test.CmdUnitTest,
		build.NewBuildCommand(),
	)

	itbasisMiddlewareApp.NewApp(cmdRoot).Run()
}
