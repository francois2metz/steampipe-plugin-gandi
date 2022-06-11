package gandi

import (
	"context"

	"github.com/go-gandi/go-gandi"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableGandiLivednsRecord() *plugin.Table {
	return &plugin.Table{
		Name:        "gandi_livedns_record",
		Description: "List gandi livedns records.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("domain"),
			Hydrate:    listLivednsRecord,
		},
		Columns: []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Transform: transform.FromQual("domain"), Description: "Domain name."},

			{Name: "rrset_href", Type: proto.ColumnType_STRING, Description: "URL for the record."},
			{Name: "rrset_name", Type: proto.ColumnType_STRING, Description: "Name of the record."},
			{Name: "rrset_type", Type: proto.ColumnType_STRING, Description: "One of: A, AAAA, ALIAS, CAA, CDS, CNAME, DNAME, DS, KEY, LOC, MX, NAPTR, NS, OPENPGPKEY, PTR, RP, SPF, SRV, SSHFP, TLSA, TXT, WKS."},
			{Name: "rrset_values", Type: proto.ColumnType_JSON, Description: "A list of values for this record."},
			{Name: "rrset_ttl", Type: proto.ColumnType_STRING, Description: "The time in seconds that DNS resolvers should cache this record."},
		},
	}
}

func listLivednsRecord(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_livedns_record.listLivednsRecord", "connection_error", err)
		return nil, err
	}

	domain := d.KeyColumnQuals["domain"].GetStringValue()

	client := gandi.NewLiveDNSClient(*config)
	records, err := client.GetDomainRecords(domain)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_livedns_record.listLivednsRecord", err)
		return nil, err
	}
	for _, record := range records {
		d.StreamListItem(ctx, record)
	}
	return nil, nil
}
