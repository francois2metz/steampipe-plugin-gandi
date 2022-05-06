package gandi

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type gandiConfig struct {
	Key *string `cty:"key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &gandiConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) gandiConfig {
	if connection == nil || connection.Config == nil {
		return gandiConfig{}
	}
	config, _ := connection.Config.(gandiConfig)
	return config
}
