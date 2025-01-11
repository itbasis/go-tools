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

	ListAllVersions(ctx context.Context) ([]sdkmSDKVersion.SDKVersion, error)
	ListAllVersionsByPrefix(ctx context.Context, prefix string) ([]sdkmSDKVersion.SDKVersion, error)

	LatestVersion(ctx context.Context) (sdkmSDKVersion.SDKVersion, error)
	LatestVersionByPrefix(ctx context.Context, prefix string) (sdkmSDKVersion.SDKVersion, error)

	Current(ctx context.Context, baseDir string) (sdkmSDKVersion.SDKVersion, error)

	Install(ctx context.Context, baseDir string) error
	InstallVersion(ctx context.Context, version string) error

	Env(ctx context.Context, baseDir string) (map[string]string, error)
	EnvByVersion(ctx context.Context, version string) (map[string]string, error)

	Exec(ctx context.Context, baseDir string, stdIn io.Reader, stdOut, stdErr io.Writer, args []string) error
}
