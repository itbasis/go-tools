package cmd

import (
	"log"

	itbasisCoreApp "github.com/itbasis/go-tools/core/app"
	"github.com/itbasis/go-tools/sdkm/cmd/root"
)

func InitApp() *itbasisCoreApp.App {
	var cmdRoot, err = root.NewRootCommand()
	if err != nil {
		log.Fatal(err)
	}

	return itbasisCoreApp.NewApp(cmdRoot)
}
