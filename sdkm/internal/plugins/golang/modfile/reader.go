package modfile

import (
	"os"
	"path"
	"path/filepath"

	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	"github.com/pkg/errors"
	"golang.org/x/mod/modfile"
)

const ErrMsgFailedReadGoModFile = "failed to read go.mod file"

func ReadGoModFile(baseDir string) (modfile.File, error) {
	var (
		goModFilePath             = path.Join(baseDir, "go.mod")
		goModFileRelPath, errPath = filepath.Rel(itbasisMiddlewareOs.Pwd(), goModFilePath)
	)

	if errPath != nil {
		return modfile.File{}, errors.Wrap(errPath, ErrMsgFailedReadGoModFile)
	}

	bytes, errRead := os.ReadFile(goModFilePath)
	if errRead != nil {
		return modfile.File{}, errors.Wrap(errRead, ErrMsgFailedReadGoModFile)
	}

	file, errParse := modfile.Parse(goModFileRelPath, bytes, nil)
	if errParse != nil {
		return modfile.File{}, errors.Wrap(errParse, ErrMsgFailedReadGoModFile)
	}

	return *file, nil
}
