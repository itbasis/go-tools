package golang

import (
	"context"
	"log/slog"
	"strings"

	sdkmPlugin "github.com/itbasis/go-tools/sdkm/pkg/plugin"
	sdkmSDKVersion "github.com/itbasis/go-tools/sdkm/pkg/sdk-version"
	"github.com/pkg/errors"
)

func (receiver *goPlugin) LatestVersion(ctx context.Context, rebuildCache bool) (sdkmSDKVersion.SDKVersion, error) {
	sdkVersion, err := receiver.sdkVersions.LatestVersion(ctx, rebuildCache)
	if err != nil {
		return sdkmSDKVersion.SDKVersion{}, err //nolint:wrapcheck // TODO
	}

	receiver.enrichSDKVersion(&sdkVersion)

	return sdkVersion, nil
}

func (receiver *goPlugin) LatestVersionByPrefix(ctx context.Context, rebuildCache bool, prefix string) (sdkmSDKVersion.SDKVersion, error) {
	slog.Debug("searching for latest version by prefix: " + prefix)

	if prefix == "" {
		return receiver.LatestVersion(ctx, rebuildCache)
	}

	sdkVersions, err := receiver.ListAllVersions(ctx, rebuildCache)
	if err != nil {
		return sdkmSDKVersion.SDKVersion{}, errors.Wrap(sdkmPlugin.ErrSDKVersionNotFound, err.Error())
	}

	for _, sdkVersion := range sdkVersions {
		if strings.HasPrefix(sdkVersion.ID, prefix) {
			receiver.enrichSDKVersion(&sdkVersion)

			return sdkVersion, nil
		}
	}

	return sdkmSDKVersion.SDKVersion{}, errors.Wrap(sdkmPlugin.ErrSDKVersionNotFound, "version by prefix "+prefix)
}
