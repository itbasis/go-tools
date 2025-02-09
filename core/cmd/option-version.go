package cmd

import (
	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
	itbasisCoreVersion "github.com/itbasis/go-tools/core/version"
	"github.com/spf13/cobra"
)

const _optionVersionKey itbasisCoreOption.Key = "option-version"

func WithDefaultVersion() Option {
	return &_optionVersion{version: itbasisCoreVersion.NewDefaultVersion()}
}
func WithCustomVersion(version itbasisCoreVersion.Version) Option {
	return &_optionVersion{version: version}
}

type _optionVersion struct {
	version itbasisCoreVersion.Version
}

func (r *_optionVersion) Key() itbasisCoreOption.Key { return _optionVersionKey }
func (r *_optionVersion) Apply(cmd *cobra.Command) error {
	cmd.Version = r.version.String()

	return nil
}
