package versions

import (
	"context"
	"log/slog"

	sdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
)

func (receiver *versions) AllVersions(ctx context.Context, rebuildCache bool) ([]sdkmSDKVersion.SDKVersion, error) {
	if rebuildCache || !receiver.cache.Valid(ctx) {
		receiver.updateCache(ctx, true, true, true)
	}

	var filterVersions = func(ctx context.Context) []sdkmSDKVersion.SDKVersion {
		var sdkVersions []sdkmSDKVersion.SDKVersion

		for _, versionType := range []sdkmSDKVersion.VersionType{sdkmSDKVersion.TypeStable, sdkmSDKVersion.TypeUnstable, sdkmSDKVersion.TypeArchived} {
			v := receiver.cache.Load(ctx, versionType)
			if len(v) == 0 {
				continue
			}

			sdkVersions = append(sdkVersions, v...)
		}

		return sdkVersions
	}

	if result := filterVersions(ctx); len(result) > 0 {
		return result, nil
	}

	slog.Debug("Trying to force refresh cache to search for versions")

	receiver.updateCache(ctx, true, true, true)

	return filterVersions(ctx), nil
}
