package gandi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-gandi",
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: isNotFoundError,
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"gandi_domain":          tableGandiDomain(),
			"gandi_certificate":     tableGandiCertificate(),
			"gandi_forward":         tableGandiForward(),
			"gandi_livedns_record":  tableGandiLivednsRecord(),
			"gandi_mailbox":         tableGandiMailbox(),
			"gandi_web_redirection": tableGandiWebRedirection(),
		},
	}
	return p
}
