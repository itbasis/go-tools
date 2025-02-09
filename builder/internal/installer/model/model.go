package model

import "golang.org/x/mod/module"

const VersionLatest = "latest"

type (
	ProviderKey    string
	DependencyName = string

	Dependencies struct {
		Go     GoDependencyList     `json:"go"`
		Github GithubDependencyList `json:"github"`
	}

	GoDependencyList     = map[DependencyName]module.Version
	GithubDependencyList = map[DependencyName]GithubDependency

	GithubDependency struct {
		Owner   string `json:"owner"`
		Repo    string `json:"repo"`
		Version string `json:"version"`
	}
)
