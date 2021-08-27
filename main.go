package main

import (
	"github.com/francois2metz/steampipe-plugin-gandi/gandi"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: gandi.Plugin})
}
