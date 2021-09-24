package gandi

import (
	"context"
	"errors"
	"os"

	"github.com/go-gandi/go-gandi"
	"github.com/go-gandi/go-gandi/domain"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*domain.Domain, error) {
	// get gandi client from cache
	cacheKey := "gandi"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*domain.Domain), nil
	}

	key := os.Getenv("GANDI_KEY")

	gandiConfig := GetConfig(d.Connection)
	if &gandiConfig != nil {
		if gandiConfig.Key != nil {
			key = *gandiConfig.Key
		}
	}

	if key == "" {
		return nil, errors.New("'key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	config := gandi.Config{}
	client := gandi.NewDomainClient(key, config)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
