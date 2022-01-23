package gandi

import (
	"context"

	"github.com/go-gandi/go-gandi"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableGandiEmail() *plugin.Table {
	return &plugin.Table{
		Name:        "gandi_email",
		Description: "List gandi email.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("domain"),
			Hydrate:    listEmail,
		},
		Columns: []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Transform: transform.FromQual("domain"), Description: "Domain name."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique id of the mailbox."},
			{Name: "address", Type: proto.ColumnType_STRING, Description: "Full email address."},

			{Name: "href", Type: proto.ColumnType_STRING, Description: "Link to mailbox details."},
			{Name: "login", Type: proto.ColumnType_STRING, Description: "Mailbox login."},
			{Name: "mailbox_type", Type: proto.ColumnType_STRING, Description: "One of: 'standard', 'premium', 'free'."},
			{Name: "quota_used", Type: proto.ColumnType_INT, Transform: transform.FromField("QuataUsed"), Description: "Quota used."},
		},
	}
}

func listEmail(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config, err := connect(ctx, d)
	client := gandi.NewEmailClient(*config)

	if err != nil {
		return nil, err
	}
	domain := d.KeyColumnQuals["domain"].GetStringValue()
	mailboxes, err := client.ListMailboxes(domain)
	if err != nil {
		return nil, err
	}
	for _, mailbox := range mailboxes {
		d.StreamListItem(ctx, mailbox)
	}
	return nil, nil
}
