package log_test

import (
	"log/slog"

	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.DescribeTable(
	"SlogAttrMap", func(attr slog.Attr, wantKey string, wantValues []slog.Attr) {
		defer ginkgo.GinkgoRecover()

		gomega.Expect(attr).To(
			gomega.SatisfyAll(
				gomega.HaveField("Key", gomega.Equal(wantKey)),
				gomega.HaveField(
					"Value", gomega.SatisfyAll(
						gomega.WithTransform(func(v slog.Value) slog.Kind { return v.Kind() }, gomega.Equal(slog.KindGroup)),
						gomega.WithTransform(
							func(v slog.Value) []slog.Attr { return v.Group() },
							gomega.ConsistOf(wantValues),
						),
					),
				),
			),
		)
	},
	ginkgo.Entry(
		nil,
		itbasisCoreLog.SlogAttrMap("key-string", map[string]int{"w": 1, "p": 2}),
		"key-string",
		[]slog.Attr{slog.Int("w", 1), slog.Int("p", 2)},
	),
	ginkgo.Entry(nil, itbasisCoreLog.SlogAttrMap("key-int", map[int]int{0: 1, 1: 2}), "key-int", []slog.Attr{slog.Int("0", 1), slog.Int("1", 2)}),
)
