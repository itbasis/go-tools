package plugins

import (
	sdkmPluginGo "github.com/itbasis/tools/sdkm/internal/plugins/golang"
	pluginGoConsts "github.com/itbasis/tools/sdkm/internal/plugins/golang/consts"
	sdkmPlugin "github.com/itbasis/tools/sdkm/pkg/plugin"
)

var (
	PluginNames = []string{pluginGoConsts.PluginName}

	Plugins = map[string]sdkmPlugin.GetPluginFunc{
		pluginGoConsts.PluginName: sdkmPluginGo.GetPlugin,
	}
)
