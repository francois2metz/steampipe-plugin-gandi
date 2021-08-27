package gandi

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type gandiConfig struct {
	Apikey *string `cty:"apikey"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"apikey": {
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
