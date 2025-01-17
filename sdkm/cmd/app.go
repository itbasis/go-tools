package cmd

import (
	"log"

	itbasisCoreApp "github.com/itbasis/tools/core/app"
	"github.com/itbasis/tools/sdkm/cmd/root"
)

func InitApp() *itbasisCoreApp.App {
	var cmdRoot, err = root.NewRootCommand()
	if err != nil {
		log.Fatal(err)
	}

	return itbasisCoreApp.NewApp(cmdRoot)
}
