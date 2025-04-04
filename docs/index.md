---
organization: francois2metz
category: ["internet"]
brand_color: "#6640fe"
display_name: "Gandi"
short_name: "gandi"
description: "Steampipe plugin for querying domains, mailboxes, certificates and more from Gandi."
og_description: "Query Gandi with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/francois2metz/gandi-social-graphic.png"
icon_url: "/images/plugins/francois2metz/gandi.svg"
---

# Gandi + Steampipe

[Gandi](https://gandi.net/) is a registrar and an hosting company.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  fqdn,
  tld,
  owner
from
  gandi_domain
```

```
+--------------------+------+---------------+
| fqdn               | tld  | owner         |
+--------------------+------+---------------+
| caresteouvert.fr   | fr   | francois2metz |
| 2metz.fr           | fr   | francois2metz |
+--------------------+------+---------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/francois2metz/gandi/tables)**

## Get started

### Install

Download and install the latest Gandi plugin:

```bash
steampipe plugin install francois2metz/gandi
```

### Credentials

| Item        | Description                                                                                                                                                                            |
|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | Gandi requires a [token](https://admin.gandi.net/organizations/account/pat).                                                                                                           |
| Radius      | Each connection represents a single Gandi account.                                                                                                                                     |
| Permissions | Tokens have the same permissions as the user who creates them. The personal access token may have access to different permissions.                                                     |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/gandi.spc`)<br />2. Credentials specified in environment variables, e.g., `GANDI_TOKEN` or `GANDI_KEY`. |

### Configuration

Installing the latest gandi plugin will create a config file (`~/.steampipe/config/gandi.spc`) with a single connection named `gandi`:

```hcl
connection "gandi" {
    plugin = "francois2metz/gandi"

    # The Personal Access Token (create it here: https://admin.gandi.net/organizations/account/pat)
    # Permissions:
    #  - View organization
    #  - See and renew domain names
    #  - See & download SSL certificates
    # token = "YOUR_GANDI_TOKEN"
}
```

### Credentials from Environment Variables

The Gandi plugin will use the following environment variables to obtain credentials **only if other argument (`token` or `key`) is not specified** in the connection:

```sh
export GANDI_TOKEN=YOUR_GANDI_TOKEN
```

## Get Involved

* Open source: https://github.com/francois2metz/steampipe-plugin-gandi
