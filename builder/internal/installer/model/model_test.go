package model_test

import (
	"encoding/json"
	"path"

	itbasisTestUtilsFiles "github.com/itbasis/go-test-utils/v5/files"
	"github.com/itbasis/go-tools/builder/internal/installer/model"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"golang.org/x/mod/module"
	"golang.org/x/tools/godoc/vfs"
)

var _ = ginkgo.Describe(
	"Unmarshal", func() {
		defer ginkgo.GinkgoRecover()

		var dependencies model.Dependencies

		gomega.Expect(
			json.Unmarshal(
				itbasisTestUtilsFiles.ReadFile(vfs.OS("testdata"), "000.json"),
				&dependencies,
			),
		).To(gomega.Succeed())

		gomega.Expect(dependencies).To(
			gstruct.MatchAllFields(
				gstruct.Fields{
					"Go": gstruct.MatchAllKeys(
						gstruct.Keys{
							"mockgen": gomega.Equal(
								module.Version{
									Path:    "go.uber.org/mock/mockgen",
									Version: model.VersionLatest,
								},
							),
						},
					),
					"Github": gstruct.MatchAllKeys(
						gstruct.Keys{
							"golangci-lint": gomega.Equal(
								model.GithubDependency{
									Owner:   "golangci",
									Repo:    "golangci-lint",
									Version: model.VersionLatest,
								},
							),
						},
					),
				},
			),
		)
	},
)

var _ = ginkgo.Describe(
	"Unmarshal default dependencies", func() {
		defer ginkgo.GinkgoRecover()

		var dependencies model.Dependencies
		gomega.Expect(
			json.Unmarshal(
				itbasisTestUtilsFiles.ReadFile(
					vfs.OS(path.Join("..", "..", "..", "cmd", "dependencies")),
					"dependencies.json",
				),
				&dependencies,
			),
		).To(gomega.Succeed())
	},
)
