package cmd

import (
	itbasisCoreOption "github.com/itbasis/tools/core/option"
	itbasisCoreVersion "github.com/itbasis/tools/core/version"
	"github.com/spf13/cobra"
)

const _optionVersionKey = "option-version"

func WithDefaultVersion() itbasisCoreOption.Option[cobra.Command] {
	return &_optionVersion{version: itbasisCoreVersion.NewDefaultVersion()}
}
func WithCustomVersion(version itbasisCoreVersion.Version) itbasisCoreOption.Option[cobra.Command] {
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
