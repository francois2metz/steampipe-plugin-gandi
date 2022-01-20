package gandi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableGandiDomain() *plugin.Table {
	return &plugin.Table{
		Name:        "gandi_domain",
		Description: "List gandi domains.",
		List: &plugin.ListConfig{
			Hydrate: listDomain,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique id of the domain."},
			{Name: "fqdn", Type: proto.ColumnType_STRING, Transform: transform.FromField("FQDN"), Description: "Fully qualified domain name, written in its native alphabet (IDN)."},
			{Name: "fqdn_unicode", Type: proto.ColumnType_STRING, Transform: transform.FromField("FQDNUnicode"), Description: "Fully qualified domain name, written in unicode."},
			{Name: "tld", Type: proto.ColumnType_STRING, Transform: transform.FromField("TLD"), Description: "The top-level domain."},
			{Name: "domain_owner", Type: proto.ColumnType_STRING, Description: "The full name of the owner."},
			{Name: "orga_owner", Type: proto.ColumnType_STRING, Description: "The username of the organization owner."},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "The username of the owner."},
			{Name: "auto_renew", Type: proto.ColumnType_BOOL, Description: "Automatic renewal status."},
			{Name: "sharing_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("SharingID"), Description: "The id of the organization."},
			{Name: "status", Type: proto.ColumnType_JSON, Description: "The status of the domain."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags associated to this domain."},
			{Name: "dates_registry_created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Dates.RegistryCreatedAt"), Description: ""},
			{Name: "dates_created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Dates.CreatedAt"), Description: "The date the domain started to be handled by Gandi."},
			{Name: "dates_updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Dates.UpdatedAt"), Description: "The last update date."},
			{Name: "dates_deletes_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Dates.DeletesAt"), Description: "The date on which the domain will be deleted at the registry."},
			{Name: "dates_hold_begins_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Dates.HoldBeginsAt"), Description: "The date from which the domain is held."},
			{Name: "dates_hold_ends_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Dates.HoldEndsAt"), Description: "The date from which the domain can’t be renewed anymore (the domain can be restored if the registry supports redemption period otherwise the domain might be destroyed at Gandi at that date)."},
			{Name: "dates_pending_delete_ends_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Dates.PendingDeleteEndsAt"), Description: "The date from which the domain will be available after a deletion."},
			{Name: "dates_registry_ends_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Dates.RegistryEndsAt"), Description: ""},
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
