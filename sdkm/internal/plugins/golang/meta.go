package golang

import (
	pluginGoConsts "github.com/itbasis/go-tools/sdkm/internal/plugins/golang/consts"
	sdkmPlugin "github.com/itbasis/go-tools/sdkm/pkg/plugin"
)

var Meta = sdkmPlugin.MetaInfo{
	ID:             pluginGoConsts.PluginID,
	GetPluginFunc:  GetPlugin,
	ExtCommandInit: CmdExtPlugin,
}
