package log_test

import (
	"github.com/itbasis/tools/middleware/log"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"SlogAttrSlice", func() {
		defer ginkgo.GinkgoRecover()

		gomega.Expect(log.SlogAttrSlice("key-int", []int{1, -1, 2}).String()).To(gomega.Equal("key-int=1-12"))
		gomega.Expect(log.SlogAttrSlice("key-int64", []int64{1, -1, 2}).String()).To(gomega.Equal("key-int64=1-12"))
		gomega.Expect(log.SlogAttrSlice("key-string", []string{"1", "-1-1", "2 "}).String()).To(gomega.Equal("key-string=1-1-12 "))
	},
)

var _ = ginkgo.Describe(
	"SlogAttrSliceWithSeparator", func() {
		defer ginkgo.GinkgoRecover()

		gomega.Expect(log.SlogAttrSliceWithSeparator("key-int", ",", []int{1, -1, 2}).String()).To(gomega.Equal("key-int=1,-1,2"))
		gomega.Expect(log.SlogAttrSliceWithSeparator("key-int64", ",", []int64{1, -1, 2}).String()).To(gomega.Equal("key-int64=1,-1,2"))
		gomega.Expect(
			log.SlogAttrSliceWithSeparator(
				"key-string",
				",",
				[]string{"1", "-1-1", "2 "},
			).String(),
		).To(gomega.Equal("key-string=1,-1-1,2 "))
	},
)
