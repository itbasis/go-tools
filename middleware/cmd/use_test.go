package cmd_test

import (
	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.DescribeTable(
	"build use", func(args []string, want string) {
		gomega.Expect(itbasisMiddlewareCmd.BuildUse(args...)).To(gomega.Equal(want))
	},
	ginkgo.Entry(nil, []string{"test"}, "test"),
	ginkgo.Entry(nil, []string{"test", "test1"}, "test test1"),
	ginkgo.Entry(nil, []string{"test ", "test1"}, "test test1"),
	ginkgo.Entry(nil, []string{" test", "test1"}, "test test1"),
	ginkgo.Entry(nil, []string{"test ", " test1"}, "test test1"),
)
