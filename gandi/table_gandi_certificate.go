package gandi

import (
	"context"

	"github.com/go-gandi/go-gandi"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableGandiCertificate() *plugin.Table {
	return &plugin.Table{
		Name:        "gandi_certificate",
		Description: "List gandi certificates.",
		List: &plugin.ListConfig{
			Hydrate: listCertificate,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getCertificate,
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
	if err != nil {
		plugin.Logger(ctx).Error("gandi_certificate.listCertificate", "connection_error", err)
		return nil, err
	}

	client := gandi.NewCertificateClient(*config)
	certificates, err := client.ListCertificates()
	if err != nil {
		plugin.Logger(ctx).Error("gandi_certificate.listCertificate", err)
		return nil, err
	}
	for _, certificate := range certificates {
		d.StreamListItem(ctx, certificate)
	}
	return nil, nil
}

func getCertificate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_certificate.getCertificate", "connection_error", err)
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetStringValue()

	client := gandi.NewCertificateClient(*config)
	result, err := client.GetCertificate(id)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_certificate.getCertificate", err)
		return nil, err
	}
	return result, nil
}
