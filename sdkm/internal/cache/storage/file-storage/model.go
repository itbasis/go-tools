package filestorage

import (
	itbasisSdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
)

type model struct {
	Updated  updated
	Versions map[itbasisSdkmSDKVersion.VersionType][]itbasisSdkmSDKVersion.SDKVersion
}
