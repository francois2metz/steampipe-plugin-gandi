# Table: gandi_domain

The `gandi_domain` table can be used to query information about your domains.

## Examples

### List domains

```sql
select
  fqdn
from
  gandi_domain;
```

### List domains without auto renew set

```sql
select
  fqdn
from
  gandi_domain
where
  auto_renew=false;
```

### List domains expiring in less than one month

```sql
select
  fqdn
from
  gandi_domain
where
  dates_registry_ends_at < now() + interval '1 month';
```

### List domains using classic namervers

```sql
select
  fqdn
from
  gandi_domain
where
  livedns_current = 'classic';
```
