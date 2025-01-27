//go:build windows

package os_test

import (
	itbasisCoreOs "github.com/itbasis/go-tools/core/os"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.DescribeTable(
	"CleanPath", func(path string, cleanPaths []string, wantPath string) {
		gomega.Expect(itbasisCoreOs.CleanPath(path, cleanPaths...)).
			To(gomega.Equal(wantPath))
	},
	ginkgo.Entry(nil, "a;b;c", []string{}, "a;b;c"),
	ginkgo.Entry(nil, "a;b;c", []string{"a"}, "b;c"),
	ginkgo.Entry(nil, "a;b;a;c", []string{"a"}, "b;c"),
	ginkgo.Entry(nil, "a;b;c", []string{"b"}, "a;c"),
	ginkgo.Entry(nil, "a;b;c", []string{"c"}, "a;b"),
	ginkgo.Entry(nil, "a;b;c", []string{"a", "c"}, "b"),
)

var _ = ginkgo.DescribeTable(
	"AddBeforePath", func(path string, addPaths []string, wantPath string) {
		gomega.Expect(itbasisCoreOs.AddBeforePath(path, addPaths...)).
			To(gomega.Equal(wantPath))
	},
	ginkgo.Entry(nil, "a;b;c", []string{}, "a;b;c"),
	ginkgo.Entry(nil, "a;b;c", []string{"d"}, "d;a;b;c"),
	ginkgo.Entry(nil, "a;b;c", []string{"d", "e"}, "d;e;a;b;c"),
)
