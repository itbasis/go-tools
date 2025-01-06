package golang

import (
	"context"
	"strings"

	sdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
)

func (receiver *goPlugin) ListAllVersions(ctx context.Context) ([]sdkmSDKVersion.SDKVersion, error) {
	return receiver.sdkVersions.AllVersions(ctx) //nolint:wrapcheck // _
}

func (receiver *goPlugin) ListAllVersionsByPrefix(ctx context.Context, prefix string) ([]sdkmSDKVersion.SDKVersion, error) {
	var allVersions, err = receiver.sdkVersions.AllVersions(ctx)

	if err != nil {
		return nil, err //nolint:wrapcheck // _
	}

	if prefix == "" {
		return allVersions, nil
	}

	var versions []sdkmSDKVersion.SDKVersion

	for _, sdkVersion := range allVersions {
		if strings.HasPrefix(sdkVersion.ID, prefix) {
			receiver.enrichSDKVersion(&sdkVersion)

			versions = append(versions, sdkVersion)
		}
	}

	return versions, nil
}
