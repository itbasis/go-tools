package exec_test

import (
	"os/exec"

	itbasisCoreExec "github.com/itbasis/go-tools/core/exec"
	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"Option Args", func() {
		var path = "test_path"

		ginkgo.Describe(
			"without restoring the previous state", func() {
				ginkgo.DescribeTable(
					"simple", func(args, optionArgs, wantArgs []string) {
						var cmd = exec.Command(path, args...)

						gomega.Expect(
							itbasisCoreOption.ApplyOptions(
								cmd,
								[]itbasisCoreExec.Option{itbasisCoreExec.WithArgs(optionArgs...)},
								nil,
							),
						).
							To(gomega.Succeed())
						gomega.Expect(cmd.Path).To(gomega.Equal(path))
						gomega.Expect(cmd.Args).To(gomega.HaveExactElements(wantArgs))
					},
					ginkgo.Entry(nil, []string{"arg0"}, []string{}, []string{path}),
					ginkgo.Entry(nil, []string{"arg0"}, []string{"arg1"}, []string{path, "arg1"}),
				)

				ginkgo.DescribeTable(
					"with preservation of path and arguments",
					func(includePrevArgs itbasisCoreExec.IncludePrevArgs, args, optionArgs, wantArgs []string) {
						var cmd = exec.Command(path, args...)

						gomega.Expect(
							itbasisCoreOption.ApplyOptions(
								cmd,
								[]itbasisCoreExec.Option{
									itbasisCoreExec.WithArgsIncludePrevious(includePrevArgs, optionArgs...),
								},
								nil,
							),
						).To(gomega.Succeed())
						gomega.Expect(cmd.Path).To(gomega.Equal(path))
						gomega.Expect(cmd.Args).To(gomega.HaveExactElements(wantArgs))
					},
					ginkgo.Entry(nil, itbasisCoreExec.IncludePrevArgsBefore, []string{"arg0"}, []string{}, []string{path, "arg0"}),
					ginkgo.Entry(
						nil,
						itbasisCoreExec.IncludePrevArgsBefore,
						[]string{"arg0"},
						[]string{"arg1"},
						[]string{path, "arg0", "arg1"},
					),
					ginkgo.Entry(
						nil,
						itbasisCoreExec.IncludePrevArgsAfter,
						[]string{"arg0"},
						[]string{"arg1"},
						[]string{path, "arg1", "arg0"},
					),
				)
			},
		)

		ginkgo.Describe(
			"with restoration of the previous state", func() {
				ginkgo.DescribeTable(
					"simple", func(args, optionArgs, wantArgs []string) {
						var cmd = exec.Command(path, args...)

						gomega.Expect(
							itbasisCoreOption.ApplyRestoreOptions(
								cmd,
								[]itbasisCoreExec.RestoreOption{itbasisCoreExec.WithRestoreArgs(optionArgs...)}, func() {
									gomega.Expect(cmd.Args).To(gomega.HaveExactElements(wantArgs))
								},
							),
						).To(gomega.Succeed())

						gomega.Expect(cmd.Path).To(gomega.Equal(path))
						gomega.Expect(cmd.Args).To(gomega.HaveExactElements(append([]string{path}, args...)))
					},
					ginkgo.Entry(nil, []string{"arg0"}, []string{}, []string{path}),
					ginkgo.Entry(
						nil, []string{"arg0"}, []string{"arg1"}, []string{path, "arg1"},
					),
				)

				ginkgo.DescribeTable(
					"with preservation of path and arguments",
					func(includePrevArgs itbasisCoreExec.IncludePrevArgs, args, optionArgs, wantArgs []string) {
						var cmd = exec.Command(path, args...)

						gomega.Expect(
							itbasisCoreOption.ApplyRestoreOptions(
								cmd,
								[]itbasisCoreExec.RestoreOption{itbasisCoreExec.WithRestoreArgsIncludePrevious(
									includePrevArgs,
									optionArgs...,
								)}, func() {
									gomega.Expect(cmd.Args).To(gomega.HaveExactElements(wantArgs))
								},
							),
						).To(gomega.Succeed())

						gomega.Expect(cmd.Path).To(gomega.Equal(path))
						gomega.Expect(cmd.Args).To(gomega.HaveExactElements(append([]string{path}, args...)))
					},
					ginkgo.Entry(nil, itbasisCoreExec.IncludePrevArgsNo, []string{"arg0"}, []string{}, []string{path}),
					ginkgo.Entry(
						nil, itbasisCoreExec.IncludePrevArgsNo, []string{"arg0"}, []string{"arg1"}, []string{path, "arg1"},
					),
					ginkgo.Entry(nil, itbasisCoreExec.IncludePrevArgsBefore, []string{"arg0"}, []string{}, []string{path, "arg0"}),
					ginkgo.Entry(
						nil, itbasisCoreExec.IncludePrevArgsBefore, []string{"arg0"}, []string{"arg1"}, []string{path, "arg0", "arg1"},
					),
					ginkgo.Entry(
						nil, itbasisCoreExec.IncludePrevArgsAfter, []string{"arg0"}, []string{"arg1"}, []string{path, "arg1", "arg0"},
					),
				)

			},
		)
	},
)
