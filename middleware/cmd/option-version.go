package cmd

import (
	itbasisMiddlewareOption "github.com/itbasis/tools/middleware/option"
	itbasisMiddlewareVersion "github.com/itbasis/tools/middleware/version"
	"github.com/spf13/cobra"
)

const _optionVersionKey = "option-version"

func WithDefaultVersion() itbasisMiddlewareOption.Option[cobra.Command] {
	return &_optionVersion{version: itbasisMiddlewareVersion.NewDefaultVersion()}
}
func WithCustomVersion(version itbasisMiddlewareVersion.Version) itbasisMiddlewareOption.Option[cobra.Command] {
	return &_optionVersion{version: version}
}

type _optionVersion struct {
	version itbasisMiddlewareVersion.Version
}

func (r *_optionVersion) Key() itbasisMiddlewareOption.Key { return _optionVersionKey }
func (r *_optionVersion) Apply(cmd *cobra.Command) error {
	cmd.Version = r.version.String()

	return nil
}
