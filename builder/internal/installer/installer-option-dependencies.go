package installer

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/itbasis/go-tools/builder/internal/installer/model"
	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
)

const _optionDependenciesKey itbasisCoreOption.Key = "dependencies"

func WithJSONData(data []byte) Option {
	return &_optionDependencies{data: data}
}

func WithFile(filePath string) Option {
	return &_optionDependencies{filePath: filePath}
}

type _optionDependencies struct {
	filePath     string
	data         []byte
	dependencies model.Dependencies
}

func (r _optionDependencies) Key() itbasisCoreOption.Key { return _optionDependenciesKey }
func (r _optionDependencies) Apply(obj *Installer) error {
	var err error

	if len(r.filePath) > 0 {
		slog.Debug("using dependencies file: " + r.filePath)

		r.data, err = os.ReadFile(r.filePath)
		if err != nil {
			return err //nolint:wrapcheck // TODO
		}
	}

	if len(r.data) > 0 {
		if err = json.Unmarshal(r.data, &r.dependencies); err != nil {
			return err //nolint:wrapcheck // TODO
		}
	}

	obj.dependencies = r.dependencies

	return nil
}
