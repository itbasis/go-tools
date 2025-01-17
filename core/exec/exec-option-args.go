package exec

import (
	"log/slog"
	"os/exec"

	itbasisCoreLog "github.com/itbasis/tools/core/log"
	itbasisCoreOption "github.com/itbasis/tools/core/option"
)

func WithArgs(args ...string) itbasisCoreOption.Option[exec.Cmd] {
	return &optionArgs{includePrevArgs: IncludePrevArgsNo, args: args}
}

func WithArgsIncludePrevious(includePrevArgs IncludePrevArgs, args ...string) itbasisCoreOption.Option[exec.Cmd] {
	return &optionArgs{includePrevArgs: includePrevArgs, args: args}
}

func WithRestoreArgs(args ...string) itbasisCoreOption.RestoreOption[exec.Cmd] {
	return &optionArgs{includePrevArgs: IncludePrevArgsNo, args: args, restore: true}
}

func WithRestoreArgsIncludePrevious(includePrevArgs IncludePrevArgs, args ...string) itbasisCoreOption.RestoreOption[exec.Cmd] {
	return &optionArgs{includePrevArgs: includePrevArgs, args: args, restore: true}
}

const _optionArgsKey = "option-args"

type IncludePrevArgs int

const (
	IncludePrevArgsNo     IncludePrevArgs = iota
	IncludePrevArgsBefore IncludePrevArgs = iota
	IncludePrevArgsAfter  IncludePrevArgs = iota
)

type optionArgs struct {
	includePrevArgs IncludePrevArgs
	restore         bool

	args []string
	prev []string
}

func (r *optionArgs) Key() itbasisCoreOption.Key { return _optionArgsKey }

func (r *optionArgs) Apply(cmd *exec.Cmd) error {
	switch r.includePrevArgs {
	case IncludePrevArgsNo:
		cmd.Args = append([]string{cmd.Path}, r.args...)

	case IncludePrevArgsBefore:
		cmd.Args = append(cmd.Args, r.args...)

	case IncludePrevArgsAfter:
		cmd.Args = append(append([]string{cmd.Path}, r.args...), cmd.Args[1:]...)

	default:
		return NewUnsupportedIncludePrevArgsError(r.includePrevArgs)
	}

	slog.Debug("applied args", itbasisCoreLog.SlogAttrSliceWithSeparator("args", " ", cmd.Args))

	return nil
}

func (r *optionArgs) Save(cmd *exec.Cmd) error {
	if r.restore {
		r.prev = cmd.Args
	}

	return nil
}

func (r *optionArgs) Restore(cmd *exec.Cmd) error {
	if r.restore {
		cmd.Args = r.prev
	}

	return nil
}
