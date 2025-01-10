package golang

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/itbasis/tools/middleware/log"
	"github.com/itbasis/tools/sdkm/internal/plugins/golang/modfile"
	sdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
)

func (receiver *goPlugin) Current(ctx context.Context, baseDir string) (sdkmSDKVersion.SDKVersion, error) {
	goModFile, errGoModFile := modfile.ReadGoModFile(baseDir)
	if errGoModFile != nil {
		slog.Error("Failed to read go.mod file", log.SlogAttrError(errGoModFile))

		return sdkmSDKVersion.SDKVersion{}, errGoModFile //nolint:wrapcheck // TODO
	}

	var (
		sdkVersion sdkmSDKVersion.SDKVersion
		err        error
	)

	if toolchain := goModFile.Toolchain; toolchain != nil {
		sdkVersion, err = receiver.LatestVersionByPrefix(ctx, toolchain.Name[2:])
	} else {
		sdkVersion, err = receiver.LatestVersionByPrefix(ctx, goModFile.Go.Version)
	}

	slog.Debug(fmt.Sprintf("sdkVersion: %++v", sdkVersion))

	if err != nil {
		return sdkmSDKVersion.SDKVersion{}, err
	}

	receiver.enrichSDKVersion(&sdkVersion)

	return sdkVersion, nil
}
