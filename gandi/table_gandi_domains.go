package gandi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableGandiDomains() *plugin.Table {
	return &plugin.Table{
		Name:        "gandi_domains",
		Description: "List gandi domains",
		List: &plugin.ListConfig{
			Hydrate: listDomain,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getDomain,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "unique id of the domain"},
			{Name: "fqdn", Type: proto.ColumnType_STRING, Transform: transform.FromField("FQDN"), Description: ""},
			{Name: "fqdn_unicode", Type: proto.ColumnType_STRING, Transform: transform.FromField("FQDNUnicode"), Description: ""},
			{Name: "tld", Type: proto.ColumnType_STRING, Transform: transform.FromField("TLD"), Description: ""},
			{Name: "href", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "domain_owner", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "orga_owner", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "nameserver_current", Type: proto.ColumnType_STRING, Transform: transform.FromField("NameServerConfig.Current"), Description: ""},
			{Name: "auto_renew", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "sharing_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("SharingID"), Description: ""},
			{Name: "registry_created_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.RegistryCreatedAt"), Description: ""},
			{Name: "created_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.CreatedAt"), Description: ""},
			{Name: "updated_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.UpdatedAt"), Description: ""},
			{Name: "deletes_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.DeletesAt"), Description: ""},
			{Name: "auth_info_expires_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.AuthInfoExpiresAt"), Description: ""},
			{Name: "hold_begins_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.HoldBeginsAt"), Description: ""},
			{Name: "hold_ends_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.HoldEndsAt"), Description: ""},
			{Name: "pending_delete_ends_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.PendingDeleteEndsAt"), Description: ""},
			{Name: "registry_ends_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.RegistryEndsAt"), Description: ""},
			{Name: "renew_begins_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.RenewBeginsAt"), Description: ""},
			{Name: "renew_ends_at", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Dates.RenewEndsAt"), Description: ""},
		},
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	domains, err := client.ListDomains()
	if err != nil {
		return nil, err
	}
	for _, domain := range domains {
		d.StreamListItem(ctx, domain)
	}
	return nil, nil
}

func getDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals
	name := quals["name"].GetStringValue()
	result, err := client.GetDomain(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}
