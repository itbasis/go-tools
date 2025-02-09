package github

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"strings"

	"github.com/google/go-github/v68/github"
	"github.com/itbasis/go-tools/builder/internal/installer/model"
	"github.com/pkg/errors"
)

const ProviderGithubKey model.ProviderKey = "github"

type Installer interface {
	Install(list model.GithubDependencyList) error
}

var InstallerEmpty = (Installer)(nil)

type installer struct {
	githubClient *github.Client
}

func NewGithubInstaller() (Installer, error) {
	return &installer{
		githubClient: github.NewClient(nil),
	}, nil
}

func (r *installer) Install(list model.GithubDependencyList) error {
	for name, dependency := range list {
		if err := r.installDependency(name, dependency); err != nil {
			return err
		}
	}

	return nil
}

func (r *installer) installDependency(name model.DependencyName, dependency model.GithubDependency) error {
	slog.Info(fmt.Sprintf("install dependency: %s[%s]", name, dependency.Version))

	var (
		githubRelease    *github.RepositoryRelease
		errGetRepository error
	)

	switch dependency.Version {
	case model.VersionLatest:
		githubRelease, _, errGetRepository = r.githubClient.Repositories.GetLatestRelease(context.Background(), dependency.Owner, dependency.Repo)

	default:
		githubRelease, _, errGetRepository = r.githubClient.Repositories.GetLatestRelease(context.Background(), dependency.Owner, dependency.Repo)
	}

	if errGetRepository != nil {
		return errGetRepository //nolint:wrapcheck // TODO
	}

	slog.Info("found repository release: " + githubRelease.GetName())

	var foundAssets []*github.ReleaseAsset

	for _, githubAsset := range githubRelease.Assets {
		var downloadURL = githubAsset.GetBrowserDownloadURL()

		if strings.Contains(downloadURL, runtime.GOOS) && strings.Contains(downloadURL, runtime.GOARCH) {
			foundAssets = append(foundAssets, githubAsset)
		}
	}

	if len(foundAssets) > 1 {
		return errors.New("found more than one github asset")
	} else if len(foundAssets) == 0 {
		return errors.New("no github asset found")
	}

	for _, githubAsset := range foundAssets {
		slog.Info("Github asset [name]: " + githubAsset.GetName())
		slog.Info("Github asset [url]: " + githubAsset.GetBrowserDownloadURL())
	}

	return nil
}
