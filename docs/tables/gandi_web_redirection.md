# Table: gandi_web_redirection

The `gandi_web_redirection` table can be used to query information about your web redirections and you must specify which domain in the where or join clause using the `domain` column.

## Examples

### List web redirections of a domain

```sql
select
  host, url, protocol
from
  gandi_web_redirection
where
  domain='example.net';
```
