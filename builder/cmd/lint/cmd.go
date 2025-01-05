package lint

import (
	"os/exec"

	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareExec "github.com/itbasis/tools/middleware/exec"
	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	"github.com/spf13/cobra"
)

var (
	_flagSkipEditorConfigChecker bool
	_flagSkipGolangCiLint        bool
)

var CmdLint = &cobra.Command{
	Use:  itbasisMiddlewareCmd.BuildUse("lint", itbasisMiddlewareCmd.UseFlags),
	Args: cobra.NoArgs,
	Run: itbasisMiddlewareCmd.WrapActionLogging(
		func(cmd *cobra.Command, _ []string) {
			withCobraOut := itbasisMiddlewareExec.WithCobraOut(cmd)

			if !_flagSkipEditorConfigChecker || itbasisMiddlewareOs.BeARegularFile(".editorconfig") {
				itbasisMiddlewareCmd.RequireNoError(cmd, _execEditorConfigChecker(withCobraOut))
			}

			if !_flagSkipGolangCiLint {
				itbasisMiddlewareCmd.RequireNoError(cmd, _execGolangCiLint(withCobraOut))
			}
		},
	),
}

func init() {
	CmdLint.Flags().BoolVar(&_flagSkipEditorConfigChecker, "skip-editor-config-checker", false, "skip editor config checker")
	CmdLint.Flags().BoolVar(&_flagSkipGolangCiLint, "skip-golangci-lint", false, "skip golangci-lint")
}

func _execEditorConfigChecker(opts ...itbasisMiddlewareOption.Option[exec.Cmd]) error {
	executable, err := itbasisMiddlewareExec.NewExecutable("editorconfig-checker", opts...)
	if err != nil {
		return err
	}

	return executable.Execute()
}

func _execGolangCiLint(opts ...itbasisMiddlewareOption.Option[exec.Cmd]) error {
	executable, err := itbasisMiddlewareExec.NewExecutable(
		"golangci-lint",
		append(
			[]itbasisMiddlewareOption.Option[exec.Cmd]{itbasisMiddlewareExec.WithArgs("run")},
			opts...,
		)...,
	)
	if err != nil {
		return err
	}

	return executable.Execute()
}
