package sdk

import (
	"path"

	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
)

func GetDefaultSdkRoot() string {
	return path.Join(itbasisMiddlewareOs.UserHomeDir(), "sdk")
}
