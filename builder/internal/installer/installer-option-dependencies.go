package installer

import (
	"log/slog"
	"os"

	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
)

const _optionDependenciesKey = "dependencies"

func WithJSONData(data []byte) Option {
	return &_optionDependencies{data: data}
}

func WithFile(filePath string) Option {
	return &_optionDependencies{filePath: filePath}
}

type _optionDependencies struct {
	filePath     string
	data         []byte
	dependencies Dependencies
}

func (r _optionDependencies) Key() itbasisCoreOption.Key { return _optionDependenciesKey }
func (r _optionDependencies) Apply(obj *Installer) error {
	var err error

	if len(r.filePath) > 0 {
		slog.Debug("using dependencies file: " + r.filePath)

		r.data, err = os.ReadFile(r.filePath)
		if err != nil {
			return err
		}
	}

	if len(r.data) > 0 {
		r.dependencies, err = ParseDependencies(r.data)
		if err != nil {
			return err
		}
	}

	obj.dependencies = r.dependencies

	return nil
}
