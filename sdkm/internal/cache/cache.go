package cache

import (
	"context"
	"sync"

	itbasisSdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
	"golang.org/x/exp/maps"
)

type cache struct {
	storeLock sync.Mutex

	cacheStorage itbasisSdkmSDKVersion.CacheStorage
	cache        map[itbasisSdkmSDKVersion.VersionType][]itbasisSdkmSDKVersion.SDKVersion
}

func NewCache() itbasisSdkmSDKVersion.Cache {
	return &cache{
		cache: map[itbasisSdkmSDKVersion.VersionType][]itbasisSdkmSDKVersion.SDKVersion{},
	}
}

func (receiver *cache) WithExternalStore(cacheStorage itbasisSdkmSDKVersion.CacheStorage) itbasisSdkmSDKVersion.Cache {
	receiver.cacheStorage = cacheStorage
	maps.Clear(receiver.cache)

	return receiver
}

func (receiver *cache) Valid(ctx context.Context) bool {
	if receiver.cacheStorage != nil {
		return receiver.cacheStorage.Valid(ctx)
	}

	return len(receiver.cache) > 0
}

func (receiver *cache) Load(ctx context.Context, versionType itbasisSdkmSDKVersion.VersionType) []itbasisSdkmSDKVersion.SDKVersion {
	receiver.storeLock.Lock()
	defer receiver.storeLock.Unlock()

	if cacheStorage := receiver.cacheStorage; cacheStorage != nil {
		if len(receiver.cache) == 0 || !cacheStorage.Valid(ctx) {
			receiver.cache = cacheStorage.Load(ctx)
		}
	}

	var list, ok = receiver.cache[versionType]
	if !ok {
		return []itbasisSdkmSDKVersion.SDKVersion{}
	}

	return list
}

func (receiver *cache) Store(
	ctx context.Context, versionType itbasisSdkmSDKVersion.VersionType, sdkVersions []itbasisSdkmSDKVersion.SDKVersion,
) {
	receiver.storeLock.Lock()
	defer receiver.storeLock.Unlock()

	receiver.cache[versionType] = sdkVersions

	if receiver.cacheStorage != nil {
		receiver.cacheStorage.Store(ctx, receiver.cache)
	}
}

func (receiver *cache) GoString() string {
	var result = "SDKVersionCache"

	if receiver.cacheStorage != nil {
		result = result + " (" + receiver.cacheStorage.GoString() + ")"
	}

	return result
}
