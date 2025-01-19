package versions

import (
	"context"
	"log/slog"

	sdkmPlugin "github.com/itbasis/tools/sdkm/pkg/plugin"
	sdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
	"github.com/pkg/errors"
)

func (receiver *versions) LatestVersion(ctx context.Context, rebuildCache bool) (sdkmSDKVersion.SDKVersion, error) {
	if rebuildCache || !receiver.cache.Valid(ctx) {
		receiver.updateCache(ctx, true, false, false)
	}

	var sdkVersions = receiver.cache.Load(ctx, sdkmSDKVersion.TypeStable)

	if len(sdkVersions) == 0 {
		slog.Debug("Trying to force a cache refresh to find the latest stable version")

		receiver.updateCache(ctx, true, false, false)
		sdkVersions = receiver.cache.Load(ctx, sdkmSDKVersion.TypeStable)
	}

	if len(sdkVersions) == 0 {
		return sdkmSDKVersion.SDKVersion{}, errors.Wrap(sdkmPlugin.ErrSDKVersionNotFound, "latest")
	}

	return sdkVersions[0], nil
}
