package installer

import (
	"encoding/json"

	"golang.org/x/mod/module"
)

type Dependencies struct {
	GoDependencies map[DependencyName]module.Version `json:"go_install"`
}

type DependencyName = string

func ParseDependencies(data []byte) (Dependencies, error) {
	result := Dependencies{}

	if err := json.Unmarshal(data, &result); err != nil {
		return Dependencies{}, err
	}

	return result, nil
}
