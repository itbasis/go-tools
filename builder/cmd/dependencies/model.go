package dependencies

import (
	_ "embed"
	"encoding/json"
)

//go:embed dependencies.json
var _data []byte

var _dependencies _model

func init() {
	if err := json.Unmarshal(_data, &_dependencies); err != nil {
		panic(err)
	}
}

type _model struct {
	GoInstall map[string]_goInstallDependency `json:"go_install"`
}

type _goInstallDependency struct {
	Package string
	Version string
}

func (r *_goInstallDependency) GoString() string { return r.ToString() }

func (r *_goInstallDependency) ToString() string {
	return r.Package + "@" + r.Version
}
