package lint

import (
	"os/exec"

	builderCmd "github.com/itbasis/tools/builder/internal/cmd"
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	itbasisMiddlewareExec "github.com/itbasis/tools/middleware/exec"
	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	_flagSkipEditorConfigChecker bool
	_flagSkipGolangCiLint        bool
)

func NewLintCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    itbasisMiddlewareCmd.BuildUse("lint", builderCmd.UseArgPackages),
		Args:   cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
		PreRun: itbasisMiddlewareCmd.LogCommand,
		Run:    _run,
	}

	cmd.Flags().BoolVar(&_flagSkipEditorConfigChecker, "skip-editor-config-checker", false, "skip editor config checker")
	cmd.Flags().BoolVar(&_flagSkipGolangCiLint, "skip-golangci-lint", false, "skip golangci-lint")

	return cmd
}

func _run(cmd *cobra.Command, args []string) {
	withCobraOut := itbasisMiddlewareExec.WithCobraOut(cmd)

	if !_flagSkipEditorConfigChecker || itbasisMiddlewareOs.BeARegularFile(".editorconfig") {
		itbasisMiddlewareCmd.RequireNoError(cmd, _execEditorConfigChecker(withCobraOut))
	}

	if !_flagSkipGolangCiLint {
		itbasisMiddlewareCmd.RequireNoError(cmd, _execGolangCiLint(builderCmd.ArgPackages(builderCmd.DefaultPackages, args), withCobraOut))
	}
}

func _execEditorConfigChecker(opts ...itbasisMiddlewareOption.Option[exec.Cmd]) error {
	executable, err := itbasisMiddlewareExec.NewExecutable("editorconfig-checker", opts...)
	if err != nil {
		return errors.Wrap(err, itbasisMiddlewareExec.ErrFailedExecuteCommand.Error())
	}

	if err := executable.Execute(); err != nil {
		return errors.Wrap(err, itbasisMiddlewareExec.ErrFailedExecuteCommand.Error())
	}

	return nil
}

func _execGolangCiLint(lintPackages string, opts ...itbasisMiddlewareOption.Option[exec.Cmd]) error {
	executable, err := itbasisMiddlewareExec.NewExecutable(
		"golangci-lint",
		append(
			[]itbasisMiddlewareOption.Option[exec.Cmd]{itbasisMiddlewareExec.WithArgs("run", lintPackages)},
			opts...,
		)...,
	)
	if err != nil {
		return errors.Wrap(err, itbasisMiddlewareExec.ErrFailedExecuteCommand.Error())
	}

	return errors.Wrap(executable.Execute(), itbasisMiddlewareExec.ErrFailedExecuteCommand.Error())
}
