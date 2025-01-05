package app

import (
	itbasisMiddlewareLog "github.com/itbasis/tools/middleware/log"
	"github.com/spf13/cobra"
)

type App struct {
	cmdRoot *cobra.Command
}

func NewApp(cmdRoot *cobra.Command) *App {
	itbasisMiddlewareLog.InitDefaultLoggerWIthConsole(cmdRoot.OutOrStdout())

	return &App{
		cmdRoot: cmdRoot,
	}
}

func (app *App) Run(args ...string) {
	if len(args) > 0 {
		app.cmdRoot.SetArgs(args)
	}

	_ = app.cmdRoot.Execute()
}
