# Table: gandi_mailbox

The `gandi_mailbox` table can be used to query information about your mailboxes and you must specify which domain in the where or join clause using the `domain` column.

## Examples

### List mailboxes of a domain

```sql
select
  address, quota_used
from
  gandi_mailbox
where
  domain='example.net';
```
