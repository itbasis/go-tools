package env_test

import (
	itbasisCoreEnv "github.com/itbasis/go-tools/core/env"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.DescribeTable(
	"MapToSlices", func(value itbasisCoreEnv.Map, want itbasisCoreEnv.List) {
		gomega.Expect(itbasisCoreEnv.MapToSlices(value)).To(gomega.Equal(want))
	},
	ginkgo.Entry(nil, itbasisCoreEnv.Map{}, itbasisCoreEnv.List{}),
	ginkgo.Entry(nil, (itbasisCoreEnv.Map)(nil), itbasisCoreEnv.List{}),
	ginkgo.Entry(nil, itbasisCoreEnv.Map{"1": "2", "test": "va"}, itbasisCoreEnv.List{"1=2", "test=va"}),
)

var _ = ginkgo.DescribeTable(
	"SlicesToMap", func(value itbasisCoreEnv.List, want itbasisCoreEnv.Map) {
		gomega.Expect(itbasisCoreEnv.SlicesToMap(value)).To(gomega.Equal(want))
	},
	ginkgo.Entry(nil, itbasisCoreEnv.List{}, itbasisCoreEnv.Map{}),
	ginkgo.Entry(nil, (itbasisCoreEnv.List)(nil), itbasisCoreEnv.Map{}),
	ginkgo.Entry(
		nil,
		itbasisCoreEnv.List{"1=2", "test=va", "test2=tt=1", "test3=tt 1"},
		itbasisCoreEnv.Map{"1": "2", "test": "va", "test2": "tt=1", "test3": "tt 1"},
	),
)

var _ = ginkgo.Describe(
	"MergeEnvs", func() {
		ginkgo.DescribeTable(
			"source as List", func(source itbasisCoreEnv.List, additional, want itbasisCoreEnv.Map) {
				gomega.Expect(itbasisCoreEnv.MergeEnvs(source, additional)).To(gomega.Equal(want))
			},
			ginkgo.Entry(nil, (itbasisCoreEnv.List)(nil), (itbasisCoreEnv.Map)(nil), (itbasisCoreEnv.Map)(nil)),
			ginkgo.Entry(nil, itbasisCoreEnv.List{}, (itbasisCoreEnv.Map)(nil), (itbasisCoreEnv.Map)(nil)),
			ginkgo.Entry(nil, (itbasisCoreEnv.List)(nil), itbasisCoreEnv.Map{}, itbasisCoreEnv.Map{}),
			ginkgo.Entry(nil, itbasisCoreEnv.List{}, itbasisCoreEnv.Map{}, itbasisCoreEnv.Map{}),
			ginkgo.Entry(nil, itbasisCoreEnv.List{"t=1"}, itbasisCoreEnv.Map{"t": "2"}, itbasisCoreEnv.Map{"t": "2"}),
			ginkgo.Entry(nil, itbasisCoreEnv.List{"t=1"}, itbasisCoreEnv.Map{"t1": "2"}, itbasisCoreEnv.Map{"t": "1", "t1": "2"}),
		)

		ginkgo.DescribeTable(
			"source as Map", func(source, additional, want itbasisCoreEnv.Map) {
				gomega.Expect(itbasisCoreEnv.MergeEnvs(source, additional)).To(gomega.Equal(want))
			},
			ginkgo.Entry(nil, (itbasisCoreEnv.Map)(nil), (itbasisCoreEnv.Map)(nil), (itbasisCoreEnv.Map)(nil)),
			ginkgo.Entry(nil, itbasisCoreEnv.Map{}, (itbasisCoreEnv.Map)(nil), (itbasisCoreEnv.Map)(nil)),
			ginkgo.Entry(nil, (itbasisCoreEnv.Map)(nil), itbasisCoreEnv.Map{}, itbasisCoreEnv.Map{}),
			ginkgo.Entry(nil, itbasisCoreEnv.Map{}, itbasisCoreEnv.Map{}, itbasisCoreEnv.Map{}),
			ginkgo.Entry(nil, itbasisCoreEnv.Map{"t": "1"}, itbasisCoreEnv.Map{"t": "2"}, itbasisCoreEnv.Map{"t": "2"}),
			ginkgo.Entry(nil, itbasisCoreEnv.Map{"t": "1"}, itbasisCoreEnv.Map{"t1": "2"}, itbasisCoreEnv.Map{"t": "1", "t1": "2"}),
		)
	},
)
