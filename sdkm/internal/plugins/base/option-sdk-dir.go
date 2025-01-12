package base

import (
	"log/slog"

	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
	sdkmLog "github.com/itbasis/tools/sdkm/internal/log"
	sdkmSdk "github.com/itbasis/tools/sdkm/internal/sdk"
)

const _optionSdkDirKey = "option-sdk-dir"

func WithDefaultSdkDir() itbasisMiddlewareOption.Option[basePlugin] {
	return &_optionSdkDir{}
}

func WithCustomSdkDir(sdkDir string) itbasisMiddlewareOption.Option[basePlugin] {
	return &_optionSdkDir{dir: sdkDir}
}

type _optionSdkDir struct {
	dir string
}

func (r *_optionSdkDir) Key() itbasisMiddlewareOption.Key { return _optionSdkDirKey }

func (r *_optionSdkDir) Apply(cmp *basePlugin) error {
	slog.Debug("apply SDK directory option", sdkmLog.SlogAttrRootDir(r.dir))

	if r.dir != "" {
		cmp.sdkDir = r.dir

		return nil
	}

	cmp.sdkDir = sdkmSdk.GetDefaultSdkRoot()

	return nil
}
