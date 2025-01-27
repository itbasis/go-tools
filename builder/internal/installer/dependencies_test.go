package installer_test

import (
	"github.com/itbasis/go-test-utils/v5/files"
	builderInstaller "github.com/itbasis/go-tools/builder/internal/installer"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"golang.org/x/mod/module"
	"golang.org/x/tools/godoc/vfs"
)

var _ = ginkgo.Describe(
	"Unmarshal", func() {
		data := files.ReadFile(vfs.OS("."), "sample.json")

		gomega.Expect(builderInstaller.ParseDependencies(data)).To(
			gomega.Equal(
				builderInstaller.Dependencies{
					GoDependencies: map[builderInstaller.DependencyName]module.Version{
						"mockgen": {
							Path:    "go.uber.org/mock/mockgen",
							Version: "latest",
						},
					},
					GithubDependencies: map[builderInstaller.DependencyName]builderInstaller.GithubDependency{
						"golangci-lint": {
							Owner:   "golangci",
							Repo:    "golangci-lint",
							Version: "latest",
						},
					},
				},
			),
		)
	},
)
