package versions

import (
	"context"

	sdkmPlugin "github.com/itbasis/tools/sdkm/pkg/plugin"
	sdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
	"github.com/pkg/errors"
)

func (receiver *versions) LatestVersion(ctx context.Context) (sdkmSDKVersion.SDKVersion, error) {
	if !receiver.cache.Valid(ctx) {
		receiver.parseVersions(ctx, sdkmSDKVersion.TypeStable, receiver.reStableGroupVersions, true)
	}

	v := receiver.cache.Load(ctx, sdkmSDKVersion.TypeStable)
	if len(v) == 0 {
		return sdkmSDKVersion.SDKVersion{}, errors.Wrap(sdkmPlugin.ErrSDKVersionNotFound, "latest")
	}

	return v[0], nil
}
