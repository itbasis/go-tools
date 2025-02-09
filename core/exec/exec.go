package exec

import (
	"log/slog"
	"os/exec"

	itbasisCoreEnv "github.com/itbasis/go-tools/core/env"
	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
)

type Executable struct {
	cmd *exec.Cmd

	cli string
}

func NewExecutable(cli string, opts ...Option) (*Executable, error) {
	var cmp = &Executable{
		cmd: exec.Command(cli),
	}

	if err := itbasisCoreOption.ApplyOptions(
		cmp.cmd, opts, map[itbasisCoreOption.Key]itbasisCoreOption.LazyOptionFunc[exec.Cmd]{
			_optionWorkDirKey: WithOsPwd,
			_optionInKey:      WithStdIn,
			_optionOutKey:     WithStdOut,
		},
	); err != nil {
		return nil, err //nolint:wrapcheck // TODO
	}

	return cmp, nil
}

func (ge *Executable) Execute(opts ...RestoreOption) error {
	var (
		cmd = ge.cmd
		err error
	)

	if applyErr := itbasisCoreOption.ApplyRestoreOptions(
		cmd, opts, func() {
			slog.Debug(
				"execute external program",
				itbasisCoreLog.SlogAttrCommand(cmd.Path, cmd.Dir, cmd.Args[1:], itbasisCoreEnv.SlogAttrEnv(cmd.Env)),
			)

			err = cmd.Run()
		},
	); applyErr != nil {
		return applyErr //nolint:wrapcheck // TODO
	}

	return err //nolint:wrapcheck // TODO
}
