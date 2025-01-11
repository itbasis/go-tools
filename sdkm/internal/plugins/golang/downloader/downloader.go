package downloader

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	sdkmHttp "github.com/itbasis/tools/sdkm/internal/http"
	pluginGoConsts "github.com/itbasis/tools/sdkm/internal/plugins/golang/consts"
	sdkmPlugin "github.com/itbasis/tools/sdkm/pkg/plugin"
	"github.com/pkg/errors"
	"golift.io/xtractr"
)

type Downloader struct {
	fmt.GoStringer

	httpClient *http.Client

	urlReleases string

	os   string
	arch string

	basePlugin sdkmPlugin.BasePlugin
}

func NewDownloader(os, arch, urlReleases string, basePlugin sdkmPlugin.BasePlugin) *Downloader {
	return &Downloader{
		os:          os,
		arch:        arch,
		urlReleases: urlReleases,
		basePlugin:  basePlugin,
		httpClient:  sdkmHttp.NewHTTPClient(time.Minute),
	}
}

func (receiver *Downloader) Download(version string) (string, error) {
	url := receiver.URLForDownload(version)
	outFilePath := filepath.Join(receiver.basePlugin.GetSDKDir(), ".download", filepath.Base(url))

	if err := os.MkdirAll(filepath.Dir(outFilePath), itbasisMiddlewareOs.DefaultDirMode); err != nil {
		return "", errors.Wrapf(sdkmPlugin.ErrDownloadFailed, "fail make directories: %s", err.Error())
	}

	outFile, errOutFile := os.Create(outFilePath)
	if errOutFile != nil {
		return "", errors.Wrapf(sdkmPlugin.ErrDownloadFailed, "fail create output file: %s", errOutFile.Error())
	}

	defer func(outFile *os.File) {
		if err := outFile.Close(); err != nil {
			panic(err)
		}
	}(outFile)

	slog.Info(fmt.Sprintf("downloading '%s' to '%s'", url, outFilePath))

	//nolint:noctx // TODO
	resp, errResp := receiver.httpClient.Get(url)
	if errResp != nil {
		return "", errors.Wrap(sdkmPlugin.ErrDownloadFailed, errResp.Error())
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Wrapf(sdkmPlugin.ErrDownloadFailed, "status code %d", resp.StatusCode)
	}

	if _, err := io.Copy(outFile, resp.Body); err != nil {
		return "", errors.Wrapf(sdkmPlugin.ErrDownloadFailed, "failed copy file: %s", err.Error())
	}

	slog.Info(fmt.Sprintf("downloaded '%s' to '%s'", url, outFilePath))

	return outFilePath, nil
}

func (receiver *Downloader) Unpack(archiveFilePath, targetDir string) error {
	slog.Debug(fmt.Sprintf("unpacking '%s' to '%s'", archiveFilePath, targetDir))

	tmpDirPath, errTmpDirPath := os.MkdirTemp("", "sdkm-"+string(pluginGoConsts.PluginID))
	if errTmpDirPath != nil {
		return errors.Wrapf(sdkmPlugin.ErrDownloadFailed, "fail create temporary dir: %s", errTmpDirPath)
	}

	defer func(path string) {
		if err := os.RemoveAll(path); err != nil {
			panic(err)
		}
	}(tmpDirPath)

	if _, _, err := xtractr.ExtractTarGzip(
		&xtractr.XFile{FilePath: archiveFilePath, OutputDir: tmpDirPath, DirMode: xtractr.DefaultDirMode, FileMode: xtractr.DefaultFileMode},
	); err != nil {
		return errors.Wrapf(sdkmPlugin.ErrDownloadFailed, "extracting %s failed", archiveFilePath)
	}

	// issue https://github.com/golift/xtractr/issues/70
	if errRename := os.Rename(path.Join(tmpDirPath, "go"), targetDir); errRename != nil {
		return errors.Wrapf(sdkmPlugin.ErrDownloadFailed, "failed rename: %s", errRename.Error())
	}

	return nil
}

func (receiver *Downloader) URLForDownload(version string) string {
	return fmt.Sprintf("%s/go%s.%s-%s.tar.gz", receiver.urlReleases, version, receiver.os, receiver.arch)
}

func (receiver *Downloader) GoString() string {
	return "downloader{os=" + receiver.os + "; arch=" + receiver.arch + "; urlReleases: " + receiver.urlReleases + "}"
}
