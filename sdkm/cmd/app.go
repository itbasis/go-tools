package cmd

import (
	"log"

	itbasisMiddlewareApp "github.com/itbasis/tools/middleware/app"
	"github.com/itbasis/tools/sdkm/cmd/root"
)

func InitApp() *itbasisMiddlewareApp.App {
	var cmdRoot, err = root.NewRootCommand()
	if err != nil {
		log.Fatal(err)
	}

	return itbasisMiddlewareApp.NewApp(cmdRoot)
}
