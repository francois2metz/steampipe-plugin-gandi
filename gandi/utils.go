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

	key := os.Getenv("GANDI_KEY")

	gandiConfig := GetConfig(d.Connection)

	if gandiConfig.Key != nil {
		key = *gandiConfig.Key
	}

	if key == "" {
		return nil, errors.New("'key' must be set in the connection configuration. Edit your connection configuration file or set the GANDI_KEY environment variable and then restart Steampipe")
	}

	config := &config.Config{APIKey: key, Timeout: -1}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, config)

	return config, nil
}
