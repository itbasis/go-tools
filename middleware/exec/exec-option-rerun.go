package exec

import (
	"os/exec"

	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
)

const _optionRerunKey = "option-rerun"

func WithRerun() itbasisMiddlewareOption.RestoreOption[exec.Cmd] { return &optionRerun{} }

type optionRerun struct{}

func (r *optionRerun) Key() itbasisMiddlewareOption.Key { return _optionRerunKey }
func (r *optionRerun) Apply(cmd *exec.Cmd) error        { return nil }
func (r *optionRerun) Save(cmd *exec.Cmd) error         { return nil }
func (r *optionRerun) Restore(cmd *exec.Cmd) error {
	cmd.Process = nil
	cmd.ProcessState = nil

	return nil
}
