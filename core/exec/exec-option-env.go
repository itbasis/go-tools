package exec

import (
	"os/exec"

	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
	itbasisCoreOs "github.com/itbasis/go-tools/core/os"
)

const _optionEnvKey = "option-env"

func WithEnv(env []string) itbasisCoreOption.Option[exec.Cmd] {
	return &_optionEnv{env: env}
}
func WithEnvAsMap(env map[string]string) itbasisCoreOption.Option[exec.Cmd] {
	return &_optionEnv{env: itbasisCoreOs.EnvMapToSlices(env)}
}
func WithRestoreEnv(env []string) itbasisCoreOption.RestoreOption[exec.Cmd] {
	return &_optionEnv{env: env, restore: true}
}

type _optionEnv struct {
	restore bool

	env  []string
	prev []string
}

func (r *_optionEnv) Key() itbasisCoreOption.Key { return _optionEnvKey }
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
