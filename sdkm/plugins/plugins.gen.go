package plugins

import (
	sdkmPluginGo "github.com/itbasis/go-tools/sdkm/internal/plugins/golang"
	pluginGoConsts "github.com/itbasis/go-tools/sdkm/internal/plugins/golang/consts"
	"github.com/itbasis/go-tools/sdkm/pkg/plugin"
)

var (
	_plugins = map[plugin.ID]plugin.MetaInfo{
		pluginGoConsts.PluginID: sdkmPluginGo.Meta,
	}

	PluginNames = []string{string(pluginGoConsts.PluginID)}
)
