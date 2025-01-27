package installer

import (
	"context"
	"log/slog"
	"runtime"
	"strings"

	"github.com/google/go-github/v68/github"
	"github.com/pkg/errors"
)

func (r *Installer) installGitHub() error {
	githubClient := github.NewClient(nil)

	for dependencyName, dependency := range r.dependencies.GithubDependencies {
		slog.Info("install Github dependency: " + dependencyName + " [" + dependency.Version + "]")

		var (
			githubRelease    *github.RepositoryRelease
			errGetRepository error
		)
		switch dependency.Version {
		case _versionLatest:
			githubRelease, _, errGetRepository = githubClient.Repositories.GetLatestRelease(context.Background(), dependency.Owner, dependency.Repo)
		default:
			githubRelease, _, errGetRepository = githubClient.Repositories.GetLatestRelease(context.Background(), dependency.Owner, dependency.Repo)
		}

		if errGetRepository != nil {
			return errGetRepository
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
	}

	return nil
}
