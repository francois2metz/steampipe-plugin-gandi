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

	apikey := os.Getenv("GANDI_APIKEY")

	gandiConfig := GetConfig(d.Connection)
	if &gandiConfig != nil {
		if gandiConfig.Apikey != nil {
			apikey = *gandiConfig.Apikey
		}
	}

	if apikey == "" {
		return nil, errors.New("'apikey' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	config := gandi.Config{}
	client := gandi.NewDomainClient(apikey, config)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
