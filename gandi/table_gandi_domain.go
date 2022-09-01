package gandi

import (
	"context"

	"github.com/go-gandi/go-gandi"
	"github.com/go-gandi/go-gandi/domain"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableGandiDomain() *plugin.Table {
	return &plugin.Table{
		Name:        "gandi_domain",
		Description: "List gandi domains.",
		List: &plugin.ListConfig{
			Hydrate: listDomain,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Unique id of the domain.",
			},
			{
				Name:        "fqdn",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("FQDN"),
				Description: "Fully qualified domain name, written in its native alphabet (IDN).",
			},
			{
				Name:        "fqdn_unicode",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("FQDNUnicode"),
				Description: "Fully qualified domain name, written in unicode.",
			},
			{
				Name:        "tld",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TLD"),
				Description: "The top-level domain.",
			},
			{
				Name:        "domain_owner",
				Type:        proto.ColumnType_STRING,
				Description: "The full name of the owner.",
			},
			{
				Name:        "orga_owner",
				Type:        proto.ColumnType_STRING,
				Description: "The username of the organization owner.",
			},
			{
				Name:        "owner",
				Type:        proto.ColumnType_STRING,
				Description: "The username of the owner.",
			},
			{
				Name:        "auto_renew",
				Type:        proto.ColumnType_BOOL,
				Description: "Automatic renewal status.",
			},
			{
				Name:        "sharing_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SharingID"),
				Description: "The id of the organization.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_JSON,
				Description: "The status of the domain.",
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: "Tags associated to this domain.",
			},
			{
				Name:        "dates_registry_created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Dates.RegistryCreatedAt"),
				Description: "The date the domain was created on the registry.",
			},
			{
				Name:        "dates_created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Dates.CreatedAt"),
				Description: "The date the domain started to be handled by Gandi.",
			},
			{
				Name:        "dates_updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Dates.UpdatedAt"),
				Description: "The last update date.",
			},
			{
				Name:        "dates_deletes_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Dates.DeletesAt"),
				Description: "The date on which the domain will be deleted at the registry.",
			},
			{
				Name:        "dates_hold_begins_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Dates.HoldBeginsAt"),
				Description: "The date from which the domain is held.",
			},
			{
				Name:        "dates_hold_ends_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Dates.HoldEndsAt"),
				Description: "The date from which the domain can’t be renewed anymore (the domain can be restored if the registry supports redemption period otherwise the domain might be destroyed at Gandi at that date).",
			},
			{
				Name:        "dates_pending_delete_ends_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Dates.PendingDeleteEndsAt"),
				Description: "The date from which the domain will be available after a deletion.",
			},
			{
				Name:        "dates_registry_ends_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Dates.RegistryEndsAt"),
				Description: "The date the domain will end at the registry.",
			},
			{
				Name:        "livedns_current",
				Hydrate:     getLiveDNS,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Current"),
				Description: "Type of nameservers currently set. classic corresponds to Gandi's classic nameservers, livedns is for the new, default, Gandi nameservers, premium_dns indicates the presence of Gandi's Premium DNS nameserver and the corresponding service subscription, and other is for custom nameservers.",
			},
			{
				Name:        "nameservers",
				Hydrate:     getLiveDNS,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Nameservers"),
				Description: "List of current nameservers.",
			},
			{
				Name:        "dnssec_available",
				Hydrate:     getLiveDNS,
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("DNSSECAvailable"),
				Description: "Indicates if DNSSEC may be applied to the domain.",
			},
			{
				Name:        "livednssec_available",
				Hydrate:     getLiveDNS,
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("LiveDNSSECAvailable"),
				Description: "Indicates if DNSSEC with liveDNS may be applied to this domain.",
			},
		},
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_domain.listDomain", "connection_error", err)
		return nil, err
	}

	client := gandi.NewDomainClient(*config)
	domains, err := client.ListDomains()
	if err != nil {
		plugin.Logger(ctx).Error("gandi_domain.listDomain", err)
		return nil, err
	}
	for _, domain := range domains {
		d.StreamListItem(ctx, domain)
	}
	return nil, nil
}

func getLiveDNS(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getLiveDNS")
	domain := h.Item.(domain.ListResponse)

	config, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_domain.getLiveDNS", "connection_error", err)
		return nil, err
	}

	client := gandi.NewDomainClient(*config)
	liveDNS, err := client.GetLiveDNS(domain.FQDN)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_domain.getLiveDNS", err)
		return nil, err
	}
	return liveDNS, nil
}
