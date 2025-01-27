package golang

import (
	"context"
	"fmt"
	"log/slog"

	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
	"github.com/itbasis/go-tools/sdkm/internal/plugins/golang/modfile"
	sdkmSDKVersion "github.com/itbasis/go-tools/sdkm/pkg/sdk-version"
)

func (receiver *goPlugin) Current(ctx context.Context, rebuildCache bool, baseDir string) (sdkmSDKVersion.SDKVersion, error) {
	goModFile, errGoModFile := modfile.ReadGoModFile(baseDir)
	if errGoModFile != nil {
		slog.Error("Failed to read go.mod file", itbasisCoreLog.SlogAttrError(errGoModFile))

		return sdkmSDKVersion.SDKVersion{}, errGoModFile //nolint:wrapcheck // TODO
	}

	var (
		sdkVersion sdkmSDKVersion.SDKVersion
		err        error
	)

	if toolchain := goModFile.Toolchain; toolchain != nil {
		sdkVersion, err = receiver.LatestVersionByPrefix(ctx, rebuildCache, toolchain.Name[2:])
	} else {
		sdkVersion, err = receiver.LatestVersionByPrefix(ctx, rebuildCache, goModFile.Go.Version)
	}

	slog.Debug(fmt.Sprintf("sdkVersion: %++v", sdkVersion))

	if err != nil {
		return sdkmSDKVersion.SDKVersion{}, err
	}

	receiver.enrichSDKVersion(&sdkVersion)

	return sdkVersion, nil
}
