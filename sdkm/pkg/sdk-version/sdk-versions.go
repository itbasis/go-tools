package sdkversion

import (
	"context"
	"fmt"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=sdk-versions.mock.go
type SDKVersions interface {
	fmt.GoStringer

	WithCache(cache Cache) SDKVersions

	AllVersions(ctx context.Context, rebuildCache bool) ([]SDKVersion, error)
	LatestVersion(ctx context.Context, rebuildCache bool) (SDKVersion, error)
}
