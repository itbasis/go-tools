package golang

import (
	"context"
	"log/slog"
	"os"
	"path"

	itbasisCoreLog "github.com/itbasis/tools/core/log"
	itbasisCoreOs "github.com/itbasis/tools/core/os"
	pluginGoConsts "github.com/itbasis/tools/sdkm/internal/plugins/golang/consts"
)

func (receiver *goPlugin) Env(ctx context.Context, rebuildCache bool, baseDir string) (map[string]string, error) {
	sdkVersion, errCurrent := receiver.Current(ctx, rebuildCache, baseDir)

	if errCurrent != nil {
		return map[string]string{}, errCurrent
	}

	return receiver.EnvByVersion(ctx, sdkVersion.ID)
}

func (receiver *goPlugin) EnvByVersion(_ context.Context, version string) (map[string]string, error) {
	var (
		goCacheDir = receiver.getGoCacheDir(version)
		goBin      = path.Join(goCacheDir, "bin")
		originPath = os.Getenv("SDKM_PATH_ORIGIN")
	)

	slog.Debug("originPath: " + originPath)

	if originPath == "" {
		originPath = os.Getenv("PATH")

		slog.Debug("originPath (using env.PATH): " + originPath)
	}

	var envs = map[string]string{
		"SDKM_PATH_ORIGIN":   originPath,
		"SDKM_GOROOT_ORIGIN": os.Getenv("GOROOT"),
		"SDKM_GOPATH_ORIGIN": os.Getenv("GOPATH"),
		"SDKM_GOBIN_ORIGIN":  os.Getenv("GOBIN"),
		//
		"GOROOT": receiver.basePlugin.GetSDKVersionDir(pluginGoConsts.PluginID, version),
		"GOPATH": goCacheDir,
		"GOBIN":  goBin,
		"PATH": itbasisCoreOs.AddBeforePath(
			originPath,
			path.Join(receiver.basePlugin.GetSDKVersionDir(pluginGoConsts.PluginID, version), "bin"),
			goBin,
			itbasisCoreOs.ExecutableDir(),
		),
	}

	slog.Debug("envs", itbasisCoreLog.SlogAttrMap("envs", envs))

	return envs, nil
}
