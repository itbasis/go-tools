package log

import "log/slog"

func SlogAttrSdkRootDir(sdkRootDir string) slog.Attr {
	return slog.String("sdk_root_dir", sdkRootDir)
}
