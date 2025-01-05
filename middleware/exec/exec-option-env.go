package exec

import (
	"os/exec"

	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
)

const _optionEnvKey = "option-env"

func WithEnv(env []string) itbasisMiddlewareOption.Option[exec.Cmd] {
	return &_optionEnv{env: env}
}
func WithRestoreEnv(env []string) itbasisMiddlewareOption.RestoreOption[exec.Cmd] {
	return &_optionEnv{env: env, restore: true}
}

type _optionEnv struct {
	restore bool

	env  []string
	prev []string
}

func (r *_optionEnv) Key() itbasisMiddlewareOption.Key { return _optionEnvKey }
func (r *_optionEnv) Apply(cmd *exec.Cmd) error {
	cmd.Env = r.env

	return nil
}

func (r *_optionEnv) Save(cmd *exec.Cmd) error {
	if r.restore {
		r.prev = cmd.Env
	}

	return nil
}
func (r *_optionEnv) Restore(cmd *exec.Cmd) error {
	if r.restore {
		cmd.Env = r.env
	}

	return nil
}
