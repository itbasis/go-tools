package cmd

import (
	"strings"

	"github.com/itbasis/tools/sdkm/plugins"
)

const (
	ArgAliasPlugin = "plugin"
)

var (
	UseArgRequirePlugins  = "{" + strings.Join(plugins.PluginNames, "|") + "}"
	UseArgOptionalPlugins = "[" + strings.Join(plugins.PluginNames, "|") + "]"
)
