package os

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"

	itbasisMiddlewareLog "github.com/itbasis/tools/middleware/log"
)

const (
	DefaultDirMode  = 0o755
	DefaultFileMode = 0o644
)

func Pwd() string {
	executable, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	return executable
}

func UserHomeDir() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	return userHomeDir
}

func ExecutableDir() string {
	executable, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}

	return filepath.Dir(executable)
}

func BeARegularFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		slog.Debug("fail get file info", itbasisMiddlewareLog.SlogAttrError(err))

		return false
	}

	return fileInfo.Mode().IsRegular()
}
