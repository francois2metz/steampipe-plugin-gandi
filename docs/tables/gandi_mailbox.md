# Table: gandi_mailbox

The `gandi_mailbox` table can be used to query information about your mailboxes and you must specify which domain in the where or join clause using the `domain` column.

## Examples

### List mailboxes of a domain

```sql
select
  address
from
  gandi_mailbox
where
  domain='example.net';
```

### List maiboxes across all your domains

```sql
select
  m.domain, m.address
from
  gandi_mailbox m
join
  gandi_domain d
on
  d.fqdn = m.domain;
```

### List maiboxes that are 90% of their quota

```sql
select
  m.domain,
  m.address,
  m.quota_used
from
  gandi_mailbox m
join
  gandi_domain d
on
  d.fqdn = m.domain
where
  case
    when m.mailbox_type = 'premium' then (100 * m.quota_used) / 50000000
    else (100 * m.quota_used) / 3000000
  end > 90;
```
