package filestorage

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/itbasis/go-clock/v2"
	itbasisCoreLog "github.com/itbasis/tools/core/log"
	itbasisCoreOs "github.com/itbasis/tools/core/os"
	sdkmPlugin "github.com/itbasis/tools/sdkm/pkg/plugin"
	itbasisSdkmSDKVersion "github.com/itbasis/tools/sdkm/pkg/sdk-version"
)

const (
	cacheExpirationDuration = 24 * time.Hour
)

var (
	emptyLoadResult = map[itbasisSdkmSDKVersion.VersionType][]itbasisSdkmSDKVersion.SDKVersion{}
)

type fileStorage struct {
	lock sync.Mutex

	filePath string
}

func NewFileCacheStorage(pluginID sdkmPlugin.ID) itbasisSdkmSDKVersion.CacheStorage {
	return NewFileCacheStorageCustomPath(path.Join(itbasisCoreOs.ExecutableDir(), ".cache", string(pluginID)+".json"))
}

func NewFileCacheStorageCustomPath(filePath string) itbasisSdkmSDKVersion.CacheStorage {
	slog.Debug("using cache with file path: " + filePath)

	return &fileStorage{filePath: filePath}
}

func (receiver *fileStorage) GoString() string {
	return "FileCacheStorage[file=" + receiver.filePath + "]"
}

func (receiver *fileStorage) Valid(ctx context.Context) bool {
	filePath := receiver.filePath

	slog.Debug("validating with file path: " + filePath)

	if filePath == "" {
		slog.Debug("file path is empty: " + filePath)

		return false
	}

	fileInfo, errStat := os.Stat(filePath)
	if errStat != nil && os.IsNotExist(errStat) {
		slog.Debug("cache file not found: " + filePath)

		return false
	} else if errStat != nil {
		slog.Error("AttrError accessing cache file", itbasisCoreLog.SlogAttrError(errStat))

		return false
	}

	if clock.FromContext(ctx).Now().Sub(fileInfo.ModTime()) >= cacheExpirationDuration {
		slog.Debug("cache file has been expired: " + filePath)

		return false
	}

	return true
}

func (receiver *fileStorage) Load(ctx context.Context) map[itbasisSdkmSDKVersion.VersionType][]itbasisSdkmSDKVersion.SDKVersion {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	var filePath = receiver.filePath

	slog.Debug("loading cache from file: " + filePath)

	if !receiver.Valid(ctx) {
		return emptyLoadResult
	}

	var bytes, errReadFile = os.ReadFile(filePath)
	if errReadFile != nil {
		slog.Error("error reading cache file: "+filePath, itbasisCoreLog.SlogAttrError(errReadFile))

		return emptyLoadResult
	}

	var model model

	if errUnmarshal := json.Unmarshal(bytes, &model); errUnmarshal != nil {
		slog.Error(
			"error unmarshalling cache file",
			itbasisCoreLog.SlogAttrError(errUnmarshal),
			itbasisCoreLog.SlogAttrFilePath(filePath),
		)

		return emptyLoadResult
	}

	slog.Debug("loaded cache from file: " + filePath)

	return model.Versions
}

func (receiver *fileStorage) Store(ctx context.Context, versions map[itbasisSdkmSDKVersion.VersionType][]itbasisSdkmSDKVersion.SDKVersion) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	filePath := receiver.filePath

	slog.Debug("storing cache to file: " + filePath)

	var bytes, errMarshal = json.Marshal(
		model{
			Updated:  updated(clock.FromContext(ctx).Now()),
			Versions: versions,
		},
	)

	if errMarshal != nil {
		slog.Error(
			"error marshalling cache file",
			itbasisCoreLog.SlogAttrError(errMarshal),
			itbasisCoreLog.SlogAttrFilePath(filePath),
		)

		return
	}

	dir := filepath.Dir(filePath)
	if errMkdir := os.MkdirAll(dir, itbasisCoreOs.DefaultDirMode); errMkdir != nil {
		slog.Error("error creating cache dir: "+dir, itbasisCoreLog.SlogAttrError(errMkdir))

		return
	}

	if errWriteFile := os.WriteFile(filePath, bytes, itbasisCoreOs.DefaultFileMode); errWriteFile != nil {
		slog.Error("error writing cache file: "+filePath, itbasisCoreLog.SlogAttrError(errWriteFile))
	}
}
