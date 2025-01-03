package exec_test

import (
	"os/exec"

	itbasisMiddlewareExec "github.com/itbasis/tools/middleware/exec"
	"github.com/itbasis/tools/middleware/option"
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

						gomega.Expect(option.ApplyOptions(cmd, []option.Option[exec.Cmd]{itbasisMiddlewareExec.WithArgs(optionArgs...)}, nil)).
							To(gomega.Succeed())
						gomega.Expect(cmd.Path).To(gomega.Equal(path))
						gomega.Expect(cmd.Args).To(gomega.HaveExactElements(wantArgs))
					},
					ginkgo.Entry(nil, []string{"arg0"}, []string{}, []string{path}),
					ginkgo.Entry(nil, []string{"arg0"}, []string{"arg1"}, []string{path, "arg1"}),
				)

				ginkgo.DescribeTable(
					"with preservation of path and arguments",
					func(includePrevArgs itbasisMiddlewareExec.IncludePrevArgs, args, optionArgs, wantArgs []string) {
						var cmd = exec.Command(path, args...)

						gomega.Expect(
							option.ApplyOptions(
								cmd,
								[]option.Option[exec.Cmd]{
									itbasisMiddlewareExec.WithArgsIncludePrevious(includePrevArgs, optionArgs...),
								},
								nil,
							),
						).To(gomega.Succeed())
						gomega.Expect(cmd.Path).To(gomega.Equal(path))
						gomega.Expect(cmd.Args).To(gomega.HaveExactElements(wantArgs))
					},
					ginkgo.Entry(nil, itbasisMiddlewareExec.IncludePrevArgsBefore, []string{"arg0"}, []string{}, []string{path, "arg0"}),
					ginkgo.Entry(
						nil,
						itbasisMiddlewareExec.IncludePrevArgsBefore,
						[]string{"arg0"},
						[]string{"arg1"},
						[]string{path, "arg0", "arg1"},
					),
					ginkgo.Entry(
						nil,
						itbasisMiddlewareExec.IncludePrevArgsAfter,
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
							option.ApplyRestoreOptions(
								cmd,
								[]option.RestoreOption[exec.Cmd]{itbasisMiddlewareExec.WithRestoreArgs(optionArgs...)}, func() {
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
					func(includePrevArgs itbasisMiddlewareExec.IncludePrevArgs, args, optionArgs, wantArgs []string) {
						var cmd = exec.Command(path, args...)

						gomega.Expect(
							option.ApplyRestoreOptions(
								cmd,
								[]option.RestoreOption[exec.Cmd]{itbasisMiddlewareExec.WithRestoreArgsIncludePrevious(
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
					ginkgo.Entry(nil, itbasisMiddlewareExec.IncludePrevArgsNo, []string{"arg0"}, []string{}, []string{path}),
					ginkgo.Entry(
						nil, itbasisMiddlewareExec.IncludePrevArgsNo, []string{"arg0"}, []string{"arg1"}, []string{path, "arg1"},
					),
					ginkgo.Entry(nil, itbasisMiddlewareExec.IncludePrevArgsBefore, []string{"arg0"}, []string{}, []string{path, "arg0"}),
					ginkgo.Entry(
						nil, itbasisMiddlewareExec.IncludePrevArgsBefore, []string{"arg0"}, []string{"arg1"}, []string{path, "arg0", "arg1"},
					),
					ginkgo.Entry(
						nil, itbasisMiddlewareExec.IncludePrevArgsAfter, []string{"arg0"}, []string{"arg1"}, []string{path, "arg1", "arg0"},
					),
				)

			},
		)
	},
)
