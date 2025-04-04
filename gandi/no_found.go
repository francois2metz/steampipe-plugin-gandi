package gandi

import (
	"context"
	"net/http"

	ganditypes "github.com/go-gandi/go-gandi/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func isNotFoundError(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
	requestError, ok := err.(*ganditypes.RequestError)
	if !ok {
		return false
	}

	if requestError.StatusCode == http.StatusNotFound {
		return true
	}
	return false
}
