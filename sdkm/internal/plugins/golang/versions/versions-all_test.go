package versions_test

import (
	"context"
	"path"

	pluginGoVersions "github.com/itbasis/tools/sdkm/internal/plugins/golang/versions"
	sdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"All Versions", func() {
		defer ginkgo.GinkgoRecover()

		ginkgo.DescribeTable(
			"different versions of answers", func(testResponseFilePath string, wantLen int, wantSDKVersions []sdkmSDKVersion.SDKVersion) {
				var server = initFakeServer(path.Join("testdata", "all-versions", testResponseFilePath))
				defer server.Close()

				sdkVersions := pluginGoVersions.NewVersions(server.URL())

				gomega.Expect(sdkVersions.AllVersions(context.Background(), false)).
					To(
						gomega.SatisfyAll(
							gomega.HaveLen(wantLen),

							gomega.ContainElements(
								sdkmSDKVersion.SDKVersion{ID: "1.23rc1", Type: sdkmSDKVersion.TypeArchived},
								sdkmSDKVersion.SDKVersion{ID: "1.22.0", Type: sdkmSDKVersion.TypeArchived},
								sdkmSDKVersion.SDKVersion{ID: "1.18", Type: sdkmSDKVersion.TypeArchived},
								sdkmSDKVersion.SDKVersion{ID: "1.18.10", Type: sdkmSDKVersion.TypeArchived},
								sdkmSDKVersion.SDKVersion{ID: "1.4beta1", Type: sdkmSDKVersion.TypeArchived},
								sdkmSDKVersion.SDKVersion{ID: "1.3rc1", Type: sdkmSDKVersion.TypeArchived},
							),

							gomega.ContainElements(wantSDKVersions),
						),
					)
			},
			ginkgo.Entry(
				nil, "001.html", 292, []sdkmSDKVersion.SDKVersion{
					{ID: "1.23rc2", Type: sdkmSDKVersion.TypeUnstable},
					{ID: "1.22.5", Type: sdkmSDKVersion.TypeStable},
				},
			),
			ginkgo.Entry(
				nil, "002.html", 295, []sdkmSDKVersion.SDKVersion{
					{ID: "1.23.0", Type: sdkmSDKVersion.TypeStable},
					{ID: "1.23rc2", Type: sdkmSDKVersion.TypeArchived},
					{ID: "1.22.6", Type: sdkmSDKVersion.TypeStable},
					{ID: "1.22.5", Type: sdkmSDKVersion.TypeArchived},
				},
			),
			ginkgo.Entry(
				nil, "003.html", 304, []sdkmSDKVersion.SDKVersion{
					{ID: "1.24rc1", Type: sdkmSDKVersion.TypeUnstable},
					{ID: "1.23.4", Type: sdkmSDKVersion.TypeStable},
					{ID: "1.23.0", Type: sdkmSDKVersion.TypeArchived},
					{ID: "1.23rc2", Type: sdkmSDKVersion.TypeArchived},
					{ID: "1.22.10", Type: sdkmSDKVersion.TypeStable},
					{ID: "1.22.6", Type: sdkmSDKVersion.TypeArchived},
					{ID: "1.22.5", Type: sdkmSDKVersion.TypeArchived},
				},
			),
		)
	},
)
