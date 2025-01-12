package plugin

import (
	"context"
	"io"

	sdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
)

//nolint:interfacebloat // TODO
//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=plugin.mock.go
type SDKMPlugin interface {
	WithVersions(versions sdkmSDKVersion.SDKVersions) SDKMPlugin
	// WithBasePlugin(basePlugin BasePlugin) SDKMPlugin
	// WithDownloader(downloader) SDKMPlugin

	ListAllVersions(ctx context.Context, rebuildCache bool) ([]sdkmSDKVersion.SDKVersion, error)
	ListAllVersionsByPrefix(ctx context.Context, rebuildCache bool, prefix string) ([]sdkmSDKVersion.SDKVersion, error)

	LatestVersion(ctx context.Context, rebuildCache bool) (sdkmSDKVersion.SDKVersion, error)
	LatestVersionByPrefix(ctx context.Context, rebuildCache bool, prefix string) (sdkmSDKVersion.SDKVersion, error)

	Current(ctx context.Context, rebuildCache bool, baseDir string) (sdkmSDKVersion.SDKVersion, error)

	Install(ctx context.Context, rebuildCache bool, baseDir string) error
	InstallVersion(ctx context.Context, version string) error

	Env(ctx context.Context, rebuildCache bool, baseDir string) (map[string]string, error)
	EnvByVersion(ctx context.Context, version string) (map[string]string, error)

	Exec(ctx context.Context, rebuildCache bool, baseDir string, stdIn io.Reader, stdOut, stdErr io.Writer, args []string) error
}
