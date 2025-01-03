package exec

import (
	"os/exec"

	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
)

const _optionEnvKey = "option-env"

func WithEnv(env []string) itbasisMiddlewareOption.Option[exec.Cmd] {
	return &optionEnv{env: env}
}
func WithRestoreEnv(env []string) itbasisMiddlewareOption.RestoreOption[exec.Cmd] {
	return &optionEnv{env: env, restore: true}
}

type optionEnv struct {
	restore bool

	env  []string
	prev []string
}

func (r *optionEnv) Key() itbasisMiddlewareOption.Key { return _optionEnvKey }
func (r *optionEnv) Apply(cmd *exec.Cmd) error {
	cmd.Env = r.env

	return nil
}

func (r *optionEnv) Save(cmd *exec.Cmd) error {
	if r.restore {
		r.prev = cmd.Env
	}

	return nil
}
func (r *optionEnv) Restore(cmd *exec.Cmd) error {
	if r.restore {
		cmd.Env = r.env
	}

	return nil
}
