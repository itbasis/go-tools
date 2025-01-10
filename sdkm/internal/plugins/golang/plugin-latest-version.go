package golang

import (
	"context"
	"log/slog"
	"strings"

	sdkmPlugin "github.com/itbasis/tools/sdkm/pkg/plugin"
	sdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
	"github.com/pkg/errors"
)

func (receiver *goPlugin) LatestVersion(ctx context.Context) (sdkmSDKVersion.SDKVersion, error) {
	return receiver.sdkVersions.LatestVersion(ctx) //nolint:wrapcheck // TODO
}

func (receiver *goPlugin) LatestVersionByPrefix(ctx context.Context, prefix string) (sdkmSDKVersion.SDKVersion, error) {
	slog.Debug("searching for latest version by prefix: " + prefix)

	if prefix == "" {
		return receiver.LatestVersion(ctx)
	}

	sdkVersions, err := receiver.ListAllVersions(ctx)
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
