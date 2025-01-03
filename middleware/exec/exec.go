package exec

import (
	"log/slog"
	"os/exec"

	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
)

type Executable struct {
	cmd *exec.Cmd

	cli string
}

func NewExecutable(cli string, opts ...itbasisMiddlewareOption.Option[exec.Cmd]) (*Executable, error) {
	var cmp = &Executable{
		cmd: exec.Command(cli),
	}

	if err := itbasisMiddlewareOption.ApplyOptions(
		cmp.cmd, opts, map[itbasisMiddlewareOption.Key]itbasisMiddlewareOption.LazyOptionFunc[exec.Cmd]{
			_optionOutKey: WithStdOut,
		},
	); err != nil {
		return nil, err
	}

	return cmp, nil
}

func (ge *Executable) Execute(opts ...itbasisMiddlewareOption.RestoreOption[exec.Cmd]) error {
	var (
		cmd = ge.cmd
		err error
	)

	if applyErr := itbasisMiddlewareOption.ApplyRestoreOptions(
		cmd, opts, func() {
			slog.Info("execute command", slog.String("command", cmd.String()))

			err = cmd.Run()
		},
	); applyErr != nil {
		return applyErr
	}

	return err
}
