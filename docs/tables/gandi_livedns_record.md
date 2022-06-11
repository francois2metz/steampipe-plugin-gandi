# Table: gandi_livedns_record

The `gandi_livedns_record` table can be used to query information about your records and you must specify which domain in the where or join clause using the `domain` column.

## Examples

### List records of your domain

```sql
select
  rrset_name, rrset_type, rrset_values
from
  gandi_livedns_record
where
  domain='example.net';
```

### List records with a specific name

```sql
select
  rrset_type, rrset_values
from
  gandi_livedns_record
where
  domain='example.net'
  and rrset_name='test';
```

### List CNAME records

```sql
select
  rrset_name, rrset_values
from
  gandi_livedns_record
where
  domain='example.net'
  and rrset_type='CNAME';
```

### Get entries with a specific value

```sql
select
  rrset_name,
from
  gandi_livedns_record
where
  domain='example.net'
  and rrset_values ? 'test';
```
