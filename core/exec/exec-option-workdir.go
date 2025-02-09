package exec

import (
	"os/exec"

	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
	itbasisCoreOs "github.com/itbasis/go-tools/core/os"
)

const _optionWorkDirKey itbasisCoreOption.Key = "itbasisCoreOption-workdir"

func WithOsPwd() Option {
	return &_optionWorkDir{dir: itbasisCoreOs.Pwd()}
}
func WithWorkDir(dir string) Option {
	return &_optionWorkDir{dir: dir}
}

type _optionWorkDir struct {
	dir string
}

func (r *_optionWorkDir) Key() itbasisCoreOption.Key { return _optionWorkDirKey }
func (r *_optionWorkDir) Apply(cmd *exec.Cmd) error {
	cmd.Dir = r.dir

	return nil
}
