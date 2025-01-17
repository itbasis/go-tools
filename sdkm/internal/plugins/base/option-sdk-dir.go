package base

import (
	"log/slog"

	itbasisCoreOption "github.com/itbasis/tools/core/option"
	sdkmLog "github.com/itbasis/tools/sdkm/internal/log"
	sdkmSdk "github.com/itbasis/tools/sdkm/internal/sdk"
)

const _optionSdkDirKey = "option-sdk-dir"

func WithDefaultSdkDir() itbasisCoreOption.Option[basePlugin] {
	return &_optionSdkDir{}
}

func WithCustomSdkDir(sdkDir string) itbasisCoreOption.Option[basePlugin] {
	return &_optionSdkDir{dir: sdkDir}
}

type _optionSdkDir struct {
	dir string
}

func (r *_optionSdkDir) Key() itbasisCoreOption.Key { return _optionSdkDirKey }

func (r *_optionSdkDir) Apply(cmp *basePlugin) error {
	slog.Debug("apply SDK directory option", sdkmLog.SlogAttrRootDir(r.dir))

	if r.dir != "" {
		cmp.sdkDir = r.dir

		return nil
	}

	cmp.sdkDir = sdkmSdk.GetDefaultSdkRoot()

	return nil
}
