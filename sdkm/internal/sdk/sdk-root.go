package sdk

import (
	"path"

	itbasisCoreOs "github.com/itbasis/tools/core/os"
)

func GetDefaultSdkRoot() string {
	return path.Join(itbasisCoreOs.UserHomeDir(), "sdk")
}
