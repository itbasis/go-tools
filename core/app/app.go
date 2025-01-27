package app

import (
	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
	"github.com/spf13/cobra"
)

type App struct {
	cmdRoot *cobra.Command
}

func NewApp(cmdRoot *cobra.Command) *App {
	itbasisCoreLog.InitDefaultLoggerWithConsole(cmdRoot.OutOrStdout())

	return &App{
		cmdRoot: cmdRoot,
	}
}

// Run If no arguments were passed, "os.Args" will be used.
func (app *App) Run(args ...string) {
	if len(args) > 0 {
		app.cmdRoot.SetArgs(args)
	}

	_ = app.cmdRoot.Execute()
}
