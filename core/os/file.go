package os

import (
	"log/slog"
	"os"

	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
)

func BeARegularFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		slog.Debug("fail get file info", itbasisCoreLog.SlogAttrError(err))

		return false
	}

	return fileInfo.Mode().IsRegular()
}
