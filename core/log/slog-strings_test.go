package log_test

import (
	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"SlogAttrSlice", func() {
		defer ginkgo.GinkgoRecover()

		gomega.Expect(itbasisCoreLog.SlogAttrSlice("key-int", []int{1, -1, 2}).String()).To(gomega.Equal("key-int=1-12"))
		gomega.Expect(itbasisCoreLog.SlogAttrSlice("key-int64", []int64{1, -1, 2}).String()).To(gomega.Equal("key-int64=1-12"))
		gomega.Expect(itbasisCoreLog.SlogAttrSlice("key-string", []string{"1", "-1-1", "2 "}).String()).To(gomega.Equal("key-string=1-1-12 "))
	},
)

var _ = ginkgo.Describe(
	"SlogAttrSliceWithSeparator", func() {
		defer ginkgo.GinkgoRecover()

		gomega.Expect(itbasisCoreLog.SlogAttrSliceWithSeparator("key-int", ",", []int{1, -1, 2}).String()).To(gomega.Equal("key-int=1,-1,2"))
		gomega.Expect(itbasisCoreLog.SlogAttrSliceWithSeparator("key-int64", ",", []int64{1, -1, 2}).String()).To(gomega.Equal("key-int64=1,-1,2"))
		gomega.Expect(
			itbasisCoreLog.SlogAttrSliceWithSeparator(
				"key-string",
				",",
				[]string{"1", "-1-1", "2 "},
			).String(),
		).To(gomega.Equal("key-string=1,-1-1,2 "))
	},
)
