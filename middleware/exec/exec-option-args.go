package exec

import (
	"log/slog"
	"os/exec"

	"github.com/itbasis/tools/middleware/log"
	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
)

func WithArgs(args ...string) itbasisMiddlewareOption.Option[exec.Cmd] {
	return &optionArgs{includePrevArgs: IncludePrevArgsNo, args: args}
}

func WithArgsIncludePrevious(includePrevArgs IncludePrevArgs, args ...string) itbasisMiddlewareOption.Option[exec.Cmd] {
	return &optionArgs{includePrevArgs: includePrevArgs, args: args}
}

func WithRestoreArgs(args ...string) itbasisMiddlewareOption.RestoreOption[exec.Cmd] {
	return &optionArgs{includePrevArgs: IncludePrevArgsNo, args: args, restore: true}
}

func WithRestoreArgsIncludePrevious(includePrevArgs IncludePrevArgs, args ...string) itbasisMiddlewareOption.RestoreOption[exec.Cmd] {
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

func (r *optionArgs) Key() itbasisMiddlewareOption.Key { return _optionArgsKey }

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

	slog.Debug("applied args", log.SlogAttrSliceWithSeparator("args", " ", cmd.Args))

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
