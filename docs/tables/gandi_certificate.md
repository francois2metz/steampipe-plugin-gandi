# Table: gandi_certificate

The `gandi_certificate` table can be used to query information about your certificates.

## Examples

### List web redirections of a domain

```sql
select
  cn, status
from
  gandi_certificate
```
