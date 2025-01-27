package installer

import (
	"encoding/json"

	"golang.org/x/mod/module"
)

const _versionLatest = "latest"

type Dependencies struct {
	GoDependencies     map[DependencyName]module.Version   `json:"go_install"`
	GithubDependencies map[DependencyName]GithubDependency `json:"github_install"`
}

type GithubDependency struct {
	Owner   string `json:"owner"`
	Repo    string `json:"repo"`
	Version string `json:"version"`
}

type DependencyName = string

func ParseDependencies(data []byte) (Dependencies, error) {
	result := Dependencies{}

	if err := json.Unmarshal(data, &result); err != nil {
		return Dependencies{}, err
	}

	return result, nil
}
