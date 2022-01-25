package gandi

import (
	"context"

	"github.com/go-gandi/go-gandi"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableGandiCertificate() *plugin.Table {
	return &plugin.Table{
		Name:        "gandi_certificate",
		Description: "List gandi certificates.",
		List: &plugin.ListConfig{
			Hydrate: listCertificate,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "UUID."},
			{Name: "cn", Type: proto.ColumnType_STRING, Transform: transform.FromField("CN"), Description: "Common Name."},
			{Name: "package_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Package.Name"), Description: ""},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "One of: 'pending', 'valid', 'revoked', 'replaced', 'replaced_rev', 'expired'."},
		},
	}
}

func listCertificate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config, err := connect(ctx, d)
	client := gandi.NewCertificateClient(*config)

	if err != nil {
		return nil, err
	}
	certificates, err := client.ListCertificates()
	if err != nil {
		return nil, err
	}
	for _, certificate := range certificates {
		d.StreamListItem(ctx, certificate)
	}
	return nil, nil
}
