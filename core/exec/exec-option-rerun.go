package exec

import (
	"os/exec"

	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
)

const _optionRerunKey itbasisCoreOption.Key = "option-rerun"

func WithRerun() RestoreOption { return &optionRerun{} }

type optionRerun struct{}

func (r *optionRerun) Key() itbasisCoreOption.Key { return _optionRerunKey }

func (r *optionRerun) Apply(_ *exec.Cmd) error { return nil }

func (r *optionRerun) Save(_ *exec.Cmd) error { return nil }

func (r *optionRerun) Restore(cmd *exec.Cmd) error {
	cmd.Process = nil
	cmd.ProcessState = nil

	return nil
}
