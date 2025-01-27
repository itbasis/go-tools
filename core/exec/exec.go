package exec

import (
	"log/slog"
	"os/exec"

	"github.com/itbasis/go-tools/core/log"
	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
)

type Executable struct {
	cmd *exec.Cmd

	cli string
}

func NewExecutable(cli string, opts ...itbasisCoreOption.Option[exec.Cmd]) (*Executable, error) {
	var cmp = &Executable{
		cmd: exec.Command(cli),
	}

	if err := itbasisCoreOption.ApplyOptions(
		cmp.cmd, opts, map[itbasisCoreOption.Key]itbasisCoreOption.LazyOptionFunc[exec.Cmd]{
			_optionInKey:  WithStdIn,
			_optionOutKey: WithStdOut,
		},
	); err != nil {
		return nil, err //nolint:wrapcheck // TODO
	}

	return cmp, nil
}

func (ge *Executable) Execute(opts ...itbasisCoreOption.RestoreOption[exec.Cmd]) error {
	var (
		cmd = ge.cmd
		err error
	)

	if applyErr := itbasisCoreOption.ApplyRestoreOptions(
		cmd, opts, func() {
			slog.Debug("execute external program", log.SlogAttrCommand(cmd.Path, cmd.Args[1:]...))

			err = cmd.Run()
		},
	); applyErr != nil {
		return applyErr //nolint:wrapcheck // TODO
	}

	return err //nolint:wrapcheck // TODO
}
