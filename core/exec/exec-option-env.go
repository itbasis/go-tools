package exec

import (
	"os/exec"

	itbasisCoreEnv "github.com/itbasis/go-tools/core/env"
	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
)

const _optionEnvKey itbasisCoreOption.Key = "option-env"

func WithEnv[E itbasisCoreEnv.ListOrMap](env E) Option {
	return &_optionEnv[E]{env: env}
}
func WithRestoreEnv[E itbasisCoreEnv.ListOrMap](env E) RestoreOption {
	return &_optionEnv[E]{env: env, restore: true}
}

type _optionEnv[E itbasisCoreEnv.ListOrMap] struct {
	restore bool

	env  E
	prev itbasisCoreEnv.List
}

func (r *_optionEnv[E]) Key() itbasisCoreOption.Key { return _optionEnvKey }
func (r *_optionEnv[E]) Apply(cmd *exec.Cmd) error {
	switch v := any(r.env).(type) {
	case itbasisCoreEnv.List:
		cmd.Env = v

	case itbasisCoreEnv.Map:
		cmd.Env = itbasisCoreEnv.MapToSlices(v)
	}

	return nil
}

func (r *_optionEnv[E]) Save(cmd *exec.Cmd) error {
	if r.restore {
		r.prev = cmd.Env
	}

	return nil
}
func (r *_optionEnv[E]) Restore(cmd *exec.Cmd) error {
	if r.restore {
		cmd.Env = r.prev
	}

	return nil
}
