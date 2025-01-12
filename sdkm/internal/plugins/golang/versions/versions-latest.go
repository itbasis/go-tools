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

	var v = receiver.cache.Load(ctx, sdkmSDKVersion.TypeStable)

	if len(v) == 0 {
		slog.Debug("Trying to force a cache refresh to find the latest stable version")

		receiver.updateCache(ctx, true, false, false)
		v = receiver.cache.Load(ctx, sdkmSDKVersion.TypeStable)
	}

	if len(v) == 0 {
		return sdkmSDKVersion.SDKVersion{}, errors.Wrap(sdkmPlugin.ErrSDKVersionNotFound, "latest")
	}

	return v[0], nil
}
