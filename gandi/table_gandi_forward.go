package gandi

import (
	"context"

	"github.com/go-gandi/go-gandi"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableGandiForward() *plugin.Table {
	return &plugin.Table{
		Name:        "gandi_forward",
		Description: "List gandi email forwards.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("domain"),
			Hydrate:    listForward,
		},
		Columns: []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Transform: transform.FromQual("domain"), Description: "Domain name."},

			{Name: "source", Type: proto.ColumnType_STRING, Description: "The source email address."},
			{Name: "destinations", Type: proto.ColumnType_JSON, Description: "A list of email addresses."},

			{Name: "href", Type: proto.ColumnType_STRING, Description: "URL to forwarding address."},
		},
	}
}

func listForward(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_forward.listForward", "connection_error", err)
		return nil, err
	}

	domain := d.KeyColumnQuals["domain"].GetStringValue()

	client := gandi.NewEmailClient(*config)
	mailboxes, err := client.GetForwards(domain)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_forward.listForward", err)
		return nil, err
	}
	for _, mailbox := range mailboxes {
		d.StreamListItem(ctx, mailbox)
	}
	return nil, nil
}
