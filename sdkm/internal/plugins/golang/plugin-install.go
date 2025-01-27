package golang

import (
	"context"
	"fmt"
	"log/slog"

	pluginGoConsts "github.com/itbasis/go-tools/sdkm/internal/plugins/golang/consts"
	"github.com/itbasis/go-tools/sdkm/pkg/plugin"
	"github.com/pkg/errors"
)

func (receiver *goPlugin) Install(ctx context.Context, rebuildCache bool, baseDir string) error {
	sdkVersion, errCurrent := receiver.Current(ctx, rebuildCache, baseDir)
	if errCurrent != nil {
		return errors.Wrapf(plugin.ErrSDKInstall, "failed get current version: %s", errCurrent.Error())
	}

	if sdkVersion.Installed {
		slog.Info(fmt.Sprintf("SDK already installed: %s", sdkVersion.ID))

		return nil
	}

	return receiver.InstallVersion(ctx, sdkVersion.ID)
}

func (receiver *goPlugin) InstallVersion(_ context.Context, version string) error {
	if receiver.basePlugin.HasInstalled(pluginGoConsts.PluginID, version) {
		slog.Info("SDK already installed: " + version)

		return nil
	}

	archiveFilePath, errDownload := receiver.downloader.Download(version)
	if errDownload != nil {
		return errors.Wrap(plugin.ErrSDKInstall, errDownload.Error())
	}

	slog.Debug(fmt.Sprintf("Downloading SDK version: %s", version), slog.String("archiveFilePath", archiveFilePath))

	if errUnpack := receiver.downloader.Unpack(
		archiveFilePath, receiver.basePlugin.GetSDKVersionDir(pluginGoConsts.PluginID, version),
	); errUnpack != nil {
		return errors.Wrap(plugin.ErrSDKInstall, errUnpack.Error())
	}

	return nil
}
