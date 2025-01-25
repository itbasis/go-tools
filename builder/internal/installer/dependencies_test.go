package installer

import (
	"github.com/itbasis/go-test-utils/v5/files"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"golang.org/x/mod/module"
	"golang.org/x/tools/godoc/vfs"
)

var _ = ginkgo.Describe(
	"Unmarshal", func() {
		data := files.ReadFile(vfs.OS("."), "sample.json")

		gomega.Expect(ParseDependencies(data)).To(
			gomega.Equal(
				Dependencies{
					GoDependencies: map[DependencyName]module.Version{
						"mockgen": {
							Path:    "go.uber.org/mock/mockgen",
							Version: "latest",
						},
					},
				},
			),
		)
	},
)
