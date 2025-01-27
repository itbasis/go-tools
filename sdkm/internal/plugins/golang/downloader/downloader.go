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

	itbasisCoreOs "github.com/itbasis/go-tools/core/os"
	sdkmHttp "github.com/itbasis/go-tools/sdkm/internal/http"
	sdkmPlugin "github.com/itbasis/go-tools/sdkm/pkg/plugin"
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

	if err := os.MkdirAll(filepath.Dir(outFilePath), itbasisCoreOs.DefaultDirMode); err != nil {
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

	var tmpDirPath = path.Clean(filepath.FromSlash(targetDir + ".tmp"))
	if err := os.MkdirAll(tmpDirPath, itbasisCoreOs.DefaultDirMode); err != nil {
		return errors.Wrapf(sdkmPlugin.ErrDownloadFailed, "fail create temporary dir: %s", err)
	}

	slog.Debug("creating tmp dir for unpack: " + tmpDirPath)

	defer func(path string) {
		if err := os.RemoveAll(path); err != nil {
			panic(err)
		}
	}(tmpDirPath)

	var errExtract error

	if filepath.Ext(archiveFilePath) == ".zip" {
		slog.Debug(fmt.Sprintf("unpacking zip archive '%s' to '%s'", archiveFilePath, tmpDirPath))

		_, _, errExtract = xtractr.ExtractZIP(&xtractr.XFile{FilePath: archiveFilePath, OutputDir: tmpDirPath})
	} else {
		slog.Debug(fmt.Sprintf("unpacking tar.gz archive '%s' to '%s'", archiveFilePath, tmpDirPath))

		_, _, errExtract = xtractr.ExtractTarGzip(
			&xtractr.XFile{FilePath: archiveFilePath, OutputDir: tmpDirPath, DirMode: xtractr.DefaultDirMode, FileMode: xtractr.DefaultFileMode},
		)
	}

	if errExtract != nil {
		return errors.Wrapf(sdkmPlugin.ErrDownloadFailed, "extracting %s failed", archiveFilePath)
	}

	// issue https://github.com/golift/xtractr/issues/70
	if errRename := os.Rename(path.Join(tmpDirPath, "go"), path.Clean(filepath.FromSlash(targetDir))); errRename != nil {
		return errors.Wrapf(sdkmPlugin.ErrDownloadFailed, "failed rename: %s", errRename.Error())
	}

	return nil
}

func (receiver *Downloader) URLForDownload(version string) string {
	switch receiver.os {
	case "windows":
		return fmt.Sprintf("%s/go%s.%s-%s.zip", receiver.urlReleases, version, receiver.os, receiver.arch)

	default:
		return fmt.Sprintf("%s/go%s.%s-%s.tar.gz", receiver.urlReleases, version, receiver.os, receiver.arch)
	}
}

func (receiver *Downloader) GoString() string {
	return "downloader{os=" + receiver.os + "; arch=" + receiver.arch + "; urlReleases: " + receiver.urlReleases + "}"
}
