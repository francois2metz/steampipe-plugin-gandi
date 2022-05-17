# Table: gandi_certificate

The `gandi_certificate` table can be used to query information about your certificates.

## Examples

### List certificates

```sql
select
  cn, status
from
  gandi_certificate
```

### List invalid certificates

```sql
select
  cn, status
from
  gandi_certificate
where
  status != 'valid'
```
