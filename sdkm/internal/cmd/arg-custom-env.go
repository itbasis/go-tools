package cmd

import (
	"log/slog"
	"os"

	itbasisMiddlewareCmd "github.com/itbasis/tools/middleware/cmd"
	sdkmLog "github.com/itbasis/tools/sdkm/internal/log"
	sdkmSdk "github.com/itbasis/tools/sdkm/internal/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	_flagSdkRootDir = "sdk-root-dir"
	_envSDKRootDir  = "SDKM_SDK_ROOT_DIR"

	_flagCacheRootDir = "cache-root-dir"
	_envCacheRootDir  = "SDKM_CACHE_ROOT_DIR"
)

func InitFlagSdkRootDir(flags *pflag.FlagSet) {
	flags.String(_flagSdkRootDir, "", "SDK root directory")
}

func InitFlagCacheRootDir(flags *pflag.FlagSet) {
	flags.String(_flagCacheRootDir, "", "Cache root directory")
}

func GetSdkRootDir(cmd *cobra.Command) string {
	return _getRootDir(cmd, _flagSdkRootDir, _envSDKRootDir)
}

func GetCacheRootDir(cmd *cobra.Command) string {
	return _getRootDir(cmd, _flagCacheRootDir, _envCacheRootDir)
}

func _getRootDir(cmd *cobra.Command, flag, envName string) string {
	slog.Debug("getting root directory", slog.String("flag", flag), slog.String("envName", envName))

	var (
		rootDir string
		err     error
	)

	rootDir, err = cmd.Flags().GetString(flag)
	itbasisMiddlewareCmd.RequireNoError(cmd, err)

	if rootDir != "" {
		slog.Debug("using from a command line argument", sdkmLog.SlogAttrRootDir(rootDir))

		return rootDir
	}

	if rootDir = os.Getenv(envName); rootDir != "" {
		slog.Debug("using from environment variables", sdkmLog.SlogAttrRootDir(rootDir))

		return rootDir
	}

	rootDir = sdkmSdk.GetDefaultSdkRoot()
	slog.Debug("using with default value", sdkmLog.SlogAttrRootDir(rootDir))

	return rootDir
}
