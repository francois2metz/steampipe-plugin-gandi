package gandi

import (
	"context"

	"github.com/go-gandi/go-gandi"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableGandiWebRedirection() *plugin.Table {
	return &plugin.Table{
		Name:        "gandi_web_redirection",
		Description: "List gandi web redirections.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("domain"),
			Hydrate:    listWebRedirection,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"domain", "host"}),
			Hydrate:    getWebRedirection,
		},
		Columns: []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Transform: transform.FromQual("domain"), Description: "Domain name."},

			{Name: "host", Type: proto.ColumnType_STRING, Description: "Source hostname (including the domain name)."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of redirection. One of: 'cloak', 'http301', 'http302'"},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "Target URL."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Last update date."},
			{Name: "protocol", Type: proto.ColumnType_STRING, Description: "One of: 'http', 'https', 'httpsonly'."},
		},
	}
}

func listWebRedirection(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config, err := connect(ctx, d)
	client := gandi.NewDomainClient(*config)

	if err != nil {
		return nil, err
	}
	domain := d.KeyColumnQuals["domain"].GetStringValue()
	redirections, err := client.ListWebRedirections(domain)
	if err != nil {
		return nil, err
	}
	for _, redirection := range redirections {
		d.StreamListItem(ctx, redirection)
	}
	return nil, nil
}

func getWebRedirection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config, err := connect(ctx, d)
	client := gandi.NewDomainClient(*config)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals

	host := quals["host"].GetStringValue()
	domain := quals["domain"].GetStringValue()

	result, err := client.GetWebRedirection(domain, host)
	if err != nil {
		return nil, err
	}
	return result, nil
}
