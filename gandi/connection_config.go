package gandi

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/schema"
)

type gandiConfig struct {
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &gandiConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) gandiConfig {
	if connection == nil || connection.GetConfig() == nil {
		return gandiConfig{}
	}
	config, _ := connection.GetConfig().(gandiConfig)
	return config
}
