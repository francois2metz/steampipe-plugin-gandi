package gandi

import (
	"context"
	"errors"
	"os"

	"github.com/go-gandi/go-gandi/config"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*config.Config, error) {
	// get gandi client from cache
	cacheKey := "gandi"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*config.Config), nil
	}

	token := os.Getenv("GANDI_TOKEN")

	gandiConfig := GetConfig(d.Connection)

	if gandiConfig.Token != nil {
		token = *gandiConfig.Token
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file or set the GANDI_TOKEN environment variable and then restart Steampipe")
	}

	config := &config.Config{PersonalAccessToken: token, Timeout: -1}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, config)

	return config, nil
}
