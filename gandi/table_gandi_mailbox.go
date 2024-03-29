package gandi

import (
	"context"

	"github.com/go-gandi/go-gandi"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGandiMailbox() *plugin.Table {
	return &plugin.Table{
		Name:        "gandi_mailbox",
		Description: "List gandi mailboxes.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("domain"),
			Hydrate:    listMailbox,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"domain", "id"}),
			Hydrate:    getMailbox,
		},
		Columns: []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Transform: transform.FromQual("domain"), Description: "Domain name."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique id of the mailbox."},

			{Name: "address", Type: proto.ColumnType_STRING, Description: "Full email address."},
			{Name: "antispam", Type: proto.ColumnType_BOOL, Description: "Antispam is enabled."},
			{Name: "autorenew_duration", Type: proto.ColumnType_INT, Transform: transform.FromField("Autorenew.Duration"), Description: "Duration for autorenew."},
			{Name: "autorenew_duration_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Autorenew.DurationType"), Description: "Type of duration ('m' for month)."},
			{Name: "autorenew_enabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Autorenew.Enabled"), Description: "Specify if autorenew is enabled on this mailbox."},
			{Name: "expires_at", Type: proto.ColumnType_TIMESTAMP, Description: "Expiry date of the mailbox."},
			{Name: "href", Type: proto.ColumnType_STRING, Description: "Link to mailbox details."},
			{Name: "login", Type: proto.ColumnType_STRING, Description: "Mailbox login."},
			{Name: "mailbox_type", Type: proto.ColumnType_STRING, Description: "The type of mailbox, one of: 'standard', 'premium', 'free'."},
			{Name: "quota_used", Type: proto.ColumnType_INT, Description: "Quota used."},
		},
	}
}

func listMailbox(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_mailbox.listMailbox", "connection_error", err)
		return nil, err
	}

	domain := d.EqualsQuals["domain"].GetStringValue()

	client := gandi.NewEmailClient(*config)
	mailboxes, err := client.ListMailboxes(domain)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_mailbox.listMailbox", err)
		return nil, err
	}
	for _, mailbox := range mailboxes {
		d.StreamListItem(ctx, mailbox)
	}
	return nil, nil
}

func getMailbox(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_mailbox.getMailbox", "connection_error", err)
		return nil, err
	}

	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	domain := quals["domain"].GetStringValue()

	client := gandi.NewEmailClient(*config)
	result, err := client.GetMailbox(domain, id)
	if err != nil {
		plugin.Logger(ctx).Error("gandi_mailbox.getMailbox", err)
		return nil, err
	}
	return result, nil
}
